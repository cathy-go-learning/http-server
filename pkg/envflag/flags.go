package envflag

import (
	"flag"
	"os"
)

var (
	Version  = flag.String("Version", "0.0.1", "")
	LogLevel = flag.String("log-level", "info", "Logging level. env var: `LOGLEVEL`")
)

func getStringEnv(envKey string, defaultValue string) string {
	val := os.Getenv(envKey)
	if val == "" {
		return defaultValue
	}

	return val
}

func Parse() {
	flag.Parse()

	*Version = getStringEnv("VERSION", *Version)
	*LogLevel = getStringEnv("LOGLEVEL", *LogLevel)
}
