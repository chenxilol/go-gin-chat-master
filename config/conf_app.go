package config

type App struct {
	Port           string `mapstructure:"port" json:"port" yaml:"port"`
	UploadFilePath string ` mapstructure:"upload_File_Path" json:"upload_File_Path" yaml:"upload_File_Path"`
	CookieKey      string `mapstructure:"cookie_key" json:"cookie_key" yaml:"cookie_key"`
	ServeType      string `mapstructure:"serve_type" json:"serve_type" yaml:"serve_type"`
	DebugMod       string `mapstructure:"debug_mod" json:"debug_mod" yaml:"debug_mod"`
}
