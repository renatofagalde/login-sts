package model

type databaseProperties struct {
	url      string
	port     uint64
	dbname   string
	user     string
	password string
}

func (m *databaseProperties) GetUrl() string {
	return m.url
}

func (m *databaseProperties) GetPort() uint64 {
	return m.port
}

func (m *databaseProperties) GetDBName() string {
	return m.dbname
}

func (m *databaseProperties) GetUser() string {
	return m.user
}

func (m *databaseProperties) GetPassword() string {
	return m.password
}
