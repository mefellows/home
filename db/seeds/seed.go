package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/mefellows/home/config"
)

func main() {
	log.Println("[DEBUG] Running DB seeding")
	c := config.NewConfig()
	seed, _ := ioutil.ReadFile("db/seeds/seed.sql")
	fmt.Println(string(seed))
	f := c.DB.Exec(string(seed))

	if f.Error != nil {
		log.Fatalln("[ERROR] unable to setup db seed:", f.Error)
	}
}
