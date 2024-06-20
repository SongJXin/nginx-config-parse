package main

import (
	"github.com/nginxinc/nginx-go-crossplane"
	"nginx-config-prase/src/output"
	"nginx-config-prase/src/util"
	"os"
)

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
