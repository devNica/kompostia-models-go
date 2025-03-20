package interfaces

type DbSchema struct {
	DatabaseName     string
	DatabasePassword string
	DatabaseUser     string
	DatabaseHost     string
	DatabasePort     string
	DatabaseSchema   string
	DatabaseSSLMode  string
}
