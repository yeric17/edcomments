package configuration

import (
	"encoding/json"
	"log"
	"os"
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
