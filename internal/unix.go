//go:build !(darwin || ios || windows)
package internal
// linux || freebsd || openbsd || netbsd || solaris || aix || dragonfly || illumos || plan9
import (
	"errors"
	"fmt"
	"os"
	"path"
	"runtime"
	"strings"

	"github.com/christopherdavenport/platformdirs-go/internal/core"
)

func customGetuid() (int, error) {
	uid := os.Getuid()
	if uid == -1 {
		return 0, errors.New("Should only be used on Unix")
	}
	return uid, nil
}

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

// Multipath string in form of `xyx(pathListSeperator)`
func withMultiPath(dir core.PlatformParams, multiPathString string) (string, error) {
	pathSepString := string(os.PathListSeparator)
	paths := strings.Split(multiPathString, pathSepString)
	if !dir.Multipath {
		paths = paths[0:1]
	}

	for i := 0; i < len(paths); i++ {
		thisPath := paths[i]
		newPath, err := core.Append_app_name_and_version(dir, thisPath)
		if err != nil {
			return "", err
		}
		paths[i] = newPath
	}
	return strings.Join(paths, pathSepString), nil
}

// data directories shared by users, if multipath enabled will return
// multiple by the OS Path List Seperator.
// e.g. `/usr/local/share/$appname/$version` or `/usr/share/$appname/$version`
// multipath - `/usr/local/share/$appname/$version:/usr/share/$appname/$version`
func SiteDataDir(dir core.PlatformParams) (string, error) {
	dataDirs := os.Getenv("XDG_DATA_DIRS")
	if dataDirs == "" || strings.Trim(dataDirs, " ") == "" {
		dataDirs = fmt.Sprintf("/usr/local/share%s/user/share", string(os.PathListSeparator))
	}
	return withMultiPath(dir, dataDirs)
}

// config directory tied to the user, e.g. `~/.config/$appname/$version` or
// `$XDG_CONFIG_HOME/$appname/$version`
func UserConfigDir(params core.PlatformParams) (string, error) {
	configHome := os.Getenv("XDG_CONFIG_HOME")
	if configHome == "" {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		configHome = path.Join(homeDir, ".config")
	}
	return core.Append_app_name_and_version(params, configHome)
}

//	config directories shared by users
//
// Without multipath - /etc/xdg/$appname/$version
// With multipath - all dirs appended with `/$appname/$version`
func SiteConfigDir(params core.PlatformParams) (string, error) {
	configDirs := os.Getenv("XDG_CONFIG_DIRS")
	if configDirs == "" {
		configDirs = "/etc/xdg"
	}
	return withMultiPath(params, configDirs)
}

func envNameOrHomePlus(env string, paths ...string) (string, error) {
	out := os.Getenv(env)
	if out == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		slice := append([]string{home}, paths...)
		out = path.Join(slice...)
	}
	return out, nil
}

func UserCacheDir(params core.PlatformParams) (string, error) {
	base, err := envNameOrHomePlus("XDG_CACHE_HOME", ".cache")
	if err != nil {
		return "", err
	}
	return core.Append_app_name_and_version(params, base)
}

func SiteCacheDir(params core.PlatformParams) (string, error) {
	return core.Append_app_name_and_version(params, "/var/tmp")
}

func UserStateDir(params core.PlatformParams) (string, error) {
	base, err := envNameOrHomePlus("XDG_STATE_HOME", ".local", "state")
	if err != nil {
		return "", err
	}
	return core.Append_app_name_and_version(params, base)
}

func UserLogDir(params core.PlatformParams) (string, error) {
	base, err := UserStateDir(params)
	if err != nil {
		return "", err
	}
	if params.Opinion {
		base = path.Join(base, "log")
		err = core.Optionally_create_directory(params, base)
		if err != nil {
			return "", err
		}
	}
	return base, nil
}

func UserDocumentsDir(params core.PlatformParams) (string, error) {
	return envNameOrHomePlus("XDG_DOCUMENTS_DIR", "Documents")
}

func UserDowloadsDir(params core.PlatformParams) (string, error) {
	return envNameOrHomePlus("XDG_DOWNLOAD_DIR", "Downloads")
}

func UserPicturesDir(params core.PlatformParams) (string, error) {
	return envNameOrHomePlus("XDG_PICTURES_DIR", "Pictures")
}

func UserVideosDir(params core.PlatformParams) (string, error) {
	return envNameOrHomePlus("XDG_VIDEOS_DIR", "Videos")
}

func UserMusicDir(params core.PlatformParams) (string, error) {
	return envNameOrHomePlus("XDG_MUSIC_DIR", "Music")
}

func UserDesktopDir(params core.PlatformParams) (string, error) {
	return envNameOrHomePlus("XDG_DESKTOP_DIR", "Desktop")
}

func UserRuntimeDir(params core.PlatformParams) (string, error) {
	base := os.Getenv("XDG_RUNTIME_DIR")
	if base == "" {
		uuid, err := customGetuid()
		if err != nil {
			return "", err
		}
		switch runtimeos := runtime.GOOS; runtimeos {
		case "freebsd", "netbsd", "openbsd":

			base = fmt.Sprintf("/var/run/user/%d", uuid)
			_, err := os.Stat(base)
			if err == nil {
				break
			}
			if os.IsNotExist(err) {
				base = fmt.Sprintf("/tmp/runtime-%d", uuid)
				break
			}
			return "", err // Error accessing

		default:
			base = fmt.Sprintf("/run/user/%d", uuid)

		}
	}
	return core.Append_app_name_and_version(params, base)
}

func SiteRuntimeDir(params core.PlatformParams) (string, error) {
	base := os.Getenv("XDG_RUNTIME_DIR")
	if base == "" {
		switch runtimeos := runtime.GOOS; runtimeos {
		case "freebsd", "netbsd", "openbsd":
			base = "/var/run"
		default:
			base = "/run"
		}
	}
	return core.Append_app_name_and_version(params, base)
}
