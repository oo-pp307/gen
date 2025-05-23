// Code generated by github.com/oo-pp307/gen. DO NOT EDIT.
// Code generated by github.com/oo-pp307/gen. DO NOT EDIT.
// Code generated by github.com/oo-pp307/gen. DO NOT EDIT.

package query

import (
	"context"
	"fmt"
	"testing"

	"github.com/oo-pp307/gen"
	"github.com/oo-pp307/gen/field"
	"github.com/oo-pp307/gen/tests/.gen/dal_4/model"
	"gorm.io/gorm/clause"
)

func init() {
	InitializeDB()
	err := _gen_test_db.AutoMigrate(&model.Person{})
	if err != nil {
		fmt.Printf("Error: AutoMigrate(&model.Person{}) fail: %s", err)
	}
}

func Test_personQuery(t *testing.T) {
	person := newPerson(_gen_test_db)
	person = *person.As(person.TableName())
	_do := person.WithContext(context.Background()).Debug()

	primaryKey := field.NewString(person.TableName(), clause.PrimaryKey)
	_, err := _do.Unscoped().Where(primaryKey.IsNotNull()).Delete()
	if err != nil {
		t.Error("clean table <people> fail:", err)
		return
	}

	_, ok := person.GetFieldByName("")
	if ok {
		t.Error("GetFieldByName(\"\") from person success")
	}

	err = _do.Create(&model.Person{})
	if err != nil {
		t.Error("create item in table <people> fail:", err)
	}

	err = _do.Save(&model.Person{})
	if err != nil {
		t.Error("create item in table <people> fail:", err)
	}

	err = _do.CreateInBatches([]*model.Person{{}, {}}, 10)
	if err != nil {
		t.Error("create item in table <people> fail:", err)
	}

	_, err = _do.Select(person.ALL).Take()
	if err != nil {
		t.Error("Take() on table <people> fail:", err)
	}

	_, err = _do.First()
	if err != nil {
		t.Error("First() on table <people> fail:", err)
	}

	_, err = _do.Last()
	if err != nil {
		t.Error("First() on table <people> fail:", err)
	}

	_, err = _do.Where(primaryKey.IsNotNull()).FindInBatch(10, func(tx gen.Dao, batch int) error { return nil })
	if err != nil {
		t.Error("FindInBatch() on table <people> fail:", err)
	}

	err = _do.Where(primaryKey.IsNotNull()).FindInBatches(&[]*model.Person{}, 10, func(tx gen.Dao, batch int) error { return nil })
	if err != nil {
		t.Error("FindInBatches() on table <people> fail:", err)
	}

	_, err = _do.Select(person.ALL).Where(primaryKey.IsNotNull()).Order(primaryKey.Desc()).Find()
	if err != nil {
		t.Error("Find() on table <people> fail:", err)
	}

	_, err = _do.Distinct(primaryKey).Take()
	if err != nil {
		t.Error("select Distinct() on table <people> fail:", err)
	}

	_, err = _do.Select(person.ALL).Omit(primaryKey).Take()
	if err != nil {
		t.Error("Omit() on table <people> fail:", err)
	}

	_, err = _do.Group(primaryKey).Find()
	if err != nil {
		t.Error("Group() on table <people> fail:", err)
	}

	_, err = _do.Scopes(func(dao gen.Dao) gen.Dao { return dao.Where(primaryKey.IsNotNull()) }).Find()
	if err != nil {
		t.Error("Scopes() on table <people> fail:", err)
	}

	_, _, err = _do.FindByPage(0, 1)
	if err != nil {
		t.Error("FindByPage() on table <people> fail:", err)
	}

	_, err = _do.ScanByPage(&model.Person{}, 0, 1)
	if err != nil {
		t.Error("ScanByPage() on table <people> fail:", err)
	}

	_, err = _do.Attrs(primaryKey).Assign(primaryKey).FirstOrInit()
	if err != nil {
		t.Error("FirstOrInit() on table <people> fail:", err)
	}

	_, err = _do.Attrs(primaryKey).Assign(primaryKey).FirstOrCreate()
	if err != nil {
		t.Error("FirstOrCreate() on table <people> fail:", err)
	}

	var _a _another
	var _aPK = field.NewString(_a.TableName(), "id")

	err = _do.Join(&_a, primaryKey.EqCol(_aPK)).Scan(map[string]interface{}{})
	if err != nil {
		t.Error("Join() on table <people> fail:", err)
	}

	err = _do.LeftJoin(&_a, primaryKey.EqCol(_aPK)).Scan(map[string]interface{}{})
	if err != nil {
		t.Error("LeftJoin() on table <people> fail:", err)
	}

	_, err = _do.Not().Or().Clauses().Take()
	if err != nil {
		t.Error("Not/Or/Clauses on table <people> fail:", err)
	}
}
