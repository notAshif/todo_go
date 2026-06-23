package main

import (
	"encoding/json"
	"log"
	"os"
)

type Storage[T any] struct {
	Filename string
}

// its create a simple new storage
func NewStorage[T any](filename string) *Storage[T] {
	log.Printf("New Storage is Createdddddddddddddddd!!!!!!!")
	return &Storage[T]{Filename: filename}
}

// this functions is used to save data in json storage
func (s *Storage[T]) Save(data T) error {
	fileData, err := json.MarshalIndent(data, "", "")

	if err != nil{
		log.Printf("FAILED TO SAVE DATAAAAAAA!!!!!! %s", err)
	}

	return os.WriteFile(s.Filename, fileData, 0)
}

//load the data 
func (s *Storage[T]) load(data *T) error {
	fileData, err := json.MarshalIndent(data, "" , "")
	
	if err != nil{
		log.Printf("FAILED TO LOAD DATAAAAAAA!!!!!! %s", err)
	}

	return json.Unmarshal(fileData, data)
}