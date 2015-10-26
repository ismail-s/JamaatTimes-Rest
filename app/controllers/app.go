package controllers

import (
	"github.com/revel/revel"
	"github.com/ismail-s/jamaattimes-rest/app/models"
)

type MasjidApp struct {
	GormController
}

func (c MasjidApp) Create() revel.Result {
	masjid := models.Masjid{Name: "Test"}
    c.Txn.NewRecord(masjid)
    c.Txn.Create(&masjid)
    return c.RenderJson(masjid)
}

func (c MasjidApp) Get() revel.Result {
	var masjids []models.Masjid
	c.Txn.Find(&masjids)
    return c.RenderJson(masjids)
}