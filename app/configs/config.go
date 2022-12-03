package configs

import (
	"log"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/json"
	"github.com/knadh/koanf/providers/file"
)

func LoadConfig() map[string]interface{} {
	// Load JSON config.
	var k = koanf.New(".")
	if err := k.Load(file.Provider("app/configs/config.json"), json.Parser()); err != nil {
		log.Fatalf("error loading config: %v", err)
	}
	// fmt.Println("Database name is  = ", k.String("db.database"))
	// fmt.Println("Database port is = ", k.Int("db.port"))
	return k.All()
}
