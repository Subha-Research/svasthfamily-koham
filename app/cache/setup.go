package cache

import (
	"net"
	"net/url"

	"github.com/Subha-Research/svasthfamily-koham/app/configs"
	redis "github.com/go-redis/redis/v9"
)

type Redis struct {
}

func (r *Redis) getRedisConnectionVariables() (string, string) {
	config := configs.LoadConfig()
	port := config["redis.port"].(string)
	host := config["redis.host"].(string)
	address := net.JoinHostPort(host, port)
	password := url.QueryEscape(config["redis.password"].(string))

	return address, password
}

func (r *Redis) buildACLRedisOptions() redis.Options {
	config := configs.LoadConfig()
	db := int(config["redis.acl_db"].(float64))
	address, password := r.getRedisConnectionVariables()
	opts := &redis.Options{
		Addr:     address,
		Password: password,
		DB:       db,
	}
	return *opts
}

func (r *Redis) buildTokenRedisOptions() redis.Options {
	config := configs.LoadConfig()
	db := int(config["redis.token_db"].(float64))

	address, password := r.getRedisConnectionVariables()
	opts := &redis.Options{
		Addr:     address,
		Password: password,
		DB:       db,
	}
	return *opts
}

func (r *Redis) SetupACLRedisDB() redis.Client {
	opts := r.buildACLRedisOptions()
	client := redis.NewClient(&opts)
	return *client
}

func (r *Redis) SetupTokenRedisDB() redis.Client {
	opts := r.buildACLRedisOptions()
	client := redis.NewClient(&opts)
	return *client
}
