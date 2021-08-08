package blockchain

type Blockchain interface {
	AddBlock(data string)
	PrintBlock()
}
