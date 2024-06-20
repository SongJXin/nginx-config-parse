package parse

type ServerConfig struct {
	Listen     []string `json:"listen,omitempty"`
	Location   string   `json:"location,omitempty"`
	ProxyPass  string   `json:"proxy_pass,omitempty"`
	ServerName []string `json:"server_name,omitempty"`
	ConfigFile string   `json:"config_file,omitempty"`
	Line       int      `json:"line,omitempty"`
}
