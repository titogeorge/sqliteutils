# sqliteutils
SQLite statement generator

[![Go Report Card](https://goreportcard.com/badge/github.com/titogeorge/sqliteutils)](https://goreportcard.com/report/github.com/titogeorge/sqliteutils)

[![Build Status](https://travis-ci.org/titogeorge/sqliteutils.svg?branch=master)](https://travis-ci.org/titogeorge/sqliteutils)

[![codecov](https://codecov.io/gh/titogeorge/sqliteutils/branch/master/graph/badge.svg)](https://codecov.io/gh/titogeorge/sqliteutils)


```go
type ATable struct {
	ID     int64
	Type   string
	AValue string
}

func (a *ATable) GetTableName() string {
	return "ATable"
}

func (a *ATable) GetIDFields() []string {
	return []string{"ID", "Type"}
}

func (a *ATable) AutoIncrementPK() bool {
	return false
}

sqliteutils.GenerateCreateStmt(&ATable{})
// CREATE TABLE IF NOT EXISTS MuitiPK( ID integer, Type TEXT, AValue TEXT, PRIMARY KEY(ID, Type)); 

sqliteutils.GenerateInsertStmt(&ATable{ID: 10, Type: "Test", AValue: "ATest"})
//INSERT INTO MuitiPK(ID, Type, AValue) VALUES (10, 'Test', 'ATest');

```


