package sqliteutils

type Dao interface{
	GetTableName() string
	GetIDField() string
}