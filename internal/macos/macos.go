package macos

import (
	"github.com/christopherdavenport/platformdirs-go/internal/core"
	"os"
	"path"
	// "strings"
)

// data directory tied to the user, e.g. `~/Library/Application Support/$appname/$version`
func UserDataDir(dir core.PlatformParams) (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	homePath := path.Join(homeDir, "Library", "Application Support")
	return core.Append_app_name_and_version(dir, homePath)
}

// data directory shared by users, e.g. `/Library/Application Support/$appname/$version`
func SiteDataDir(dir core.PlatformParams) (string, error) {
	base := "/Library/Application Support"
	return core.Append_app_name_and_version(dir, base)
}

// config directory tied to the user, same as `UserDataDir`
func UserConfigDir(dir core.PlatformParams) (string, error) {
	return UserDataDir(dir)
}

// config directory shared by the users, same as `SiteDataDir`
func SiteConfigDir(dir core.PlatformParams) (string, error) {
	return SiteDataDir(dir)
}

// cache directory tied to the user, e.g. `~/Library/Caches/$appname/$version`
func UserCacheDir(dir core.PlatformParams) (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	homePath := path.Join(homeDir, "Library", "Caches")
	return core.Append_app_name_and_version(dir, homePath)
}

// cache directory shared by users, e.g. `/Library/Caches/$appname/$version`
func SiteCacheDir(dir core.PlatformParams) (string, error) {
	base := "/Library/Caches"
	return core.Append_app_name_and_version(dir, base)
}

// state directory tied to the user, same as `UserDataDir`
func UserStateDir(dir core.PlatformParams) (string, error) {
	return UserDataDir(dir)
}

// log directory tied to the user, e.g. `~/Library/Logs/$appname/$version`
func UserLogDir(dir core.PlatformParams) (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	homePath := path.Join(homeDir, "Library", "Logs")
	return core.Append_app_name_and_version(dir, homePath)
}

// documents directory tied to the user, e.g. `~/Documents`
func UserDocumentsDir(dir core.PlatformParams) (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return path.Join(homeDir, "Documents"), nil
}

// downloads directory tied to the user, e.g. `~/Downloads`
func UserDownloadsDir(dir core.PlatformParams) (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return path.Join(homeDir, "Dowloads"), nil
}

// pictures directory tied to the user, e.g. “~/Pictures“
func UserPicturesDir(dir core.PlatformParams) (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return path.Join(homeDir, "Pictures"), nil

}

// videos directory tied to the user, e.g. `~/Movies`
func UserVideosDir(dir core.PlatformParams) (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return path.Join(homeDir, "Movies"), nil
}

// music directory tied to the user, e.g. “~/Music`
func UserMusicDir(dir core.PlatformParams) (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return path.Join(homeDir, "Music"), nil
}

// desktop directory tied to the user, e.g. `~/Desktop`
func UserDesktopDir(dir core.PlatformParams) (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return path.Join(homeDir, "Desktop"), nil
}

// runtime directory tied to the user, e.g. `~/Library/Caches/TemporaryItems/$appname/$version`
func UserRuntimeDir(dir core.PlatformParams) (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	homePath := path.Join(homeDir, "Library", "Caches", "TemporaryItems")
	return core.Append_app_name_and_version(dir, homePath)
}

// runtime directory shared by users, same as `user_runtime_dir`
func SiteRuntimeDir(dir core.PlatformParams) (string, error) {
	return UserRuntimeDir(dir)
}
