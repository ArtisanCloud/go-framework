package main

import (
	service "github.com/ArtisanCloud/go-framework/app/services"
	"github.com/ArtisanCloud/go-framework/cache"
	"github.com/ArtisanCloud/go-framework/config"
	"github.com/ArtisanCloud/go-framework/database"
	logger "github.com/ArtisanCloud/go-framework/loggerManager"
	tester "github.com/ArtisanCloud/go-framework/tests"
)
import _ "github.com/ArtisanCloud/go-framework/config"
import "github.com/ArtisanCloud/go-framework/routes"

import (
	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func main() {

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

	// setup ssh key path
	service.SetupSSHKeyPath(config.AppConfigure.SSH)

	// Initialize the cache
	_ = cache.SetupCache()

	// Initialize the database
	_ = database.SetupDatabase()

	// Initialize the logger
	_ = logger.SetupLog()

	// Initialize the Logger
	tester.TestFun()

	// Router the router as the default one provided by Gin
	Router = gin.Default()

	// Initialize the routes
	routes.InitializeRoutes(Router)

	// Start serving the application
	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080®")
	err := Router.Run(config.AppConfigure.Server.Host + ":" + config.AppConfigure.Server.Port)
	if err != nil {
		logger.Error("router error:", map[string]interface{}{
			"err": err,
		})
	}

}
