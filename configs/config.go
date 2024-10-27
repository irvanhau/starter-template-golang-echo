package configs

type ProgramConfig struct {
	Server    int
	DBPort    int
	DBHost    string
	DBUser    string
	DBPass    string
	DBName    string
	Secret    string
	RefSecret string
}
