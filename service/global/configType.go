package global

// 字段记得大写，供外部访问
type ConfigInformation struct {
	//Dbconfig
	Dbconfig DbConfig
	//Serconfig
	Serconfig SerConfig
	// redisConfig
	Redisconfig RedisConfig

	Emailconfig EmailConfig
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

// RedisConfig
type RedisConfig struct {
	DbHost   string
	DbName   string
	Password string
	Port     string
	Dbnum    int
}

// Email
type EmailConfig struct {
	User string
	Pass string
	Host string
	Port int
}
