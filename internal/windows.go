//go:build windows
package internal

import (
	"errors"
	"fmt"
	"github.com/christopherdavenport/platformdirs-go/internal/core"

	"path/filepath"

	"golang.org/x/sys/windows/registry"
)

/*
* Get folder from the registry.
*
* This is a fallback technique at best. I'm not sure if using the registry for these guarantees us the correct answer
* for all CSIDL_* names.
*
* Since this is my only mechanism. Ouch. Lets look into something in the future
 */
func get_win_folder_from_registry(csidl_name string) (string, error) {
	shell_folder_names := map[string]string{
		"CSIDL_APPDATA":        "AppData",
		"CSIDL_COMMON_APPDATA": "Common AppData",
		"CSIDL_LOCAL_APPDATA":  "Local AppData",
		"CSIDL_PERSONAL":       "Personal",
		"CSIDL_DOWNLOADS":      "{374DE290-123F-4565-9164-39C4925E467B}",
		"CSIDL_MYPICTURES":     "My Pictures",
		"CSIDL_MYVIDEO":        "My Video",
		"CSIDL_MYMUSIC":        "My Music",
	}
	shell_folder_name := shell_folder_names[csidl_name]
	if shell_folder_name == "" {
		return "", errors.New(fmt.Sprintf("Unknown CSIDL name: %s", csidl_name))
	}
	k, err := registry.OpenKey(registry.CURRENT_USER, `Software\Microsoft\\Windows\CurrentVersion\Explorer\Shell Folders`, 0x0001)
	if err != nil {
		return "", err
	}

	value, _, err := k.GetStringValue(shell_folder_name)
	if err != nil {
		return "", err
	}

	return value, nil
}

/*
	data directory tied to the user, e.g.
		``%USERPROFILE%\\AppData\\Local\\$appauthor\\$appname`` (not roaming) or
		``%USERPROFILE%\\AppData\\Roaming\\$appauthor\\$appname`` (roaming)
*/
func UserDataDir(dir core.PlatformParams) (string, error) {
	var csidl string
	if dir.Roaming {
		csidl = "CSIDL_APPDATA"
	} else {
		csidl = "CSIDL_LOCAL_APPDATA"
	}
	folder, err := get_win_folder_from_registry(csidl)
	if err != nil {
		return "", err
	}

	return core.Append_Windows(dir, folder, "")
}

/*
data directory shared by users, e.g. ``C:\\ProgramData\\$appauthor\\$appname``"""
*/
func SiteDataDir(dir core.PlatformParams) (string, error) {
	csidl := "CSIDL_COMMON_APPDATA"
	folder, err := get_win_folder_from_registry(csidl)
	if err != nil {
		return "", err
	}

	return core.Append_Windows(dir, folder, "")
}

func UserConfigDir(dir core.PlatformParams) (string, error) {
	return UserDataDir(dir)
}

func SiteConfigDir(dir core.PlatformParams) (string, error) {
	return SiteDataDir(dir)
}

func UserCacheDir(dir core.PlatformParams) (string, error) {
	csidl := "CSIDL_LOCAL_APPDATA"
	folder, err := get_win_folder_from_registry(csidl)
	if err != nil {
		return "", err
	}

	return core.Append_Windows(dir, folder, "Cache")
}

func SiteCacheDir(dir core.PlatformParams) (string, error) {
	csidl := "CSIDL_COMMON_APPDATA"
	folder, err := get_win_folder_from_registry(csidl)
	if err != nil {
		return "", err
	}

	return core.Append_Windows(dir, folder, "Cache")
}

func UserStateDir(dir core.PlatformParams) (string, error) {
	return UserDataDir(dir)
}

/*
log directory tied to the user, same as `user_data_dir` if not opinionated else ``Logs`` in it
*/
func UserLogDir(dir core.PlatformParams) (string, error) {
	path, err := UserDataDir(dir)
	if err != nil {
		return "", err
	}
	if dir.Opinion {
		path = filepath.Join(path, "Logs")
		err = core.Optionally_create_directory(dir, path)
		if err != nil {
			return "", err
		}

	}
	return path, nil
}

/*
documents directory tied to the user e.g. `%USERPROFILE%\\Documents`
*/
func  UserDocumentsDir(dir core.PlatformParams) (string, error) {
	csidl := "CSIDL_PERSONAL"
	folder, err := get_win_folder_from_registry(csidl)
	if err != nil {
		return "", err
	}
	return filepath.Clean(folder), nil
}

/*
downloads directory tied to the user e.g. `%USERPROFILE%\\Downloads`
*/
func  UserDowloadsDir(dir core.PlatformParams) (string, error) {
	csidl := "CSIDL_DOWNLOADS"
	folder, err := get_win_folder_from_registry(csidl)
	if err != nil {
		return "", err
	}
	return filepath.Clean(folder), nil
}

/*
pictures directory tied to the user e.g. `%USERPROFILE%\\Pictures`
*/
func  UserPicturesDir(dir core.PlatformParams) (string, error) {
	csidl := "CSIDL_MYPICTURES"
	folder, err := get_win_folder_from_registry(csidl)
	if err != nil {
		return "", err
	}
	return filepath.Clean(folder), nil
}

func  UserVideosDir(dir core.PlatformParams) (string, error) {
	csidl := "CSIDL_MYVIDEO"
	folder, err := get_win_folder_from_registry(csidl)
	if err != nil {
		return "", err
	}
	return filepath.Clean(folder), nil
}

func  UserMusicDir(dir core.PlatformParams) (string, error) {
	csidl := "CSIDL_MYMUSIC"
	folder, err := get_win_folder_from_registry(csidl)
	if err != nil {
		return "", err
	}
	return filepath.Clean(folder), nil
}

func  UserDesktopDir(dir core.PlatformParams) (string, error) {
	csidl := "CSIDL_DESKTOPDIRECTORY"
	folder, err := get_win_folder_from_registry(csidl)
	if err != nil {
		return "", err
	}
	return filepath.Clean(folder), nil
}

/*
	runtime directory tied to the user, e.g.
	`%USERPROFILE%\\AppData\\Local\\Temp\\$appauthor\\$appname``
*/
func UserRuntimeDir(dir core.PlatformParams) (string, error) {
	folder, err := get_win_folder_from_registry("CSIDL_LOCAL_APPDATA")
	if err != nil {
		return "", nil
	}
	base := filepath.Join(folder, "Temp")
	return core.Append_Windows(dir, base, "")
}

func SiteRuntimeDir(dir core.PlatformParams) (string, error) {
	return UserRuntimeDir(dir)
}
