package sqliteutils

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

//GenerateCreateStmt generate the sqlite create statement for given struct
func GenerateCreateStmt(dao Dao) string {
	var sBuilder strings.Builder
	sBuilder.WriteString("CREATE TABLE ")
	sBuilder.WriteString(dao.GetTableName())
	sBuilder.WriteString("( ")
	val := reflect.ValueOf(dao).Elem()
	for i := 0; i < val.NumField(); i++ {
		sBuilder.WriteString(val.Type().Field(i).Name)
		sBuilder.WriteString(" ")
		sBuilder.WriteString(GetSQLiteAffinity(val.Type().Field(i).Type.Kind()))
		if dao.GetIDField() == val.Type().Field(i).Name {
			sBuilder.WriteString(" primary key")
			if dao.AutoIncrementPK() {
				sBuilder.WriteString(" AUTOINCREMENT")
			}
		}

		if i != val.NumField()-1 {
			sBuilder.WriteString(", ")
		}
	}
	sBuilder.WriteString("); ")
	return sBuilder.String()
}

//GetSQLiteAffinity returns SQLITE type for go type
func GetSQLiteAffinity(fType reflect.Kind) string {
	switch fType {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return "integer"
	case reflect.Float32, reflect.Float64:
		return "REAL"
	default:
		return "TEXT"
	}
}

//GenerateInsertStmt create the insert statement for the given struct
func GenerateInsertStmt(dao Dao) string {
	var sBuilder strings.Builder
	sBuilder.WriteString("INSERT INTO ")
	sBuilder.WriteString(dao.GetTableName())
	sBuilder.WriteString("(")
	val := reflect.ValueOf(dao).Elem()

	var fields []reflect.StructField
	fields = make([]reflect.StructField, val.NumField(), val.NumField())
	for i := 0; i < val.NumField(); i++ {
		fields[i] = val.Type().Field(i)

	}
	for c, sf := range fields {
		if dao.AutoIncrementPK() && dao.GetIDField() == sf.Name {
			continue
		}
		sBuilder.WriteString(sf.Name)
		if c != val.NumField()-1 {
			sBuilder.WriteString(", ")
		}
	}
	sBuilder.WriteString(") VALUES (")
	for c, sf := range fields {
		if "Id" == sf.Name {
			continue
		}
		sBuilder.WriteString(getStringForValue(sf, val))
		if c != val.NumField()-1 {
			sBuilder.WriteString(", ")
		}
	}
	sBuilder.WriteString("); ")
	return sBuilder.String()
}

func getStringForValue(sf reflect.StructField, val reflect.Value) (value string) {

	switch sf.Type.Kind() {
	case reflect.String:
		value += "\""
		value += reflect.Indirect(val).FieldByName(sf.Name).Interface().(string)
		value += "\""
		break
	case reflect.Int:
		value = strconv.FormatInt(int64(reflect.Indirect(val).FieldByName(sf.Name).Interface().(int)), 10)
		break
	case reflect.Int8:
		value = strconv.FormatInt(int64(reflect.Indirect(val).FieldByName(sf.Name).Interface().(int8)), 10)
		break
	case reflect.Int16:
		value = strconv.FormatInt(int64(reflect.Indirect(val).FieldByName(sf.Name).Interface().(int16)), 10)
		break
	case reflect.Int32:
		value = strconv.FormatInt(int64(reflect.Indirect(val).FieldByName(sf.Name).Interface().(int32)), 10)
		break
	case reflect.Int64:
		value = strconv.FormatInt(reflect.Indirect(val).FieldByName(sf.Name).Interface().(int64), 10)
		break
	case reflect.Uint:
		value = strconv.FormatUint(uint64(reflect.Indirect(val).FieldByName(sf.Name).Interface().(uint)), 10)
		break
	case reflect.Uint8:
		value = strconv.FormatUint(uint64(reflect.Indirect(val).FieldByName(sf.Name).Interface().(uint8)), 10)
		break
	case reflect.Uint16:
		value = strconv.FormatUint(uint64(reflect.Indirect(val).FieldByName(sf.Name).Interface().(uint16)), 10)
		break
	case reflect.Uint32:
		value = strconv.FormatUint(uint64(reflect.Indirect(val).FieldByName(sf.Name).Interface().(uint32)), 10)
		break
	case reflect.Uint64:
		value = strconv.FormatUint(reflect.Indirect(val).FieldByName(sf.Name).Interface().(uint64), 10)
		break
	case reflect.Float32:
		value = strconv.FormatFloat(float64(reflect.Indirect(val).FieldByName(sf.Name).Interface().(float32)), 'f', -1, 32)
		break
	case reflect.Float64:
		value = strconv.FormatFloat(reflect.Indirect(val).FieldByName(sf.Name).Interface().(float64), 'f', -1, 32)
		break
	case reflect.Bool:
		value += "\""
		value += strconv.FormatBool(reflect.Indirect(val).FieldByName(sf.Name).Interface().(bool))
		value += "\""
		break
	case reflect.Struct:
		value += "\""
		value += fmt.Sprintf("%s", reflect.Indirect(val).FieldByName(sf.Name).Interface())
		value += "\""
		break

	default:

		fmt.Printf("Invalid Data Type Name: %s, Kind: %s, Type: %s, value: %s", sf.Name, sf.Type.Kind(), sf.Type, reflect.Indirect(val).FieldByName(sf.Name))
	}
	return
}
