package parse

import (
	crossplane "github.com/nginxinc/nginx-go-crossplane"
)

func HttpParse(parsedConfig crossplane.Directives) ([]ProxyConfig, []UpstreamConfig) {
	var serverConfigs []ProxyConfig
	var upstreamConfigs []UpstreamConfig
	for _, directive := range parsedConfig {
		if directive.Directive == "server" {
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
					serverConfig := ProxyConfig{
						Listen:     listenPort,
						Location:   context,
						ProxyPass:  proxyPass,
						ServerName: serverName,
						Line:       line,
					}
					serverConfigs = append(serverConfigs, serverConfig)
				case "upstream":
					upstreamConfigs = append(upstreamConfigs, UpstreamParse(subDirective))
				}
				continue
			}
		} else if directive.Directive == "upstream" {
			upstreamConfigs = append(upstreamConfigs, UpstreamParse(directive))
		}
	}
	return serverConfigs, upstreamConfigs
}
