package services

import (
	"log"
)

func PingHandler(name string, pass string) error {
	log.Println(name)
	log.Println(pass)

	return nil
}
