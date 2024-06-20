package main

import (
	"github.com/nginxinc/nginx-go-crossplane"
	"github.com/spf13/cobra"
	"nginx-config-parse/src/output"
	"nginx-config-parse/src/util"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "nginx-config-parse",
	Short: "My application does wonderful things",
}

func main() {
	if len(os.Args) != 2 {
		util.Logger.Error("Usage: go run main.go <path_to_nginx_conf>")
		return
	}

	confPath := os.Args[1]

	// 解析Nginx配置文件
	config, err := crossplane.Parse(confPath, &crossplane.ParseOptions{})
	if err != nil {
		util.Logger.Fatalf("Error parsing Nginx config file: %v", err)
	}

	output.ConsolePrintServerConfig(config)
}
