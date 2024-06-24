package cmd

import (
	crossplane "github.com/nginxinc/nginx-go-crossplane"
	"github.com/spf13/cobra"
	"nginx-config-parse/src/output"
	"nginx-config-parse/src/util"
)

func NewDisplayCmd() *cobra.Command {
	// Add sub commands here
	displayCmd := &cobra.Command{
		Use:   "display",
		Short: "display nginx network topology in the console",
		Run: func(cmd *cobra.Command, args []string) {
			var configFilePath = "/etc/nginx/nginx.conf"
			if cmd.Flags().Changed("config") {
				configFilePath = cmd.Flag("config").Value.String()
			}
			util.Logger.Debugf("config file path: %v", configFilePath)
			config, err := crossplane.Parse(configFilePath, &crossplane.ParseOptions{})
			if err != nil {
				util.Logger.Fatalf("Error parsing Nginx config file: %v", err)
			}
			output.ConsolePrint(config)
		},
	}
	//displayCmd.Flags().StringP("config", "c", "/etc/nginx/nginx.conf", "nginx config path")
	return displayCmd
}
