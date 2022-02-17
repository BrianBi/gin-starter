package configs

type Database struct {
	Connection string `mapstructure:"connection"`
	Mysql struct{
		Host 				string 	`mapstructure:"host"`
		Port 				string 	`mapstructure:"port"`
		Database 			string 	`mapstructure:"database"`
		Username 			string 	`mapstructure:"username"`
		Password 			string 	`mapstructure:"password"`
		Charset 			string 	`mapstructure:"charset"`
		MaxIdleConnections 	int 	`mapstructure:"max_idle_connections"`
		MaxOpenConnections 	int 	`mapstructure:"max_open_connections"`
		MaxLifeSeconds 		int 	`mapstructure:"max_life_seconds"`
	} `mapstructure:"mysql"`
	Sqlite struct{
		Database 	string 		`mapstructure:"database"`
	} `mapstructure:"sqlite"`
}