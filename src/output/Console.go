package output

import (
	"fmt"
	crossplane "github.com/nginxinc/nginx-go-crossplane"
	"nginx-config-parse/src/parse"
	"nginx-config-parse/src/util"
	"os"
	"strings"
	"text/tabwriter"
)

func ConsolePrint(config *crossplane.Payload) {
	serverConfigs, upstreamConfigs := parse.Parse(config)
	ConsolePrintProxyConfig(serverConfigs)
	ConsolePrintUpstreamConfig(upstreamConfigs)
}

func ConsolePrintProxyConfig(proxyConfigs []parse.ProxyConfig) {
	fmt.Println()
	fmt.Println("Server config:")
	// 解析每个 server 块并提取监听端口和转发目标

	writer := tabwriter.NewWriter(os.Stdout, 1, 0, 2, ' ', tabwriter.TabIndent)
	_, err := fmt.Fprintln(writer, "Listen\tServerName\tLocation\tProxyPass\tConfigFile")
	if err != nil {
		util.Logger.Error("Print server config failed:", err)
		return
	}
	for _, c := range proxyConfigs {
		_, err := fmt.Fprintln(
			writer,
			strings.Join(c.Listen, " ")+"\t"+
				strings.Join(c.ServerName, " ")+"\t"+
				c.Location+"\t"+
				c.ProxyPass+"\t"+
				fmt.Sprintf("%s:%d", c.ConfigFile, c.Line))
		if err != nil {
			util.Logger.Error("Print server config failed:", err)
			return
		}
	}
	err = writer.Flush()
	if err != nil {
		util.Logger.Error("Print server config failed:", err)
		return
	}
	fmt.Println()
}
func ConsolePrintUpstreamConfig(upstreamConfig []parse.UpstreamConfig) {
	fmt.Println()
	fmt.Println("Upstream config:")
	writer := tabwriter.NewWriter(os.Stdout, 1, 0, 2, ' ', tabwriter.TabIndent)
	_, err := fmt.Fprintln(writer, "Upstream\tPolicy\tServices\tConfigFile")
	if err != nil {
		util.Logger.Error("Print server config failed:", err)
		return
	}
	for _, c := range upstreamConfig {
		for i, s := range c.Servers {
			if i == 0 {
				_, err := fmt.Fprintln(
					writer,
					c.Upstream+"\t"+
						strings.Join(c.Policy, ";")+"\t"+
						s.String()+"\t"+
						fmt.Sprintf("%s:%d", c.ConfigFile, c.Line))
				if err != nil {
					util.Logger.Error("Print server config failed:", err)
					return
				}
			} else {
				_, err := fmt.Fprintln(
					writer,
					" \t \t"+s.String()+"\t\t")
				if err != nil {
					util.Logger.Error("Print server config failed:", err)
					return
				}

			}

		}
		err = writer.Flush()
		if err != nil {
			util.Logger.Error("Print server config failed:", err)
			return
		}
	}
	fmt.Println()
}
