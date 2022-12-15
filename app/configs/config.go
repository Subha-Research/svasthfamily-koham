package configs

import (
	"log"
	"path/filepath"
	"runtime"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/json"
	"github.com/knadh/koanf/providers/file"
)

func LoadConfig() map[string]interface{} {
	// Load JSON config.
	var _, b, _, _ = runtime.Caller(0)
	var basepath = filepath.Dir(b)
	var config_path = filepath.Join(basepath, "config.json")

	var k = koanf.New(".")
	if err := k.Load(file.Provider(config_path), json.Parser()); err != nil {
		log.Fatalf("error loading config: %v", err)
	}
	// fmt.Println("Database name is  = ", k.String("db.database"))
	// fmt.Println("Database port is = ", k.Int("db.port"))
	return k.All()
}
