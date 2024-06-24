package parse

import (
	crossplane "github.com/nginxinc/nginx-go-crossplane"
	"nginx-config-parse/src/util"
)

func Parse(config *crossplane.Payload) ([]ProxyConfig, []UpstreamConfig) {
	var serverConfigs []ProxyConfig
	var upstreamConfigs []UpstreamConfig
	for _, parsedFile := range config.Config {
		util.Logger.Infof("parsed file %s status %s\n", parsedFile.File, util.CompareAndColorize("ok", parsedFile.Status))
		if len(parsedFile.Errors) > 0 {
			util.Logger.Error("Errors:", parsedFile.Errors)
			continue
		}
		serverConfig, upstreamConfig := ConfigParse(parsedFile.Parsed)
		for i := range serverConfig {
			serverConfig[i].ConfigFile = parsedFile.File
		}
		for i := range upstreamConfig {
			upstreamConfig[i].ConfigFile = parsedFile.File
		}
		serverConfigs = append(serverConfigs, serverConfig...)
		upstreamConfigs = append(upstreamConfigs, upstreamConfig...)
	}
	return serverConfigs, upstreamConfigs
}

func ConfigParse(parsedConfig crossplane.Directives) ([]ProxyConfig, []UpstreamConfig) {
	var serverConfigs []ProxyConfig
	var UpstreamConfigs []UpstreamConfig
	for _, directive := range parsedConfig {
		if directive.Directive == "http" {
			serverConfig, upstreamConfig := HttpParse(directive.Block)
			UpstreamConfigs = append(UpstreamConfigs, upstreamConfig...)
			serverConfigs = append(serverConfigs, serverConfig...)
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

				serverConfig := ProxyConfig{
					Listen:     listenPort,
					Location:   context,
					ProxyPass:  proxyPass,
					ServerName: serverName,
					Line:       line,
				}
				serverConfigs = append(serverConfigs, serverConfig)

			}

		} else if directive.Directive == "upstream" {
			UpstreamConfigs = append(UpstreamConfigs, UpstreamParse(directive))
		}
	}
	return serverConfigs, UpstreamConfigs
}
