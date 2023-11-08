package cobi

import (
	"encoding/json"
	"fmt"

	"github.com/catalogfi/cobi/cobictl"
	"github.com/catalogfi/cobi/handlers"
	"github.com/spf13/cobra"
)

func Create(rpcClient cobictl.Client) *cobra.Command {
	var (
		account       uint32
		orderPair     string
		sendAmount    string
		receiveAmount string
	)

	var cmd = &cobra.Command{
		Use:   "create",
		Short: "Create a new order",
		Run: func(c *cobra.Command, args []string) {

			CreateOrder := handlers.RequestCreate{
				UserAccount:   account,
				OrderPair:     orderPair,
				SendAmount:    sendAmount,
				ReceiveAmount: receiveAmount,
			}

			jsonData, err := json.Marshal(CreateOrder)
			if err != nil {
				cobra.CheckErr(fmt.Errorf("failed to marshal payload: %w", err))
			}

			resp, err := rpcClient.SendPostRequest("createNewOrder", jsonData)
			if err != nil {
				cobra.CheckErr(fmt.Errorf("failed to send request: %w", err))
			}
			var OrderId uint64

			if err := json.Unmarshal(resp, &OrderId); err != nil {
				cobra.CheckErr(fmt.Errorf("failed to unmarshal response: %w", err))
			}

			fmt.Printf("successfully created order with id %d\n", OrderId)

		},
	}

	cmd.Flags().Uint32Var(&account, "account", 0, "Account to be used (default: 0)")
	cmd.Flags().StringVar(&orderPair, "order-pair", "", "User should provide the order pair")
	cmd.MarkFlagRequired("order-pair")
	cmd.Flags().StringVar(&sendAmount, "send-amount", "", "User should provide the send amount")
	cmd.MarkFlagRequired("send-amount")
	cmd.Flags().StringVar(&receiveAmount, "receive-amount", "", "User should provide the receive amount")
	cmd.MarkFlagRequired("receive-amount")
	return cmd
}
