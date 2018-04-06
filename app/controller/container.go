package controller

import (
	"github.com/alok-ch/bigtable-go-helloworld/app/models"
	"github.com/alok-ch/bigtable-go-helloworld/config"
	"github.com/alok-ch/bigtable-go-helloworld/pkg/logger"
	"context"
)

// App encapsulates the App environment
type App struct {
	Cfg           *config.Config
	Log           logger.ILogger
	Ctx           context.Context
	ClientService models.ClientService
	AdminService  models.AdminClientService
}


