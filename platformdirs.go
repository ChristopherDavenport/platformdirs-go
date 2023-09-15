package platformdirs

import (
	"errors"
	"runtime"

	"github.com/christopherdavenport/platformdirs-go/internal/core"
	"github.com/christopherdavenport/platformdirs-go/internal/darwin"
	"github.com/christopherdavenport/platformdirs-go/internal/unix"
)

type PlatformDirs struct {
	AppName      string
	AppAuthor    string
	Version      string // ""
	Roaming      bool   // False
	Multipath    bool   // False
	Opinion      bool   // True
	EnsureExists bool   // False
}

func New(appName string, appAuthor string) PlatformDirs {
	return PlatformDirs{
		AppName:      appName,
		AppAuthor:    appAuthor,
		Version:      "",
		Roaming:      false,
		Multipath:    false,
		Opinion:      true,
		EnsureExists: false,
	}
}

func (r PlatformDirs) transform() core.PlatformParams {
	return core.PlatformParams{
		AppName:      r.AppName,
		AppAuthor:    r.AppAuthor,
		Version:      r.Version,
		Roaming:      r.Roaming,
		Multipath:    r.Multipath,
		Opinion:      r.Opinion,
		EnsureExists: r.EnsureExists,
	}
}

// type PlatformDirs interface {
//     UserDataDir() (string, error)
// 	// SiteDataDir() (string, error)
// 	// UserConfigDir() (string, error)
// 	// SiteConfigDir() (string, error)
// 	// UserCacheDir() (string, error)
// 	// SiteCacheDir() (string, error)
// 	// UserStateDir() (string, error)
// 	// UserLogDir() (string, error)
// 	// UserDocumentsDir() (string, error)
// 	// UserDownloadsDir() (string, error)
// 	// UserPicturesDir() (string, error)
// 	// UserVideosDir() (string, error)
// 	// UserMusicDir() (string, error)
// 	// UserDesktopDir() (string, error)
// 	// UserRuntimeDir() (string, error)
// 	// SiteRuntimeDir() (string, error)
// }

func (r PlatformDirs) UserDataDir() (string, error) {
	return osSwitch(r, darwin.UserDataDir, unix.UserDataDir, unImplemented)
}

func unImplemented(core.PlatformParams) (string, error) {
	return "", errors.New("PlatformDirs does not know how to work with that GOOS yet")
}

func osSwitch(
	dirs PlatformDirs,
	mac func(core.PlatformParams) (string, error),
	unix func(core.PlatformParams) (string, error),
	windows func(core.PlatformParams) (string, error),
) (string, error) {
	t := dirs.transform()
	switch os := runtime.GOOS; os {
	case "darwin", "ios":
		return mac(t)
	case "linux", "freebsd", "openbsd", "netbsd", "solaris", "aix", "dragonfly", "illumos", "plan9":
		return unix(t)
	case "windows":
		return windows(t)
	default: // Android, and js???
		return unImplemented(t)
	}
}
