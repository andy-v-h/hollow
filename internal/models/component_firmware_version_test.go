// Code generated by SQLBoiler 4.11.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/randomize"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/strmangle"
)

func testComponentFirmwareVersionsUpsert(t *testing.T) {
	t.Parallel()

	if len(componentFirmwareVersionAllColumns) == len(componentFirmwareVersionPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := ComponentFirmwareVersion{}
	if err = randomize.Struct(seed, &o, componentFirmwareVersionDBTypes, true); err != nil {
		t.Errorf("Unable to randomize ComponentFirmwareVersion struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert ComponentFirmwareVersion: %s", err)
	}

	count, err := ComponentFirmwareVersions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, componentFirmwareVersionDBTypes, false, componentFirmwareVersionPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ComponentFirmwareVersion struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert ComponentFirmwareVersion: %s", err)
	}

	count, err = ComponentFirmwareVersions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testComponentFirmwareVersions(t *testing.T) {
	t.Parallel()

	query := ComponentFirmwareVersions()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testComponentFirmwareVersionsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ComponentFirmwareVersion{}
	if err = randomize.Struct(seed, o, componentFirmwareVersionDBTypes, true, componentFirmwareVersionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ComponentFirmwareVersion struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := ComponentFirmwareVersions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testComponentFirmwareVersionsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ComponentFirmwareVersion{}
	if err = randomize.Struct(seed, o, componentFirmwareVersionDBTypes, true, componentFirmwareVersionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ComponentFirmwareVersion struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := ComponentFirmwareVersions().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := ComponentFirmwareVersions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testComponentFirmwareVersionsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ComponentFirmwareVersion{}
	if err = randomize.Struct(seed, o, componentFirmwareVersionDBTypes, true, componentFirmwareVersionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ComponentFirmwareVersion struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ComponentFirmwareVersionSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := ComponentFirmwareVersions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testComponentFirmwareVersionsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ComponentFirmwareVersion{}
	if err = randomize.Struct(seed, o, componentFirmwareVersionDBTypes, true, componentFirmwareVersionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ComponentFirmwareVersion struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := ComponentFirmwareVersionExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if ComponentFirmwareVersion exists: %s", err)
	}
	if !e {
		t.Errorf("Expected ComponentFirmwareVersionExists to return true, but got false.")
	}
}

func testComponentFirmwareVersionsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ComponentFirmwareVersion{}
	if err = randomize.Struct(seed, o, componentFirmwareVersionDBTypes, true, componentFirmwareVersionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ComponentFirmwareVersion struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	componentFirmwareVersionFound, err := FindComponentFirmwareVersion(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if componentFirmwareVersionFound == nil {
		t.Error("want a record, got nil")
	}
}

func testComponentFirmwareVersionsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ComponentFirmwareVersion{}
	if err = randomize.Struct(seed, o, componentFirmwareVersionDBTypes, true, componentFirmwareVersionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ComponentFirmwareVersion struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = ComponentFirmwareVersions().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testComponentFirmwareVersionsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ComponentFirmwareVersion{}
	if err = randomize.Struct(seed, o, componentFirmwareVersionDBTypes, true, componentFirmwareVersionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ComponentFirmwareVersion struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := ComponentFirmwareVersions().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testComponentFirmwareVersionsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	componentFirmwareVersionOne := &ComponentFirmwareVersion{}
	componentFirmwareVersionTwo := &ComponentFirmwareVersion{}
	if err = randomize.Struct(seed, componentFirmwareVersionOne, componentFirmwareVersionDBTypes, false, componentFirmwareVersionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ComponentFirmwareVersion struct: %s", err)
	}
	if err = randomize.Struct(seed, componentFirmwareVersionTwo, componentFirmwareVersionDBTypes, false, componentFirmwareVersionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ComponentFirmwareVersion struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = componentFirmwareVersionOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = componentFirmwareVersionTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := ComponentFirmwareVersions().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testComponentFirmwareVersionsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	componentFirmwareVersionOne := &ComponentFirmwareVersion{}
	componentFirmwareVersionTwo := &ComponentFirmwareVersion{}
	if err = randomize.Struct(seed, componentFirmwareVersionOne, componentFirmwareVersionDBTypes, false, componentFirmwareVersionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ComponentFirmwareVersion struct: %s", err)
	}
	if err = randomize.Struct(seed, componentFirmwareVersionTwo, componentFirmwareVersionDBTypes, false, componentFirmwareVersionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ComponentFirmwareVersion struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = componentFirmwareVersionOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = componentFirmwareVersionTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ComponentFirmwareVersions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func componentFirmwareVersionBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *ComponentFirmwareVersion) error {
	*o = ComponentFirmwareVersion{}
	return nil
}

func componentFirmwareVersionAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *ComponentFirmwareVersion) error {
	*o = ComponentFirmwareVersion{}
	return nil
}

func componentFirmwareVersionAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *ComponentFirmwareVersion) error {
	*o = ComponentFirmwareVersion{}
	return nil
}

func componentFirmwareVersionBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *ComponentFirmwareVersion) error {
	*o = ComponentFirmwareVersion{}
	return nil
}

func componentFirmwareVersionAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *ComponentFirmwareVersion) error {
	*o = ComponentFirmwareVersion{}
	return nil
}

func componentFirmwareVersionBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *ComponentFirmwareVersion) error {
	*o = ComponentFirmwareVersion{}
	return nil
}

func componentFirmwareVersionAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *ComponentFirmwareVersion) error {
	*o = ComponentFirmwareVersion{}
	return nil
}

func componentFirmwareVersionBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *ComponentFirmwareVersion) error {
	*o = ComponentFirmwareVersion{}
	return nil
}

func componentFirmwareVersionAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *ComponentFirmwareVersion) error {
	*o = ComponentFirmwareVersion{}
	return nil
}

func testComponentFirmwareVersionsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &ComponentFirmwareVersion{}
	o := &ComponentFirmwareVersion{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, componentFirmwareVersionDBTypes, false); err != nil {
		t.Errorf("Unable to randomize ComponentFirmwareVersion object: %s", err)
	}

	AddComponentFirmwareVersionHook(boil.BeforeInsertHook, componentFirmwareVersionBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	componentFirmwareVersionBeforeInsertHooks = []ComponentFirmwareVersionHook{}

	AddComponentFirmwareVersionHook(boil.AfterInsertHook, componentFirmwareVersionAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	componentFirmwareVersionAfterInsertHooks = []ComponentFirmwareVersionHook{}

	AddComponentFirmwareVersionHook(boil.AfterSelectHook, componentFirmwareVersionAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	componentFirmwareVersionAfterSelectHooks = []ComponentFirmwareVersionHook{}

	AddComponentFirmwareVersionHook(boil.BeforeUpdateHook, componentFirmwareVersionBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	componentFirmwareVersionBeforeUpdateHooks = []ComponentFirmwareVersionHook{}

	AddComponentFirmwareVersionHook(boil.AfterUpdateHook, componentFirmwareVersionAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	componentFirmwareVersionAfterUpdateHooks = []ComponentFirmwareVersionHook{}

	AddComponentFirmwareVersionHook(boil.BeforeDeleteHook, componentFirmwareVersionBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	componentFirmwareVersionBeforeDeleteHooks = []ComponentFirmwareVersionHook{}

	AddComponentFirmwareVersionHook(boil.AfterDeleteHook, componentFirmwareVersionAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	componentFirmwareVersionAfterDeleteHooks = []ComponentFirmwareVersionHook{}

	AddComponentFirmwareVersionHook(boil.BeforeUpsertHook, componentFirmwareVersionBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	componentFirmwareVersionBeforeUpsertHooks = []ComponentFirmwareVersionHook{}

	AddComponentFirmwareVersionHook(boil.AfterUpsertHook, componentFirmwareVersionAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	componentFirmwareVersionAfterUpsertHooks = []ComponentFirmwareVersionHook{}
}

func testComponentFirmwareVersionsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ComponentFirmwareVersion{}
	if err = randomize.Struct(seed, o, componentFirmwareVersionDBTypes, true, componentFirmwareVersionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ComponentFirmwareVersion struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ComponentFirmwareVersions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testComponentFirmwareVersionsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ComponentFirmwareVersion{}
	if err = randomize.Struct(seed, o, componentFirmwareVersionDBTypes, true); err != nil {
		t.Errorf("Unable to randomize ComponentFirmwareVersion struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(componentFirmwareVersionColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := ComponentFirmwareVersions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testComponentFirmwareVersionToManyFirmwareComponentFirmwareSetMaps(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a ComponentFirmwareVersion
	var b, c ComponentFirmwareSetMap

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, componentFirmwareVersionDBTypes, true, componentFirmwareVersionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ComponentFirmwareVersion struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, componentFirmwareSetMapDBTypes, false, componentFirmwareSetMapColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, componentFirmwareSetMapDBTypes, false, componentFirmwareSetMapColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.FirmwareID = a.ID
	c.FirmwareID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.FirmwareComponentFirmwareSetMaps().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.FirmwareID == b.FirmwareID {
			bFound = true
		}
		if v.FirmwareID == c.FirmwareID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := ComponentFirmwareVersionSlice{&a}
	if err = a.L.LoadFirmwareComponentFirmwareSetMaps(ctx, tx, false, (*[]*ComponentFirmwareVersion)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.FirmwareComponentFirmwareSetMaps); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.FirmwareComponentFirmwareSetMaps = nil
	if err = a.L.LoadFirmwareComponentFirmwareSetMaps(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.FirmwareComponentFirmwareSetMaps); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testComponentFirmwareVersionToManyAddOpFirmwareComponentFirmwareSetMaps(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a ComponentFirmwareVersion
	var b, c, d, e ComponentFirmwareSetMap

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, componentFirmwareVersionDBTypes, false, strmangle.SetComplement(componentFirmwareVersionPrimaryKeyColumns, componentFirmwareVersionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*ComponentFirmwareSetMap{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, componentFirmwareSetMapDBTypes, false, strmangle.SetComplement(componentFirmwareSetMapPrimaryKeyColumns, componentFirmwareSetMapColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*ComponentFirmwareSetMap{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddFirmwareComponentFirmwareSetMaps(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.FirmwareID {
			t.Error("foreign key was wrong value", a.ID, first.FirmwareID)
		}
		if a.ID != second.FirmwareID {
			t.Error("foreign key was wrong value", a.ID, second.FirmwareID)
		}

		if first.R.Firmware != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Firmware != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.FirmwareComponentFirmwareSetMaps[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.FirmwareComponentFirmwareSetMaps[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.FirmwareComponentFirmwareSetMaps().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testComponentFirmwareVersionsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ComponentFirmwareVersion{}
	if err = randomize.Struct(seed, o, componentFirmwareVersionDBTypes, true, componentFirmwareVersionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ComponentFirmwareVersion struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testComponentFirmwareVersionsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ComponentFirmwareVersion{}
	if err = randomize.Struct(seed, o, componentFirmwareVersionDBTypes, true, componentFirmwareVersionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ComponentFirmwareVersion struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ComponentFirmwareVersionSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testComponentFirmwareVersionsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ComponentFirmwareVersion{}
	if err = randomize.Struct(seed, o, componentFirmwareVersionDBTypes, true, componentFirmwareVersionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ComponentFirmwareVersion struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := ComponentFirmwareVersions().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	componentFirmwareVersionDBTypes = map[string]string{`ID`: `uuid`, `Component`: `string`, `Vendor`: `string`, `Model`: `string`, `Filename`: `string`, `Version`: `string`, `Checksum`: `string`, `UpstreamURL`: `string`, `RepositoryURL`: `string`, `CreatedAt`: `timestamptz`, `UpdatedAt`: `timestamptz`}
	_                               = bytes.MinRead
)

func testComponentFirmwareVersionsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(componentFirmwareVersionPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(componentFirmwareVersionAllColumns) == len(componentFirmwareVersionPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &ComponentFirmwareVersion{}
	if err = randomize.Struct(seed, o, componentFirmwareVersionDBTypes, true, componentFirmwareVersionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ComponentFirmwareVersion struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ComponentFirmwareVersions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, componentFirmwareVersionDBTypes, true, componentFirmwareVersionPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ComponentFirmwareVersion struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testComponentFirmwareVersionsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(componentFirmwareVersionAllColumns) == len(componentFirmwareVersionPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &ComponentFirmwareVersion{}
	if err = randomize.Struct(seed, o, componentFirmwareVersionDBTypes, true, componentFirmwareVersionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ComponentFirmwareVersion struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ComponentFirmwareVersions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, componentFirmwareVersionDBTypes, true, componentFirmwareVersionPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ComponentFirmwareVersion struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(componentFirmwareVersionAllColumns, componentFirmwareVersionPrimaryKeyColumns) {
		fields = componentFirmwareVersionAllColumns
	} else {
		fields = strmangle.SetComplement(
			componentFirmwareVersionAllColumns,
			componentFirmwareVersionPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := ComponentFirmwareVersionSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}
