package controller

import (
	"game/internal/models"
	"game/internal/service"
)

func InitRoutes() {
	hero := models.Hero{Name: "Янацист", Health: 100}

	hero = service.Attack(hero)
	hero = service.Attack(hero)
	hero = service.Attack(hero)
	hero = service.Attack(hero)
	hero = service.Attack(hero)
	hero = service.Attack(hero)
	
	
	
	hero = service.Heal(hero)

}
