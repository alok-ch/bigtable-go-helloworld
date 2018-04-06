package controller

import "github.com/alok-ch/bigtable-go-helloworld/app/models"

func (a *App) HelloWorld() {
	input := &models.ClientInput{}
	input.TableName = "Hello-Bigtable"
	input.ColumnFamilyName = "cf1"
	input.ColumnName = "greeting"
	input.Ctx = a.Ctx
	input.InputRow = []string{"Hello World!", "Hello Cloud Bigtable!", "Hello golang!"}

	a.performAdminCreation(input)
	a.performClientActions(input)
	a.performCleanup(input)

}

func (a *App) performAdminCreation(input *models.ClientInput) {
	err := a.AdminService.CreateTable(input)

	if err != nil {
		a.Log.Error("error while creating table :", err)
		return
	}

	err = a.AdminService.CreateColumnFamily(input)

	if err != nil {
		a.Log.Error("error while creating column family :", err)
		return
	}
}

func (a *App) performClientActions(input *models.ClientInput) {
	errorRows, err := a.ClientService.WriteRows(input)
	if err != nil {
		if err != nil {
			a.Log.Error("error while fetching rows :", err)
			return
		}
	}

	for _, val := range errorRows {
		a.Log.Error("error while creating row :", val.Error())
	}

	dataRows, err := a.ClientService.GetAllRows(input)

	for _, val := range dataRows {
		a.Log.Info("item row :", val.Row)
		a.Log.Info("item row :", string(val.Value))

	}

}

func (a *App) performCleanup(input *models.ClientInput) {
	err := a.AdminService.DeleteTable(input)

	if err != nil {
		a.Log.Error("Could not delete table  %s: %v",input.TableName,err)
		return
	}

	a.Log.Info("sucessfully deleted table : %s",input.TableName)

	err = a.AdminService.Close()
	if err != nil {
		a.Log.Error("Could not close connection   : %v",err)
		return
	}
}
