package configuration

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	//Es necesario el driver para que gorm pueda usar mysql
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

//Configuration es la estructura para config.json
type Configuration struct {
	Server   string
	Port     string
	User     string
	Password string
	Database string
}

//GetConfiguration es una funcion que retorna la configuracion
func GetConfiguration() Configuration {
	var c Configuration
	file, err := os.Open("./config.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&c)
	if err != nil {
		log.Fatal(err)
	}

	return c
}

//GetConnection obtiene una conexion a la 	BD
func GetConnection() *gorm.DB {
	c := GetConfiguration()
	//"user:password@tcp(server:port)/database?charset=utf8&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		c.User,
		c.Password,
		c.Server,
		c.Port,
		c.Database,
	)
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
