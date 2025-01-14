package serverservice

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"go.hollow.sh/serverservice/internal/models"
)

var (
	errSrvComponentPayload = errors.New("error in server component payload")
)

// serverComponentList returns a response with the list of components that matched the params.
func (r *Router) serverComponentList(c *gin.Context) {
	pager := parsePagination(c)

	params, err := parseQueryServerComponentsListParams(c)
	if err != nil {
		badRequestResponse(c, "invalid server component list params", err)
		return
	}

	dbSC, count, err := r.getServerComponents(c, params, pager)
	if err != nil {
		dbErrorResponse(c, err)
		return
	}

	serverComponents := ServerComponentSlice{}

	for _, dbSC := range dbSC {
		sc := ServerComponent{}
		if err := sc.fromDBModel(dbSC); err != nil {
			failedConvertingToVersioned(c, err)
			return
		}

		serverComponents = append(serverComponents, sc)
	}

	pd := paginationData{
		pageCount:  len(serverComponents),
		totalCount: count,
		pager:      pager,
	}

	listResponse(c, serverComponents, pd)
}

// serverComponentGet returns a response with the list of components referenced by the server UUID.
func (r *Router) serverComponentGet(c *gin.Context) {
	srv, err := r.loadServerFromParams(c)
	if err != nil {
		if errors.Is(err, ErrUUIDParse) {
			badRequestResponse(c, "", err)
			return
		}

		dbErrorResponse(c, err)

		return
	}

	pager := parsePagination(c)

	// - include Attributes, VersionedAttributes and ServerComponentyType relations
	mods := []qm.QueryMod{
		qm.Load("Attributes"),
		qm.Load("VersionedAttributes", qm.Where("(namespace, created_at, server_component_id) IN (select namespace, max(created_at), server_component_id from versioned_attributes group by namespace, server_component_id)")),
		qm.Load("ServerComponentType"),
	}

	dbComps, err := srv.ServerComponents(mods...).All(c.Request.Context(), r.DB)
	if err != nil {
		dbErrorResponse(c, err)
		return
	}

	count := int64(0)

	comps, err := convertDBServerComponents(dbComps)
	if err != nil {
		failedConvertingToVersioned(c, err)
		return
	}

	pd := paginationData{
		pageCount:  len(comps),
		totalCount: count,
		pager:      pager,
	}

	listResponse(c, comps, pd)
}

// serverComponentsCreate stores a ServerComponentSlice object into the backend store.
func (r *Router) serverComponentsCreate(c *gin.Context) {
	// load server based on the UUID parameter
	server, err := r.loadServerFromParams(c)
	if err != nil {
		if errors.Is(err, ErrUUIDParse) {
			badRequestResponse(c, "", err)
			return
		}

		dbErrorResponse(c, err)

		return
	}

	// check server exists
	if server == nil {
		notFoundResponse(c, "server resource referenced by UUID does not exist: "+server.ID)
		return
	}

	// components payload
	var serverComponents ServerComponentSlice
	if err := c.ShouldBindJSON(&serverComponents); err != nil {
		badRequestResponse(
			c,
			"",
			errors.Wrap(
				errSrvComponentPayload, "failed to unmarshal JSON as ServerComponentSlice: "+err.Error()),
		)

		return
	}

	if len(serverComponents) == 0 {
		badRequestResponse(
			c,
			"",
			errors.Wrap(errSrvComponentPayload, "ServerComponentSlice is empty"),
		)

		return
	}

	// component data is written in a transaction along with versioned attributes
	tx, err := r.DB.BeginTx(c.Request.Context(), nil)
	if err != nil {
		dbErrorResponse(c, err)
		return
	}

	// rollback is a no-op when the transaction is successful
	// nolint:errcheck // TODO(joel): log gerror instead of ignoring
	defer tx.Rollback()

	for _, srvComponent := range serverComponents {
		dbSrvComponent := srvComponent.toDBModel(server.ID)

		// Set server component UUID.
		//
		// The INSERT into the Attributes and VersionedAttributes has a check constraint
		// for server_id (dbSrvComponent.ServerUUID), server_component_id (dbSrvComponent.ID) being NOT NULL,
		//
		// Generally the INSERT into the server_components table returns a UUID generated by the database
		// and the dbSrvComponent.ID is set to the returned UUID,
		//
		// Although, since we're in a transaction here which
		// INSERTs the component data along with the attributes, versioned attributes in separate statements,
		// the dbSrvComponent.ID is not set. For this to work, it would require a CTE within which
		// the returning ID can be assigned to the dbSrvComponent.ID.
		//
		// For now its easier to just set the UUID here.
		if dbSrvComponent.ID == uuid.Nil.String() {
			dbSrvComponent.ID = uuid.New().String()
		}

		// insert component
		err := dbSrvComponent.Insert(c.Request.Context(), tx, boil.Infer())
		if err != nil {
			dbErrorResponse(c, err)
			return
		}

		// insert versioned attributes
		for _, versionedAttributes := range srvComponent.VersionedAttributes {
			dbVersionedAttributes := versionedAttributes.toDBModel()
			dbVersionedAttributes.ServerComponentID = null.StringFrom(dbSrvComponent.ID)

			err = dbSrvComponent.AddVersionedAttributes(c.Request.Context(), tx, true, dbVersionedAttributes)
			if err != nil {
				dbErrorResponse(c, err)
				return
			}
		}

		// insert attributes
		for _, attributes := range srvComponent.Attributes {
			dbAttributes, err := attributes.toDBModel()
			if err != nil {
				dbErrorResponse(c, err)
				return
			}

			dbAttributes.ServerComponentID = null.StringFrom(dbSrvComponent.ID)

			err = dbSrvComponent.AddAttributes(c.Request.Context(), tx, true, dbAttributes)
			if err != nil {
				dbErrorResponse(c, err)
				return
			}
		}
	}

	if err := tx.Commit(); err != nil {
		dbErrorResponse(c, err)
		return
	}

	createdResponse(c, "")
}

// serverComponentUpdate updates existing server component attributes
func (r *Router) serverComponentUpdate(c *gin.Context) {
	// load server based on the UUID parameter
	server, err := r.loadServerFromParams(c)
	if err != nil {
		if errors.Is(err, ErrUUIDParse) {
			badRequestResponse(c, "", err)
			return
		}

		dbErrorResponse(c, err)

		return
	}

	// check server exists
	if server == nil {
		notFoundResponse(c, "server resource referenced by UUID does not exist: "+server.ID)
		return
	}

	// components payload
	var serverComponents ServerComponentSlice
	if err := c.ShouldBindJSON(&serverComponents); err != nil {
		badRequestResponse(
			c,
			"",
			errors.Wrap(
				errSrvComponentPayload, err.Error()),
		)

		return
	}

	if len(serverComponents) == 0 {
		badRequestResponse(
			c,
			"",
			errors.Wrap(errSrvComponentPayload, "ServerComponentSlice is empty"),
		)

		return
	}

	// component data is written in a transaction along with versioned attributes
	tx, err := r.DB.BeginTx(c.Request.Context(), nil)
	if err != nil {
		dbErrorResponse(c, err)
		return
	}

	// rollback is a no-op when the transaction is successful
	// nolint:errcheck // TODO(joel): log gerror instead of ignoring
	defer tx.Rollback()

	for _, srvComponent := range serverComponents {
		// convert object to db model type and keep the received component UUID
		dbSrvComponent := srvComponent.toDBModel(server.ID)

		// check component ID is non nil
		if dbSrvComponent.ID == "" || dbSrvComponent.ID == uuid.Nil.String() {
			badRequestResponse(
				c,
				"component update requires a non-nil UUID",
				errSrvComponentPayload,
			)

			return
		}

		// check server component exists
		exists, err := models.ServerComponentExists(c.Request.Context(), tx, srvComponent.UUID.String())
		if err != nil {
			badRequestResponse(c, "check component resource exists error", err)
			return
		}

		if !exists {
			badRequestResponse(
				c,
				"",
				errors.Wrap(
					errSrvComponentPayload, "component resource referenced by UUID does not exist: "+
						srvComponent.UUID.String(),
				),
			)

			return
		}

		// update component
		_, err = dbSrvComponent.Update(c.Request.Context(), tx, boil.Infer())
		if err != nil {
			dbErrorResponse(c, err)
			return
		}

		// update component versioned attributes
		for _, versionedAttributes := range srvComponent.VersionedAttributes {
			dbVersionedAttributes := versionedAttributes.toDBModel()
			dbVersionedAttributes.ServerComponentID = null.StringFrom(dbSrvComponent.ID)

			err = dbSrvComponent.AddVersionedAttributes(c.Request.Context(), tx, true, dbVersionedAttributes)
			if err != nil {
				dbErrorResponse(c, err)
				return
			}
		}

		// update component  attributes
		for _, attributes := range srvComponent.Attributes {
			dbAttributes, err := attributes.toDBModel()
			if err != nil {
				dbErrorResponse(c, err)
				return
			}

			dbAttributes.ServerComponentID = null.StringFrom(dbSrvComponent.ID)

			err = dbSrvComponent.AddAttributes(c.Request.Context(), tx, true, dbAttributes)
			if err != nil {
				dbErrorResponse(c, err)
				return
			}
		}
	}

	if err := tx.Commit(); err != nil {
		dbErrorResponse(c, err)
		return
	}

	updatedResponse(c, "")
}

// serverComponentDelete deletes a server component.
func (r *Router) serverComponentDelete(c *gin.Context) {
	// load server based on the UUID parameter
	server, err := r.loadServerFromParams(c)
	if err != nil {
		if errors.Is(err, ErrUUIDParse) {
			badRequestResponse(c, "", err)
			return
		}

		dbErrorResponse(c, err)

		return
	}

	if _, err := server.ServerComponents().DeleteAll(c.Request.Context(), r.DB); err != nil {
		dbErrorResponse(c, err)

		return
	}

	deletedResponse(c)
}
