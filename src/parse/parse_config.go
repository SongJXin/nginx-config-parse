package parse

import (
	crossplane "github.com/nginxinc/nginx-go-crossplane"
)

func ConfigParse(parsedConfig crossplane.Directives) []ServerConfig {
	var serverConfigs []ServerConfig
	for _, directive := range parsedConfig {
		if directive.Directive == "http" {
			serverConfigs = append(serverConfigs, HttpParse(directive.Block)...)
		} else if directive.Directive == "stream" {
			serverConfigs = append(serverConfigs, StreamParse(directive.Block)...)
		} else if directive.Directive == "server" { //针对单独文件的情况
			var listenPort []string
			var proxyPass string
			var context string
			var serverName []string
			var line int
			for _, subDirective := range directive.Block {
				switch subDirective.Directive {
				case "listen":
					listenPort = append(listenPort, ListenParse(subDirective.Args))
					continue
				case "server_name":
					serverName = append(serverName, ServerNameParse(subDirective.Args)...)
					continue
				case "location":
					context, proxyPass, line = LocationParse(subDirective)
				case "proxy_pass":
					proxyPass, line = ProxyPassParse(subDirective)
				}

				serverConfig := ServerConfig{
					Listen:     listenPort,
					Location:   context,
					ProxyPass:  proxyPass,
					ServerName: serverName,
					Line:       line,
				}
				serverConfigs = append(serverConfigs, serverConfig)

			}

		}
	}
	return serverConfigs
}
