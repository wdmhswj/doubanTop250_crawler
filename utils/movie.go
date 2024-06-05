package utils

import "os"

func FileDirExist(file string) (bool, error) {
	_, err := os.Stat(file)
	if err == nil {
		return true, nil
	} else if os.IsNotExist(err) {
		return false, nil
	} else {
		return false, err
	}
}
