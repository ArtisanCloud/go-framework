package boostrap

import (
	service "github.com/ArtisanCloud/go-framework/app/services"
	"github.com/ArtisanCloud/go-framework/cache"
	"github.com/ArtisanCloud/go-framework/config"
	"github.com/ArtisanCloud/go-framework/database"
	logger "github.com/ArtisanCloud/go-framework/loggerManager"
	"github.com/ArtisanCloud/go-framework/resources/lang"
)

func InitProject() {
	// Initialize the global config
	envConfigName := "environment"
	dbConfigName := "database"
	cacheConfigName := "cache"
	logConfigName := "log"
	config.LoadEnvConfig(nil, &envConfigName, nil)
	config.LoadDatabaseConfig(nil, &dbConfigName, nil)
	config.LoadCacheConfig(nil, &cacheConfigName, nil)
	config.LoadVersion()
	config.LoadLogConfig(nil, &logConfigName, nil)

	// load locale
	lang.LoadLanguages()

	// setup ssh key path
	service.SetupSSHKeyPath(config.AppConfigure.SSH)

	// Initialize the cache
	_ = cache.SetupCache()

	// Initialize the database
	_ = database.SetupDatabase()

	// Initialize the logger
	_ = logger.SetupLog()
}