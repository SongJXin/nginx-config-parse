package parse

import crossplane "github.com/nginxinc/nginx-go-crossplane"

func UpstreamParse(Directive *crossplane.Directive) UpstreamConfig {
	upstream := Directive.Args[0]
	line := Directive.Line
	var policy []string
	var serverConfigs []ServerConfig
	for _, serverDirective := range Directive.Block {
		if serverDirective.Directive != "server" {
			policy = append(policy, serverDirective.Directive)
		} else {
			var serverConfig ServerConfig
			serverConfig.IPPort = serverDirective.Args[0]
			serverConfig.Policy = serverDirective.Args[1:]
			serverConfigs = append(serverConfigs, serverConfig)
		}
	}
	return UpstreamConfig{
		Upstream: upstream,
		Line:     line,
		Policy:   policy,
		Servers:  serverConfigs,
	}
}
