package main

import (
	"github.com/GnauqTheBeast/Blockchain/model"
	"log"
)

func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {
	bc := model.NewBlockchain()

	bc.AddTransaction("Quang", "Ha", 10)
	bc.CreateBlock(1, bc.LastBlock().Hash())

	bc.AddTransaction("Ha", "Quang", 5)
	bc.CreateBlock(2, bc.LastBlock().Hash())

	bc.Print()
}
