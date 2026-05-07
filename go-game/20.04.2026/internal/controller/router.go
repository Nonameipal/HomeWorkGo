package controller

import (
	"game/internal/models"
	"game/internal/service"
)

func InitRoutes() {
	hero := models.Hero{Name: "ЯНенацист", Health: 100}

	service.Attack(&hero)
	service.Attack(&hero)
	service.Attack(&hero)
	service.Attack(&hero)
	service.Attack(&hero)
	service.Attack(&hero)
	
	
	
	service.Heal(&hero)

}
