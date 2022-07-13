package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"sync"
)

var instance interface{}
var once sync.Once

func GetConfig(config interface{}, configPath string) interface{} {
	once.Do(func() {
		log.Println("Reading config")
		instance = config

		if err := cleanenv.ReadConfig(configPath, instance); err != nil {
			var helpText = "Error while reading config"
			description, _ := cleanenv.GetDescription(instance, &helpText)

			log.Println(description)
			log.Fatal(err)
		}

	})

	return instance
}
