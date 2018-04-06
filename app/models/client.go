package models

import (
	"cloud.google.com/go/bigtable"
	"fmt"
	"context"
)

type ClientService interface {
	WriteRows(input *ClientInput) ([]error,error)
	GetAllRows(input *ClientInput) ([]bigtable.ReadItem,error)
}

type ClientInput struct {
	TableName string
	InputRow []string
	ColumnFamilyName string
	ColumnName string
	Ctx context.Context
}



type Client struct {
	NewClient *bigtable.Client
}

func (c *Client)getTable(input *ClientInput) (*bigtable.Table) {
	return c.NewClient.Open(input.TableName)
}

func (c *Client)WriteRows(input *ClientInput) ([]error,error) {
	muts := make([]*bigtable.Mutation, len(input.InputRow))
	rowKeys := make([]string, len(input.InputRow))
	for i, greeting := range input.InputRow {
		muts[i] = bigtable.NewMutation()
		muts[i].Set(input.ColumnFamilyName, input.ColumnName, bigtable.Now(), []byte(greeting))
		rowKeys[i] = fmt.Sprintf("%s%d", input.ColumnName, i)
	}
	tbl := c.getTable(input)
	return tbl.ApplyBulk(input.Ctx, rowKeys, muts)

}

func (c *Client)GetAllRows(input *ClientInput) ([]bigtable.ReadItem,error){
	items := make([]bigtable.ReadItem,0)
	tbl := c.getTable(input)
	err := tbl.ReadRows(input.Ctx, bigtable.PrefixRange(input.ColumnName), func(row bigtable.Row) bool {
		item := row[input.ColumnFamilyName][0]
		items = append(items, item)
		return true
	}, bigtable.RowFilter(bigtable.ColumnFilter(input.ColumnName)))

	return items,err

}

