package cmd

import (
	"github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ngnt",
		Short: "nginx Network topology",
		Long:  "display nginx network topology in console or api",
	}
	cmd.PersistentFlags().StringP("config", "c", "/etc/nginx/nginx.conf", "nginx config path")
	cmd.AddCommand(NewDisplayCmd())
	return cmd
}
