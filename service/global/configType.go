package global

type ConfigInformation struct {
	//Dbconfig
	Dbconfig DbConfig
	//Serconfig
	Serconfig SerConfig
}

// Sqlconfig
type DbConfig struct {
	DbType     string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string
}

// Serconfig
type SerConfig struct {
	ServiceMode string
	HttpPort    string
}
