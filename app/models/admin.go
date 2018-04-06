package models

import (
	"cloud.google.com/go/bigtable"
	"errors"
)

type AdminClientService interface {
	FetchTableList(input *ClientInput) ([]string, error)
	CreateTable(input *ClientInput) error
	TableInfo(input *ClientInput) (*bigtable.TableInfo, error)
	CreateColumnFamily(input *ClientInput) error
	DeleteTable(input *ClientInput) error
	Close() error
}

type Admin struct {
	NewAdminClient *bigtable.AdminClient
}

func (a *Admin) FetchTableList(input *ClientInput) ([]string, error) {
	return a.NewAdminClient.Tables(input.Ctx)
}

func (a *Admin) CreateTable(input *ClientInput) error {
	tables, err := a.FetchTableList(input)

	if err != nil {
		return err
	}

	if sliceContains(tables, input.TableName) {
		return errors.New("table already exist")
	}

	return a.NewAdminClient.CreateTable(input.Ctx, input.TableName)

}

func (a *Admin) TableInfo(input *ClientInput) (*bigtable.TableInfo, error) {
	return a.NewAdminClient.TableInfo(input.Ctx, input.TableName)
}

func (a *Admin) CreateColumnFamily(input *ClientInput) error {
	tblInfo, err := a.TableInfo(input)

	if err != nil {
		return err
	}

	if sliceContains(tblInfo.Families, input.ColumnFamilyName) {
		return errors.New("column family already exist")
	}
	return a.NewAdminClient.CreateColumnFamily(input.Ctx, input.TableName, input.ColumnFamilyName)
}


func (a *Admin)DeleteTable(input *ClientInput) error {
  return a.NewAdminClient.DeleteTable(input.Ctx,input.TableName)
}

func (a *Admin)Close() error {
	return a.NewAdminClient.Close()
}

func sliceContains(list []string, target string) bool {
	for _, s := range list {
		if s == target {
			return true
		}
	}
	return false
}
