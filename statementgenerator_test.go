package sqliteutils_test

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/titogeorge/sqliteutils"
)

type AllTypes struct {
	Id        int64
	Auint     uint
	Auint8    uint8
	Auint16   uint16
	Auint32   uint32
	Auint64   uint64
	ABool     bool
	AString   string
	Aint      int
	Aint8     int8
	Aint16    int16
	Aint32    int32
	Aint64    int64
	Afloat32  float32
	Afloat64  float64
	Timestamp time.Time
	InnerJson Inner
}

func (a *AllTypes) GetTableName() string {
	return "AllTypes"
}

func (a *AllTypes) GetIDField() string {
	return "Id"
}

func (a *AllTypes) AutoIncrementPK() bool {
	return true
}

type AllTypes_NPK struct {
	SomeId    int64
	Auint     uint
	Auint8    uint8
	Auint16   uint16
	Auint32   uint32
	Auint64   uint64
	ABool     bool
	AString   string
	Aint      int
	Aint8     int8
	Aint16    int16
	Aint32    int32
	Aint64    int64
	Afloat32  float32
	Afloat64  float64
	Timestamp time.Time
	InnerJson Inner
}

func (a *AllTypes_NPK) GetTableName() string {
	return "AllTypesNPK"
}

func (a *AllTypes_NPK) GetIDField() string {
	return "SomeId"
}

func (a *AllTypes_NPK) AutoIncrementPK() bool {
	return false
}

type Inner struct {
	A int
	B string
}

func (i Inner) String() string {
	b, _ := json.Marshal(i)
	return string(b)
}

func Test_affinity(t *testing.T) {

	value := sqliteutils.GetSQLiteAffinity(reflect.Uint64)
	if value != "integer" {
		t.Error("Uint64 should be integer")
	}
	value = sqliteutils.GetSQLiteAffinity(reflect.String)
	if value != "TEXT" {
		t.Error("String should be TEXT")
	}
	value = sqliteutils.GetSQLiteAffinity(reflect.Uint)
	if value != "integer" {
		t.Error("Uint should be integer")
	}
	value = sqliteutils.GetSQLiteAffinity(reflect.Float32)
	if value != "REAL" {
		t.Error("Float32 should be REAL")
	}

	value = sqliteutils.GetSQLiteAffinity(reflect.Float64)
	if value != "REAL" {
		t.Error("Float32 should be REAL")
	}

	value = sqliteutils.GetSQLiteAffinity(reflect.Struct)
	if value != "TEXT" {
		t.Error("Struct should be TEXT")
	}

}

func Test_GenerateCreateStmt(t *testing.T) {

	value := sqliteutils.GenerateCreateStmt(&AllTypes{})
	expected := "CREATE TABLE AllTypes( Id integer primary key AUTOINCREMENT, Auint integer, Auint8 integer, Auint16 integer, Auint32 integer, Auint64 integer, ABool TEXT, AString TEXT, Aint integer, Aint8 integer, Aint16 integer, Aint32 integer, Aint64 integer, Afloat32 REAL, Afloat64 REAL, Timestamp TEXT, InnerJson TEXT); "
	if value != expected {
		t.Errorf("String should be %s instead of %s", expected, value)
	}
}

func Test_GenerateInsertStmt(t *testing.T) {
	dateString := "2018-06-16T00:00:00.00Z"
	time1, err := time.Parse(time.RFC3339, dateString)
	if err != nil {
		fmt.Println("Error while parsing date :", err)
	}
	inner := &Inner{A: 10, B: "Test"}
	value := sqliteutils.GenerateInsertStmt(&AllTypes{Auint: 34, Auint8: 23, Auint16: 1234, Auint32: 1234, Auint64: 1234, ABool: true, AString: "dfgdfgfd", Aint: 1234, Aint8: 32, Aint16: 1234, Aint32: 1234, Afloat32: 1234, Afloat64: 1234, Aint64: 1234, Timestamp: time1, InnerJson: *inner})
	//t.Log(value)
	expected := "INSERT INTO AllTypes(Auint, Auint8, Auint16, Auint32, Auint64, ABool, AString, Aint, Aint8, Aint16, Aint32, Aint64, Afloat32, Afloat64, Timestamp, InnerJson) VALUES (34, 23, 1234, 1234, 1234, \"true\", \"dfgdfgfd\", 1234, 32, 1234, 1234, 1234, 1234, 1234, \"2018-06-16 00:00:00 +0000 UTC\", \"{\"A\":10,\"B\":\"Test\"}\"); "
	if value != expected {
		t.Errorf("String should be %s instead of %s", expected, value)
	}
}

func Test_GenerateInsertStmt_no_auto_inc(t *testing.T) {
	dateString := "2018-06-16T00:00:00.00Z"
	time1, err := time.Parse(time.RFC3339, dateString)
	if err != nil {
		fmt.Println("Error while parsing date :", err)
	}
	inner := &Inner{A: 10, B: "Test"}
	value := sqliteutils.GenerateInsertStmt(&AllTypes_NPK{SomeId: 10, Auint: 34, Auint8: 23, Auint16: 1234, Auint32: 1234, Auint64: 1234, ABool: true, AString: "dfgdfgfd", Aint: 1234, Aint8: 32, Aint16: 1234, Aint32: 1234, Afloat32: 1234, Afloat64: 1234, Aint64: 1234, Timestamp: time1, InnerJson: *inner})
	//t.Log(value)
	expected := "INSERT INTO AllTypesNPK(SomeId, Auint, Auint8, Auint16, Auint32, Auint64, ABool, AString, Aint, Aint8, Aint16, Aint32, Aint64, Afloat32, Afloat64, Timestamp, InnerJson) VALUES (10, 34, 23, 1234, 1234, 1234, \"true\", \"dfgdfgfd\", 1234, 32, 1234, 1234, 1234, 1234, 1234, \"2018-06-16 00:00:00 +0000 UTC\", \"{\"A\":10,\"B\":\"Test\"}\"); "
	if value != expected {
		t.Errorf("String should be %s instead of %s", expected, value)
	}
}
