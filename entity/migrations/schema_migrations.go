package migrations

type SchemaMigrations struct {
	Version int
	Dirty   bool
}

func (SchemaMigrations) TableName() string {
	return "public.schema_migrations"
}
