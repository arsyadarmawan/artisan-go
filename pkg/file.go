package pkg

import "os"

func CheckFolderIsExist(domainPath string) error {
	if _, err := os.Stat(domainPath); err != nil {
		if os.IsNotExist(err) {
			return err
		}
	}
	return nil
}
