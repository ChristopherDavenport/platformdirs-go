package linux

import (
	"github.com/christopherdavenport/platformdirs-go/internal/core"
	"os"
	"path"
	"strings"
)

/*
*
data directory tied to the user, e.g. “~/.local/share/$appname/$version“ or

	``$XDG_DATA_HOME/$appname/$version``
*/
func UserDataDir(dir core.PlatformParams) (string, error) {
	homePath := os.Getenv("XDG_DATA_HOME")
	if homePath == "" || strings.Trim(homePath, " ") == "" {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		homePath = path.Join(homeDir, ".local", "share")
	}
	return core.Append_app_name_and_version(dir, homePath)
}
