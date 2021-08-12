//+build wireinject

package modules

import (
	"github.com/bang9211/ossicones/implements/ossiconesblockchain"
	"github.com/bang9211/ossicones/interfaces/blockchain"
	"github.com/google/wire"
)

var MySet = wire.NewSet(
	wire.InterfaceValue(
		new(blockchain.Blockchain),
		ossiconesblockchain.ObtainBlockchain(),
	)
)

// var MySet = wire.NewSet(wire.InterfaceValue(new(io.Reader), os.Stdin))

func initBlockchain() (blockchain.Blockchain, error) {
	wire.Build(MySet)
	return nil, nil
}
