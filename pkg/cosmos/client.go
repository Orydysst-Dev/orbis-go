package cosmos

import (
	"context"
	"fmt"
	"os/user"
	"strings"

	"github.com/ignite/cli/v28/ignite/pkg/cosmosaccount"
	"github.com/ignite/cli/v28/ignite/pkg/cosmosclient"
	"github.com/sourcenetwork/orbis-go/config"
)

type Client struct {
	cosmosclient.Client
	Account cosmosaccount.Account
	Address string
	// RpcClient *rpcclient.WSClient
}

func New(ctx context.Context, cfg config.Cosmos) (*Client, error) {
	opts := []cosmosclient.Option{
		cosmosclient.WithNodeAddress(cfg.RPCAddress),
		cosmosclient.WithAddressPrefix(cfg.AddressPrefix),
		cosmosclient.WithFees(cfg.Fees),
		cosmosclient.WithKeyringBackend(cosmosaccount.KeyringBackend(cfg.KeyringBackend)), // TODO
	}
	home := cfg.Home
	if home != "" {
		if strings.HasPrefix(home, "~/") {
			user, err := user.Current()
			if err != nil {
				return nil, fmt.Errorf("couldn't resolve user home path: %w", err)
			}
			home = strings.Replace(home, "~", user.HomeDir, 1)
		}
		opts = append(opts, cosmosclient.WithHome(home))
	}

	client, err := cosmosclient.New(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("new cosmos client: %w", err)
	}

	account, err := client.Account(cfg.AccountName)
	if err != nil {
		return nil, fmt.Errorf("get account by name: %w", err)
	}

	address, err := account.Address(cfg.AddressPrefix)
	if err != nil {
		return nil, fmt.Errorf("get account address: %w", err)
	}

	// rpcClient, err := rpcclient.NewWS(cfg.RPCAddress, "/websocket")
	// if err != nil {
	// 	return nil, fmt.Errorf("new rpc client: %w", err)
	// }

	err = client.RPC.Start()
	if err != nil {
		return nil, fmt.Errorf("rpc client start: %w", err)
	}

	return &Client{
		Client:  client,
		Account: account,
		Address: address,
		// RpcClient: rpcClient,
	}, nil
}
