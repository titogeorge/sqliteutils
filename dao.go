package sqliteutils

//Dao interface need to be implemented for all statement generation utils
type Dao interface {
	GetTableName() string
	GetIDField() string
}
