package config

var RuntimeConf = RuntimeConfig{}

type RuntimeConfig struct {
	Datasource Datasource `yaml:"datasource"`
	Server     Server     `yaml:"server"`
	Jwt        Jwt        `yaml:"jwt"`
}

type Datasource struct {
	DbType       string `yaml:"dbType"`
	Url          string `yaml:"url"`
	Host         string `yaml:"host"`
	Port         string `yaml:"port"`
	UserName     string `yaml:"userName"`
	Password     string `yaml:"password"`
	DatabaseName string `yaml:"databaseName"`
}

type Server struct {
	Port    int    `yaml:"port"`
	Profile string `yaml:"profile"`
}

type Jwt struct {
	Secret string `yaml:"secret"`
}
