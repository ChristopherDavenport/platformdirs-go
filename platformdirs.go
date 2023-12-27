package platformdirs

import (
	"errors"
	"fmt"
	"runtime"

	"github.com/christopherdavenport/platformdirs-go/internal/core"
	"github.com/christopherdavenport/platformdirs-go/internal/macos"
	"github.com/christopherdavenport/platformdirs-go/internal/unix"
	"github.com/christopherdavenport/platformdirs-go/internal/windows"
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
	return osSwitch(r, macos.UserDataDir, unix.UserDataDir, windows.UserDataDir)
}

func (r PlatformDirs) SiteDataDir() (string, error) {
	return osSwitch(r, macos.SiteDataDir, unix.SiteDataDir, windows.SiteDataDir)
}

func (r PlatformDirs) UserConfigDir() (string, error) {
	return osSwitch(r, macos.UserConfigDir, unix.UserConfigDir, windows.UserConfigDir)
}

func (r PlatformDirs) SiteConfigDir() (string, error) {
	return osSwitch(r, macos.SiteConfigDir, unix.SiteConfigDir, windows.SiteConfigDir)
}

func (r PlatformDirs) UserCacheDir() (string, error) {
	return osSwitch(r, macos.UserCacheDir, unix.UserCacheDir, windows.UserCacheDir)
}

func (r PlatformDirs) SiteCacheDir() (string, error) {
	return osSwitch(r, macos.SiteConfigDir, unix.SiteConfigDir, windows.SiteCacheDir)
}

func (r PlatformDirs) UserStateDir() (string, error) {
	return osSwitch(r, macos.UserStateDir, unix.UserStateDir, windows.UserStateDir)
}

func (r PlatformDirs) UserLogDir() (string, error) {
	return osSwitch(r, macos.UserLogDir, unix.UserLogDir, windows.UserLogDir)
}

func (r PlatformDirs) UserDocumentsDir() (string, error) {
	return osSwitch(r, macos.UserDocumentsDir, unix.UserDocumentsDir, windows.UserDocumentsDir)
}

func (r PlatformDirs) UserDownloadsDir() (string, error) {
	return osSwitch(r, macos.UserDownloadsDir, unix.UserDownloadsDir, windows.UserDowloadsDir)
}

func (r PlatformDirs) UserPicturesDir() (string, error) {
	return osSwitch(r, macos.UserPicturesDir, unix.UserPicturesDir, windows.UserPicturesDir)
}

func (r PlatformDirs) UserVideosDir() (string, error) {
	return osSwitch(r, macos.UserVideosDir, unix.UserVideosDir, windows.UserVideosDir)
}

func (r PlatformDirs) UserMusicDir() (string, error) {
	return osSwitch(r, macos.UserMusicDir, unix.UserMusicDir, windows.UserMusicDir)
}

func (r PlatformDirs) UserDesktopDir() (string, error) {
	return osSwitch(r, macos.UserDesktopDir, unix.UserDesktopDir, windows.UserDesktopDir)
}

func (r PlatformDirs) UserRuntimeDir() (string, error) {
	return osSwitch(r, macos.UserRuntimeDir, unix.UserRuntimeDir, windows.UserRuntimeDir)
}

func (r PlatformDirs) SiteRuntimeDir() (string, error) {
	return osSwitch(r, macos.SiteRuntimeDir, unix.SiteRuntimeDir, windows.SiteRuntimeDir)
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
