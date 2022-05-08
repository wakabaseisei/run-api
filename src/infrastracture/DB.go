package infrastructure

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type DB struct {
	Host       string
	Username   string
	Password   string
	DBName     string
	Connection *gorm.DB
}

func NewDB() *DB {
	c := NewConfig()
	return newDB(&DB{
		Host:     c.DB.Production.Host,
		Username: c.DB.Production.Username,
		Password: c.DB.Production.Password,
		DBName:   c.DB.Production.DBName,
	})
}

func NewDevDB() *DB {
	c := NewConfig()
	return newDB(&DB{
		Host:     c.DB.Dev.Host,
		Username: c.DB.Dev.Username,
		Password: c.DB.Dev.Password,
		DBName:   c.DB.Dev.DBName,
	})
}

func NewRunDB() *DB {
	c := NewConfig()
	return newDB(&DB{
		Host:     c.DB.Run.Host,
		Username: c.DB.Run.Username,
		Password: c.DB.Run.Password,
		DBName:   c.DB.Run.DBName,
	})
}

func newDB(d *DB) *DB {
	// https://github.com/go-sql-driver/mysql#examples
	db, err := gorm.Open("mysql", d.Username+":"+d.Password+"@tcp("+d.Host+")/"+d.DBName+"?charset=utf8&parseTime=True&loc=Local")
	// dbURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", d.Username, d.Password, d.Host, "3306", d.DBName)
	// db, err := gorm.Open("mysql", dbURI)
	if err != nil {
		panic(err.Error())
	}
	d.Connection = db
	return d
}

// Begin begins a transaction
func (db *DB) Begin() *gorm.DB {
	return db.Connection.Begin()
}

func (db *DB) Connect() *gorm.DB {
	return db.Connection
}
