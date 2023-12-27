package platformdirs

import (
	"github.com/christopherdavenport/platformdirs-go/internal/core"
	"github.com/christopherdavenport/platformdirs-go/internal"
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
//   UserDataDir() (string, error)
// 	 SiteDataDir() (string, error)
// 	 UserConfigDir() (string, error)
// 	SiteConfigDir() (string, error)
// 	UserCacheDir() (string, error)
//  SiteCacheDir() (string, error)
// 	UserStateDir() (string, error)
// 	UserLogDir() (string, error)
// 	UserDocumentsDir() (string, error)
// 	UserDownloadsDir() (string, error)
// 	UserPicturesDir() (string, error)
// 	UserVideosDir() (string, error)
// 	UserMusicDir() (string, error)
// 	UserDesktopDir() (string, error)
// 	UserRuntimeDir() (string, error)
// 	SiteRuntimeDir() (string, error)
// }

func (r PlatformDirs) UserDataDir() (string, error) {
	return internal.UserDataDir(r.transform())
}

func (r PlatformDirs) SiteDataDir() (string, error) {
	return internal.SiteDataDir(r.transform())
}

func (r PlatformDirs) UserConfigDir() (string, error) {
	return internal.UserConfigDir(r.transform())
}

func (r PlatformDirs) SiteConfigDir() (string, error) {
	return internal.SiteConfigDir(r.transform())
}

func (r PlatformDirs) UserCacheDir() (string, error) {
	return internal.UserCacheDir(r.transform())
}

func (r PlatformDirs) SiteCacheDir() (string, error) {
	return internal.SiteCacheDir(r.transform())
}

func (r PlatformDirs) UserStateDir() (string, error) {
	return internal.UserStateDir(r.transform())
}

func (r PlatformDirs) UserLogDir() (string, error) {
	return internal.UserLogDir(r.transform())
}

func (r PlatformDirs) UserDocumentsDir() (string, error) {
	return internal.UserDocumentsDir(r.transform())
}

func (r PlatformDirs) UserDownloadsDir() (string, error) {
	return internal.UserDowloadsDir(r.transform())
}

func (r PlatformDirs) UserPicturesDir() (string, error) {
	return internal.UserPicturesDir(r.transform())
}

func (r PlatformDirs) UserVideosDir() (string, error) {
	return internal.UserVideosDir(r.transform())
}

func (r PlatformDirs) UserMusicDir() (string, error) {
	return internal.UserMusicDir(r.transform())
}

func (r PlatformDirs) UserDesktopDir() (string, error) {
	return internal.UserDesktopDir(r.transform())
}

func (r PlatformDirs) UserRuntimeDir() (string, error) {
	return internal.UserRuntimeDir(r.transform())
}

func (r PlatformDirs) SiteRuntimeDir() (string, error) {
	return internal.SiteRuntimeDir(r.transform())
}
