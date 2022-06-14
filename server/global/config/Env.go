package config

var AppInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Port    int    `json:"Port"`
}

/*
var SysEnv struct {
	LocalIP       string `json:"LocalIP"`
	MongoAddress  string `json:"MongoAddress"`
	MongoUserName string `json:"MongoUserName"`
	MongoPassword string `json:"MongoPassword"`
	RunMod        int    `json:"RunMod"` // 0 为生产环境 1 为其它运行模式
}
var AppEnv struct {
	Port int `json:"Port"`
}

func LoadSysEnv(envPath string) {
	viper.SetConfigFile(envPath)
	err := viper.ReadInConfig()
	if err != nil {
		errStr := errors.New("SysEnv 读取配置文件出错: %+v", err)
		LogErr(errStr)
		panic(errStr)
	}
	viper.Unmarshal(&SysEnv)
}

func LoadAppEnv() {
	viper.SetConfigFile(File.AppEnv)

	err := viper.ReadInConfig()
	if err != nil {
		errStr := errors.New("AppEnv 读取配置文件出错: %+v", err)
		LogErr(errStr)
		panic(errStr)
	}
	viper.Unmarshal(&AppEnv)
}
*/
