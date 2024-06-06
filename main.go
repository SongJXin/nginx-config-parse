package main

import (
	"fmt"
	"github.com/nginxinc/nginx-go-crossplane"
	"log"
	"os"
	"strings"
	"text/tabwriter"
)

const (
	Reset = "\033[0m"
	Red   = "\033[31m"
	Green = "\033[32m"
)

type ServerConfig struct {
	Listen     string
	Context    string
	ProxyPass  string
	ServerName string
	ConfigFile string
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <path_to_nginx_conf>")
		return
	}

	confPath := os.Args[1]

	// 解析Nginx配置文件
	config, err := crossplane.Parse(confPath, &crossplane.ParseOptions{})
	if err != nil {
		log.Fatalf("Error parsing Nginx config file: %v", err)
	}
	// 解析每个 server 块并提取监听端口和转发目标
	var serverConfigs []ServerConfig
	for _, parsedFile := range config.Config {
		fmt.Printf("parsed file %s status %s\n", parsedFile.File, compareAndColorize("ok", parsedFile.Status))
		if len(parsedFile.Errors) > 0 {
			fmt.Println("Errors:", parsedFile.Errors)
			continue
		}
		serverConfig := parseServerBlocks(parsedFile.Parsed)
		for i := range serverConfig {
			serverConfig[i].ConfigFile = parsedFile.File
		}
		serverConfigs = append(serverConfigs, serverConfig...)
	}
	printServerConfig(serverConfigs)
}
func parseServerBlocks(parsedConfig crossplane.Directives) []ServerConfig {
	var serverConfigs []ServerConfig
	for _, directive := range parsedConfig {
		if directive.Directive == "http" || directive.Directive == "stream" {
			serverConfigs = append(serverConfigs, parseServerBlocks(directive.Block)...)
		} else {
			if directive.Directive == "server" {
				var listenPort string
				var proxyPass string
				var context string
				var serverName string
				for _, subDirective := range directive.Block {
					switch subDirective.Directive {
					case "listen":
						if len(subDirective.Args) > 0 {
							listenPort = subDirective.Args[0]
						}
					case "server_name":
						if len(subDirective.Args) > 0 {
							serverName = strings.Join(subDirective.Args, ";")
						}
					case "location":
						context = subDirective.Args[0]
						for _, locationDirective := range subDirective.Block {
							if locationDirective.Directive == "proxy_pass" {
								if len(locationDirective.Args) > 0 {
									proxyPass = locationDirective.Args[0]
								}
							}
						}
					case "proxy_pass":
						if len(subDirective.Args) > 0 {
							proxyPass = subDirective.Args[0]
						}
					}
					serverConfig := ServerConfig{
						Listen:     listenPort,
						Context:    context,
						ProxyPass:  proxyPass,
						ServerName: serverName,
					}

					serverConfigs = append(serverConfigs, serverConfig)
				}

			}
		}
	}
	return serverConfigs
}
func printServerConfig(config []ServerConfig) {
	writer := tabwriter.NewWriter(os.Stdout, 1, 0, 2, ' ', tabwriter.TabIndent)
	fmt.Fprintln(writer, "Listen\tServer Name\tContext\tProxy Pass\tConfig File")
	for _, c := range config {
		if c.ProxyPass == "" {
			continue
		}
		fmt.Fprintln(writer, c.Listen+"\t"+c.ServerName+"\t"+c.Context+"\t"+c.ProxyPass+"\t"+c.ConfigFile)
	}
	writer.Flush()
}

// compareAndColorize 比较两个字符串并返回带颜色的字符串
func compareAndColorize(str1, str2 string) string {
	if str1 == str2 {
		return fmt.Sprintf("%s%s%s", Green, str1, Reset)
	}
	return fmt.Sprintf("%s%s%s", Red, str2, Reset)
}
