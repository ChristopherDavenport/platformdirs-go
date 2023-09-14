package darwin

import (
	"github.com/christopherdavenport/platformdirs-go/internal/core"
	"os"
	"path"
	// "strings"
)

func UserDataDir(dir core.PlatformParams) (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	homePath := path.Join(homeDir, "Library", "Application Support")
	return core.Append_app_name_and_version(dir, homePath)
}


