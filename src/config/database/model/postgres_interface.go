package model

func NewPostgresProperties(url string, port uint64, dbname string, user string, password string) PostgresPropertiesInterface {
	return &databaseProperties{
		url:      url,
		port:     port,
		dbname:   dbname,
		user:     user,
		password: password,
	}
}

type PostgresPropertiesInterface interface {
	GetUrl() string

	GetPort() uint64

	GetDBName() string

	GetUser() string

	GetPassword() string
}
