package models

import (
	"fmt"
	"math/rand"
)

type Hero struct {
	Name   string
	Health int
}

func (h *Hero) Attack() {
	if h.Health <= 0 {
		return
	}
	dmg := rand.Intn(35)
	h.Health -= dmg
	if h.Health < 0 {
		h.Health = 0
	}
	fmt.Printf("%s атакует на %d. хп : %d\n", h.Name, dmg, h.Health)

	if h.Health == 0 {
		fmt.Println("Это фиаско братан")
		fmt.Println("Гамер Овер")
	}
}

func (h *Hero) Heal() {
	if h.Health <= 0 {
		return
	}
	heal := rand.Intn(40)
	h.Health += heal
	if h.Health > 100 {
		h.Health = 100
		fmt.Println("Уоу бро успокойся, это тебе не котики хватит наяривать")
	}
	fmt.Printf("%s лечится на %d. хп : %d\n", h.Name, heal, h.Health)
}
