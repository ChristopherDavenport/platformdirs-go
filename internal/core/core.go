package core

import (
	"os"
	"path"
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
