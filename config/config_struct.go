package config

type (
	Config struct {
		Database Database `yaml:"database"`
		Logger   Logger   `yaml:"logger"`
		Server   Server   `yaml:"server"`
		Redis    Redis    `yaml:"redis"`
		Auth     Auth     `yaml:"auth"`
	}
	Auth struct {
		Secretkey string `yaml:"secretkey"`
		Issue     string `yaml:"issue"`
	}
	Database struct {
		Postgresql Postgresql `yaml:"postgresql"`
	}
	Server struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	}
	Redis struct {
		Server   string `yaml:"server"`
		Password string `yaml:"password"`
		DB       int    `yaml:"db"`
		Port     string `yaml:"port"`
	}
	Postgresql struct {
		Database  string `yaml:"database"`
		Username  string `yaml:"username"`
		Password  string `yaml:"password"`
		Host      string `yaml:"host"`
		Port      int    `yaml:"port"`
		Adabter   string `yaml:"adabter"`
		Time_zone string `yaml:"time_zone"`
		Charset   string `yaml:"charset"`
	}
	Logger struct {
		Max_Age          string `yaml:"max_age"`
		Max_Size         string `yaml:"max_size"`
		Filename_Pattern string `yaml:"filename_pattern"`
		Rotation_Time    string `yaml:"rotation_time"`
		Internal_Path    string `yaml:"internal_path"`
	}
)
