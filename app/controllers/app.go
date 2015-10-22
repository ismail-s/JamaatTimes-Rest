package controllers

import (
	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

type MasjidObj struct {
	Id   int
	Name string
}

type Masjid struct {
	*revel.Controller
}

func (c Masjid) Get() revel.Result {
	m := MasjidObj{Id: 1, Name: "test"}
	return c.RenderJson(m)
}
