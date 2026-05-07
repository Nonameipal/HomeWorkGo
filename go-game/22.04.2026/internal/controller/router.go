package controller

import (
	"game/internal/models"
)

func InitRoutes() {
	h := models.Hero{Name: "ЯНенацист", Health: 100}

	h.Attack()
	h.Attack()
	h.Attack()
	h.Attack()
	h.Attack()
	h.Attack()

	h.Heal()

}
