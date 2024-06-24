package parse

import (
	"fmt"
	"strings"
)

type ProxyConfig struct {
	Listen     []string `json:"listen,omitempty"`
	Location   string   `json:"location,omitempty"`
	ProxyPass  string   `json:"proxy_pass,omitempty"`
	ServerName []string `json:"server_name,omitempty"`
	ConfigFile string   `json:"config_file,omitempty"`
	Line       int      `json:"line,omitempty"`
}

type UpstreamConfig struct {
	Upstream   string         `json:"upstream,omitempty"`
	Servers    []ServerConfig `json:"servers,omitempty"`
	Line       int            `json:"line,omitempty"`
	ConfigFile string         `json:"config_file,omitempty"`
	Policy     []string       `json:"policy,omitempty"`
}

type ServerConfig struct {
	IPPort string   `json:"ip_port,omitempty"`
	Policy []string `json:"policy,omitempty"`
}

func (serverConfig ServerConfig) String() string {
	return fmt.Sprintf("%s %s", serverConfig.IPPort, strings.Join(serverConfig.Policy, ";"))
}
