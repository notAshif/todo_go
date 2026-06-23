package main

import (
	"fmt"
	"log"
)

func main(){
	fmt.Println("---------------WELCOME TO TODO CLI---------------")
	todos := Todos{}

	//load data into json storage
	storage := NewStorage[Todos]("todos.json")
	err := storage.load(&todos);
	if err != nil{
		log.Panicf("WARNINGU+fe0f COULD NOT LOAD TODO'S FROM STORAGE. START FRESH TODO... %s", err)
	}

	//parse and excute cmd
	cmd := newCommand()
	cmd.Execute(&todos)

	//save data
	err = storage.Save(todos);
	if err != nil{
		log.Println("FAILED TO SAVE DATA (TODO'S)....")
	}

}