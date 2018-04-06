package database

import (
	"cloud.google.com/go/bigtable"
	"golang.org/x/net/context"
)

type ConnectionInitialiser interface {
	InitialiseAdminClient(ctx context.Context) (*bigtable.AdminClient, error)
	InitialiseNewClient(ctx context.Context)(*bigtable.Client,error)
}

type DBConfig struct {
	Project  string `json:"project"`
	Instance string `json:"instance"`
}

func (d *DBConfig) InitialiseAdminClient(ctx context.Context) (*bigtable.AdminClient, error) {
	return bigtable.NewAdminClient(ctx, d.Project, d.Instance)
}

func (d *DBConfig)InitialiseNewClient(ctx context.Context) (*bigtable.Client,error) {
	return bigtable.NewClient(ctx,d.Project, d.Instance)
}



