package config

type HttpServerConfig struct {
	ServerIp        string
	IdleTimeout     int
	ReadTimeout     int
	WriteTimeout    int
	ShutDownTimeout int
}

type MysqlConfig struct {
	ConnStr string
}

type KafkaConfig struct {
	ServerIp        string
	GroupId         string
	AutoOffsetReset string
	Topic           string
}

type LoggerConfig struct {
	PathUrl      string
	LogExtension string
}

type DefaulConfig struct {
	JwtSecret string
}

// type mysqlConfig struct {
// 	env             string
// 	logLevel        string
// 	port            string
// 	IdleTimeout     int
// 	ReadTimeout     int
// 	WriteTimeout    int
// 	ShutDownTimeout int
// 	RequestTimeout  int
// 	MaxJsonSize     int
// }
