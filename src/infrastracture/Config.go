package infrastructure

import "os"

type Config struct {
	DB struct {
		Production struct {
			Host     string
			Username string
			Password string
			DBName   string
		}
		Dev struct {
			Host     string
			Username string
			Password string
			DBName   string
		}
		Run struct {
			Host     string
			Username string
			Password string
			DBName   string
			Port     string
		}
	}
	Routing struct {
		Port string
	}
}

func NewConfig() *Config {

	c := new(Config)

	password := os.Getenv("CLOUD_SQL_PASSWORD")
	host := os.Getenv("CLOUD_SQL_HOST")

	// GKE用
	c.DB.Production.Host = "mysql-0.mysql"
	c.DB.Production.Username = "go_user"
	c.DB.Production.Password = password
	c.DB.Production.DBName = "go_database"

	// docker-compose用
	c.DB.Dev.Host = "mysql"
	c.DB.Dev.Username = "go_user"
	c.DB.Dev.Password = "password"
	c.DB.Dev.DBName = "go_database"

	// Cloud Run用
	c.DB.Run.Host = host
	c.DB.Run.Username = "go_user"
	c.DB.Run.Password = password
	c.DB.Run.DBName = "go_database"

	c.Routing.Port = ":8080"

	return c
}
