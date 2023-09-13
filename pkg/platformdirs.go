package platformdirs

import (
  "path"
  "os"
)

type PlatformParams struct {
	appname string
	appauthor string
	version string // ""
	roaming bool // False
	multipath bool // False
	opinion bool // True
	ensure_exists bool // False
}

func newParams(appName string, appAuthor string) *PlatformParams {
	return &PlatformParams{
	  appname: appName,
	  appauthor: appAuthor,
	  version: "",
	  roaming: false,
	  multipath: false,
	  opinion: true,
	  ensure_exists: false,
	}
  }

type PlatformDirs interface {
    UserDataDir() (string, error)
	// SiteDataDir() (string, error)
	// UserConfigDir() (string, error)
	// SiteConfigDir() (string, error)
	// UserCacheDir() (string, error)
	// SiteCacheDir() (string, error)
	// UserStateDir() (string, error)
	// UserLogDir() (string, error)
	// UserDocumentsDir() (string, error)
	// UserDownloadsDir() (string, error)
	// UserPicturesDir() (string, error)
	// UserVideosDir() (string, error)
	// UserMusicDir() (string, error)
	// UserDesktopDir() (string, error)
	// UserRuntimeDir() (string, error)
	// SiteRuntimeDir() (string, error)
}

func New(appName string, appAuthor string) PlatformDirs {
	return newParams(appName, appAuthor)
}

func (r *PlatformParams) UserDataDir() (string, error) {
   return user_data_dir(r)
}



func append_app_name_and_version(dir *PlatformParams, base string) (string, error) {
	var params []string
	if dir.appname != "" {
		params = append(params, dir.appname)
		if (dir.version != "") {
			params = append(params, dir.version)
		}
	}
	params = append([]string{base}, params...)
	pathValue := path.Join(params...)
	err := optionally_create_directory(dir, pathValue)
	if err != nil {
		return "", err
	}
	return pathValue, nil
}

func optionally_create_directory(dir *PlatformParams, path string) error {
	if (dir.ensure_exists){
	  err := os.MkdirAll(path, 0755)
	  if os.IsExist(err) {
		return nil
	  }
	  return err
	}
	return nil
}

func user_data_dir(params *PlatformParams) (string, error) {
	return "", nil
}