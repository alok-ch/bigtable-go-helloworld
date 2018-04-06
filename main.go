package main

import (
	"github.com/alok-ch/bigtable-go-helloworld/app/controller"
	"github.com/alok-ch/bigtable-go-helloworld/app/models"
	"github.com/alok-ch/bigtable-go-helloworld/config"
	"github.com/alok-ch/bigtable-go-helloworld/pkg/database"
	"github.com/alok-ch/bigtable-go-helloworld/pkg/logger"
	"golang.org/x/net/context"
)

func main() {

	//initialize logger
	log := &logger.RealLogger{}
	log.Initialise()

	//initialize config

	cfgProvider := config.EnvAppConfigProvider{}
	cfgMap, cfgErr := cfgProvider.ProvideEnv(config.AppConfigList)
	if cfgErr != nil {
		panic(cfgErr)
	}

	cfg := config.ConstructAppConfig(cfgMap)

	dbConfig := database.DBConfig{
		Project:  cfg.Project,
		Instance: cfg.Instance,
	}
	ctx := context.Background()

	adminClient, admError := dbConfig.InitialiseAdminClient(ctx)

	if admError != nil {
		panic("unable to initialise admin client with error :" + admError.Error())
	}

	dbClient, dbClientError := dbConfig.InitialiseNewClient(ctx)

	if dbClientError != nil {
		panic("unable to initialise db client with error :" + dbClientError.Error())
	}

	a := &controller.App{
		Log:           log,
		Cfg:           cfg,
		Ctx:           ctx,
		AdminService:  &models.Admin{adminClient},
		ClientService: &models.Client{dbClient},
	}
	a.HelloWorld()

}
