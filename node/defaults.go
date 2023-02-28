package node

import (
	"os"
	"os/user"
	"path/filepath"
	"runtime"
)

const (
	DefaultHTTPHost    = "localhost" // Default host interface for the HTTP RPC server
	DefaultHTTPPort    = 8545        // Default TCP port for the HTTP RPC server
	DefaultWSHost      = "localhost" // Default host interface for the websocket RPC server
	DefaultWSPort      = 8546        // Default TCP port for the websocket RPC server
	DefaultGraphQLHost = "localhost" // Default host interface for the GraphQL server
	DefaultGraphQLPort = 8547        // Default TCP port for the GraphQL server
	DefaultAuthHost    = "localhost" // Default host interface for the authenticated apis
	DefaultAuthPort    = 8551        // Default port for the authenticated apis
)

var (
	DefaultAuthCors    = []string{"localhost"} // Default cors domain for the authenticated apis
	DefaultAuthVhosts  = []string{"localhost"} // Default virtual hosts for the authenticated apis
	DefaultAuthOrigins = []string{"localhost"} // Default origins for the authenticated apis
	DefaultAuthPrefix  = ""                    // Default prefix for the authenticated apis
	DefaultAuthModules = []string{"eth", "engine"}
)

// DefaultConfig contains reasonable default settings.
var DefaultConfig = Config{
	DataDir: DefaultDataDir(),
}

func DefaultDataDir() string {
	home := homeDir()
	if home != "" {
		switch runtime.GOOS {
		case "darwin":
			return filepath.Join(home, "Library", "Ethereum")
		case "windows":
			fallback := filepath.Join(home, "AppData", "Roaming", "Ethereum")
			appData := windowsAppData()
			if appData == "" || isNonEmptyDir(fallback) {
				return fallback
			}
			return filepath.Join(appData, "Ethereum")
		default:
			return filepath.Join(home, ".ethereum")
		}
	}
	return ""
}

func homeDir() string {
	if home := os.Getenv("HOME"); home != "" {
		return home
	}
	if usr, err := user.Current(); err == nil {
		return usr.HomeDir
	}
	return ""
}
func windowsAppData() string {
	v := os.Getenv("LOCALAPPDATA")
	if v == "" {
		panic("enviroment variable LocalAppData is undefined")
	}
	return v
}

func isNonEmptyDir(dir string) bool {
	f, err := os.Open(dir)
	if err != nil {
		return false
	}
	names, _ := f.Readdir(1)
	f.Close()
	return len(names) > 0
}
