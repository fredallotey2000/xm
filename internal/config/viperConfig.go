package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type Configuration struct {
	HttpServerConfig
	MysqlConfig
	KafkaConfig
	LoggerConfig
}

func NewViperConfig(cfgFile string) Configuration {
	setupConfigFile(cfgFile)
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		log.Println("Error while reading config file", err)
	}
	config := getConfiguration()
	return config
}

func setupConfigFile(cfgFile string) {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.SetConfigName("config")
		viper.AddConfigPath("./")
		viper.AddConfigPath("../")
		viper.AddConfigPath("../../")
		viper.SetConfigType("yaml")
	}
}

func getConfiguration() Configuration {
	return Configuration{
		HttpServerConfig{
			ServerIp: getStringOrPanic("http-server-ip"),
		},
		MysqlConfig{
			ConnStr: getStringOrPanic("mysql-connstr"),
		},
		KafkaConfig{
			ServerIp:        getStringOrPanic("kafka-server-ip"),
			GroupId:         getStringOrPanic("kafka-group-id"),
			AutoOffsetReset: getStringOrPanic("kafka-auto-offset-reset"),
			Topic:           getStringOrPanic("kafka-topic"),
		},
		LoggerConfig{
			PathUrl:      getStringOrPanic("logs-path"),
			LogExtension: getStringOrPanic("logs-extension"),
		},
	}
}

func panicIfError(err error) {
	if err != nil {
		panic(fmt.Errorf("unable to load config: %v", err))
	}
}

func checkKey(key string) {
	if !viper.IsSet(key) {
		panicIfError(fmt.Errorf("%s key is not set", key))
	}
}

func getStringOrPanic(key string) string {
	checkKey(key)
	v := viper.GetString(key)
	return v
}

// func getBoolOrPanic(key string) bool {
// 	checkKey(key)
// 	v := viper.GetBool(key)
// 	return v
// }

// func value2string(v interface{}) (ret string, errstr string) {

// 	switch x := v.(type) {
// 	case bool:
// 		if x {
// 			ret = "true"
// 		} else {
// 			ret = "false"
// 		}
// 	case int:
// 		ret = fmt.Sprintf("%d", x)
// 	case string:
// 		ret = x
// 	case fmt.Stringer:
// 		ret = x.String()
// 	default:
// 		return "", fmt.Sprintf("Invalid value type %T", v)
// 	}

// 	return ret, ""
// }
