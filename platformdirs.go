package platformdirs

import (
	"errors"
	"fmt"
	"runtime"

	"github.com/christopherdavenport/platformdirs-go/internal/core"
	"github.com/christopherdavenport/platformdirs-go/internal/macos"
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
	return osSwitch(r, macos.UserDataDir, unix.UserDataDir, unImplemented)
}

func (r PlatformDirs) SiteDataDir() (string, error) {
	return osSwitch(r, macos.SiteDataDir, unix.SiteDataDir, unImplemented)
}

func (r PlatformDirs) UserConfigDir() (string, error) {
	return osSwitch(r, macos.UserConfigDir, unix.UserConfigDir, unImplemented)
}

func (r PlatformDirs) SiteConfigDir() (string, error) {
	return osSwitch(r, macos.SiteConfigDir, unix.SiteConfigDir, unImplemented)
}

func (r PlatformDirs) UserCacheDir() (string, error) {
	return osSwitch(r, macos.UserCacheDir, unix.UserCacheDir, unImplemented)
}

func (r PlatformDirs) SiteCacheDir() (string, error) {
	return osSwitch(r, macos.SiteConfigDir, unix.SiteConfigDir, unImplemented)
}

func (r PlatformDirs) UserStateDir() (string, error) {
	return osSwitch(r, macos.UserStateDir, unix.UserStateDir, unImplemented)
}

func (r PlatformDirs) UserLogDir() (string, error) {
	return osSwitch(r, macos.UserLogDir, unix.UserLogDir, unImplemented)
}

func (r PlatformDirs) UserDocumentsDir() (string, error) {
	return osSwitch(r, macos.UserDocumentsDir, unix.UserDocumentsDir, unImplemented)
}

func (r PlatformDirs) UserDownloadsDir() (string, error) {
	return osSwitch(r, macos.UserDownloadsDir, unix.UserDownloadsDir, unImplemented)
}

func (r PlatformDirs) UserPicturesDir() (string, error) {
	return osSwitch(r, macos.UserPicturesDir, unix.UserPicturesDir, unImplemented)
}

func (r PlatformDirs) UserVideosDir() (string, error) {
	return osSwitch(r, macos.UserVideosDir, unix.UserVideosDir, unImplemented)
}

func (r PlatformDirs) UserMusicDir() (string, error) {
	return osSwitch(r, macos.UserMusicDir, unix.UserMusicDir, unImplemented)
}

func (r PlatformDirs) UserDesktopDir() (string, error) {
	return osSwitch(r, macos.UserDesktopDir, unix.UserDesktopDir, unImplemented)
}

func (r PlatformDirs) UserRuntimeDir() (string, error) {
	return osSwitch(r, macos.UserRuntimeDir, unix.UserRuntimeDir, unImplemented)
}

func (r PlatformDirs) SiteRuntimeDir() (string, error) {
	return osSwitch(r, macos.SiteRuntimeDir, unix.SiteRuntimeDir, unImplemented)
}

func unImplemented(core.PlatformParams) (string, error) {
	os := runtime.GOOS
	msg := fmt.Sprintf("PlatformDirs does not know how to work with GOOS %s yet", os)
	return "", errors.New(msg)
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
