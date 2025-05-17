package chain

type ChainClient interface {
	NewChainClient() *ChainClient
	Connect()
	Disconnect()
	Subscribe(filter any) error
	Unsubscribe() error
}
