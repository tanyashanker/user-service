package models

type DBConfig struct {
	MongoSecrets *Secret
	RedisSecrets *Secret
}

type Secret struct {
	Host     string
	UserName string
	Password string
	DBName   string
	Port     int
}
