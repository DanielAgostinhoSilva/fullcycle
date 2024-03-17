package main

import (
	"encoding/json"
	"github.com/DanielAgostinhoSilva/fullcycle/custom-tags/transform"
	"log"
)

type User struct {
	ID    int    `transform:"upper"`
	Name  string `transform:"upper"`
	Email string `transform:"lower"`
}

func main() {
	user := &User{ID: 1, Name: "Fulano", Email: "TESTE@tEst.com"}
	err := transform.T(user)
	if err != nil {
		panic(err)
	}
	userJson, _ := json.Marshal(*user)
	log.Println(string(userJson))
}
