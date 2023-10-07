package helpers

import (
	"fmt"
	"os"
)

func CreateFolderIfNotExists(path string) error {

	const permissions = 0755

	fi, err := os.Stat(path)
	if err != nil && !os.IsNotExist(err) {
		return err
	}

	switch {
	case os.IsNotExist(err):
		return os.MkdirAll(path, permissions)
	case !fi.IsDir():
		return fmt.Errorf("a file named %s is already existing", path)
	default:
		return nil
	}
}

func CreateFileIfNotExists(path string, data []byte) error {
	const permissions = 0755

	fi, err := os.Stat(path)

	if err != nil && os.IsNotExist(err) {
		file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, permissions)
		if err != nil {
			return err
		}

		defer file.Close()

		if len(data) > 0 {
			_, err = file.Write(data)
		}

		if err != nil {
			return err
		}

		return nil
	}

	if fi.IsDir() {
		return fmt.Errorf("a folder named %s is already existing", path)
	}

	return err
}
