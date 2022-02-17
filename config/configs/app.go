package configs

type App struct {
	Name 		string 	`mapstructure:"name"`
	Debug 		bool 	`mapstructure:"debug"`
	Url 		string 	`mapstructure:"url"`
	Port 		string 	`mapstructure:"port"`
	Timezone 	string 	`mapstructure:"timezone"`
	Locale 		string 	`mapstructure:"locale"`
}