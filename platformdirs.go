package platformdirs

import (
	"errors"
	"runtime"

	"github.com/christopherdavenport/platformdirs-go/internal/core"
	"github.com/christopherdavenport/platformdirs-go/internal/unix"
	"github.com/christopherdavenport/platformdirs-go/internal/darwin"
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
	switch os := runtime.GOOS; os {
	case "darwin":
		return darwin.UserDataDir(r.transform())
	case "linux":
		return unix.UserDataDir(r.transform())
	default:
		return "", errors.New("PlatformDirs does not know how to work with that GOOS yet")
	}
}
