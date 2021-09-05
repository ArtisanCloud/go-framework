package cache

import (
	"github.com/ArtisanCloud/go-framework/config"
	"github.com/ArtisanCloud/go-libs/cache"
	"github.com/ArtisanCloud/go-libs/object"
)

var (
	CacheConnection *cache.GRedis

)

func SetupCache() (err error) {

	c := config.CacheConn
	//fmt.Dump(c)

	options := cache.RedisOptions{
		Host:       c.Host,
		Password:   c.Password,
		DB:         c.DB,
		SSLEnabled: c.SSLEnabled,
	}

	CacheConnection = cache.NewGRedis(&options)
	//fmt2.Printf("CacheConnection:%+v \r\n", CacheConnection.Pool.String())

	//CacheMapLockers = make(map[string]*sync.Mutex)

	return nil

}

func GetKeyPrefix() string {
	strAppName := object.Snake(config.AppConfigure.Name, "_")
	return strAppName + "_database_" + strAppName + "_cache:"
}
