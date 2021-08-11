package blockchain

type Block interface {
	CalculateHash()
}

type Blockchain interface {
	AddBlock(data string)
	AllBlocks() []interface{}
	PrintBlock()
}
