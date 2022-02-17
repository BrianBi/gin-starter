package configs

type Log struct {
	Level				string	`mapstructure:"level"`
	Type				string	`mapstructure:"type"`
	Filename			string	`mapstructure:"filename"`
	MaxSize				int		`mapstructure:"max_size"`
	MaxBackup			int		`mapstructure:"max_backup"`
	MaxAge				int		`mapstructure:"max_age"`
	Compress			bool	`mapstructure:"compress"`
	TimeLayout			string	`mapstructure:"time_layout"`
}
