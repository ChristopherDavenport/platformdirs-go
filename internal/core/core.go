package core

import (
	"os"
	"path"
	"path/filepath"
)

type PlatformParams struct {
	AppName      string
	AppAuthor    string
	Version      string // ""
	Roaming      bool   // False
	Multipath    bool   // False
	Opinion      bool   // True
	EnsureExists bool   // False
}

// Windows Likes the App Author Included
func Append_Windows(dir PlatformParams, base string, opinionValue string) (string, error) {
	var params []string

	if dir.AppName != "" {
		if dir.AppAuthor != "" {
			params = append(params, dir.AppAuthor)
		} else {
			params = append(params, dir.AppName)
		}
		params = append(params, dir.AppName)
		if dir.Opinion && opinionValue != "" {
			params = append(params, opinionValue)
		}

		if dir.Version != "" {
			params = append(params, dir.Version)
		}
	}
	params = append([]string{base}, params...)
	pathValue := filepath.Clean(path.Join(params...))
	err := Optionally_create_directory(dir, pathValue)
	if err != nil {
		return "", err
	}
	return pathValue, nil
}

func Append_app_name_and_version(dir PlatformParams, base string) (string, error) {
	var params []string
	if dir.AppName != "" {
		params = append(params, dir.AppName)
		if dir.Version != "" {
			params = append(params, dir.Version)
		}
	}
	params = append([]string{base}, params...)
	pathValue := path.Join(params...)
	err := Optionally_create_directory(dir, pathValue)
	if err != nil {
		return "", err
	}
	return pathValue, nil
}

func Optionally_create_directory(dir PlatformParams, path string) error {
	if dir.EnsureExists {
		err := os.MkdirAll(path, 0755)
		if os.IsExist(err) {
			return nil
		}
		return err
	}
	return nil
}
