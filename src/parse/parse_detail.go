package parse

import crossplane "github.com/nginxinc/nginx-go-crossplane"

func ListenParse(args []string) string {
	if len(args) > 0 {
		return args[0]
	}
	return ""
}

func ServerNameParse(args []string) []string {
	if len(args) > 0 {
		return args
	}
	return nil
}

func LocationParse(Directive *crossplane.Directive) (string, string, int) {
	context := Directive.Args[0]
	var proxyPass string
	var line int
	for _, locationDirective := range Directive.Block {
		if locationDirective.Directive == "proxy_pass" {
			proxyPass, line = ProxyPassParse(locationDirective)
		}
	}
	return context, proxyPass, line
}

func ProxyPassParse(locationDirective *crossplane.Directive) (string, int) {
	if len(locationDirective.Args) > 0 {
		return locationDirective.Args[0], locationDirective.Line
	}
	return "", 0
}
