package main

import "github.com/bang9211/ossicones/implements/ossiconesblockchain"

func main() {
	bc := ossiconesblockchain.NewBlockchain()
	bc.AddBlock("Genesis Block")
	bc.AddBlock("Second Block")
	bc.AddBlock("Thrid Block")
	bc.PrintBlock()
}
