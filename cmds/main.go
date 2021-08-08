package main

import "github.com/bang9211/ossicones/implements/ossiconesblockchain"

func main() {
	bc := ossiconesblockchain.ObtainBlockchain()
	bc.AddBlock("First Block")
	bc.AddBlock("Second Block")
	bc.AddBlock("Thrid Block")
	bc.PrintBlock()
}
