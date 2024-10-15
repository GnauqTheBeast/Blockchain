package main

import (
	"fmt"
	"github.com/GnauqTheBeast/Blockchain/model"
)

func main() {
	w := model.NewWallet()
	fmt.Println(w.PublicKey())
	fmt.Println(w.PrivateKeyStr())
	fmt.Println(w.PublicKey())
	fmt.Println(w.PublicKeyStr())
	//
	//bc := model.NewBlockchain("my_blockchain_address")
	//
	//bc.AddTransaction("Quang", "Ha", 10)
	//bc.Mining()
	//
	//bc.AddTransaction("Ha", "Quang", 5)
	//bc.Mining()
	//
	//bc.Print()
	//
	//fmt.Println(bc.CalcuateTotalAmount("my_blockchain_address"))
}
