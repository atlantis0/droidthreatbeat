// Config is put into a different package to prevent cyclic imports in case
// it is needed in several locations

package config

import "time"

type Config struct {
	Period         time.Duration `config:"period"`
	ThreatFilePath string        `config:"threat_file"`
}

var DefaultConfig = Config{
	Period:         30 * time.Second,
	ThreatFilePath: "threats.json",
}
