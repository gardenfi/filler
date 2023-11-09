package cobictl

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/catalogfi/cobi/cobid/handlers"
	"github.com/spf13/cobra"
)

func KillService(rpcClient Client) *cobra.Command {
	var (
		service handlers.Service
	)
	var cmd = &cobra.Command{
		Use:   "kill",
		Short: "kills a running service in daemon",
		Run: func(c *cobra.Command, args []string) {
			if service != handlers.Executor && service != handlers.Autofiller && service != handlers.AutoCreator {
				cobra.CheckErr(errors.New("invalid service type"))
			}

			FillOrder := handlers.KillSerivce{
				ServiceType: service,
			}

			jsonData, err := json.Marshal(FillOrder)
			if err != nil {
				cobra.CheckErr(fmt.Errorf("failed to marshal payload: %w", err))
			}

			resp, err := rpcClient.SendPostRequest("killService", jsonData)
			if err != nil {
				cobra.CheckErr(fmt.Errorf("failed to send request: %w", err))
			}

			fmt.Println(string(resp))
		}}
	cmd.Flags().Var(&service, "service", "allowed: \"executor\", \"autofiller\", \"autocreator\"")
	cmd.MarkFlagRequired("order-id")
	return cmd
}
