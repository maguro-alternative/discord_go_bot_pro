package table

type PGTable struct {
	SchemaName string `db:"schemaname"`
	TableName  string `db:"tablename"`
	TableOwner string `db:"tableowner"`
}
