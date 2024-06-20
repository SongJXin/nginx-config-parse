package output

import (
	"fmt"
	crossplane "github.com/nginxinc/nginx-go-crossplane"
	"nginx-config-prase/src/parse"
	"nginx-config-prase/src/util"
	"os"
	"strings"
	"text/tabwriter"
)

func ConsolePrintServerConfig(config *crossplane.Payload) {
	// 解析每个 server 块并提取监听端口和转发目标
	var serverConfigs []parse.ServerConfig
	for _, parsedFile := range config.Config {
		util.Logger.Infof("parsed file %s status %s\n", parsedFile.File, util.CompareAndColorize("ok", parsedFile.Status))
		if len(parsedFile.Errors) > 0 {
			util.Logger.Error("Errors:", parsedFile.Errors)
			continue
		}
		serverConfig := parse.ConfigParse(parsedFile.Parsed)
		for i := range serverConfig {
			serverConfig[i].ConfigFile = parsedFile.File
		}
		serverConfigs = append(serverConfigs, serverConfig...)
	}
	writer := tabwriter.NewWriter(os.Stdout, 1, 0, 2, ' ', tabwriter.TabIndent)
	_, err := fmt.Fprintln(writer, "Listen\tLocation\tContext\tProxyPass\tConfigFile")
	if err != nil {
		util.Logger.Error("Print server config failed:", err)
		return
	}
	for _, c := range serverConfigs {
		if c.ProxyPass == "" {
			continue
		}
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
}
