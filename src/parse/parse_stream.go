package parse

import (
	crossplane "github.com/nginxinc/nginx-go-crossplane"
)

func StreamParse(parsedConfig crossplane.Directives) []ServerConfig {
	var serverConfigs []ServerConfig
	for _, directive := range parsedConfig {
		if directive.Directive == "server" {
			var listenPort []string
			var proxyPass string
			var context string
			var line int
			for _, subDirective := range directive.Block {
				switch subDirective.Directive {
				case "listen":
					listenPort = append(listenPort, ListenParse(subDirective.Args))
					continue
				case "proxy_pass":
					proxyPass, line = ProxyPassParse(subDirective)
					continue
				}
			}
			serverConfig := ServerConfig{
				Listen:    listenPort,
				Location:  context,
				ProxyPass: proxyPass,
				Line:      line,
			}
			serverConfigs = append(serverConfigs, serverConfig)

		}
	}

	return serverConfigs
}
