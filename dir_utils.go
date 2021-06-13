package gofnd

import (
	"fmt"
	"os"
)

func CreateDirs(dirs []string) (err error) {
	for _, dir := range dirs {
		if err = os.MkdirAll(dir, os.ModePerm); err != nil {
			return fmt.Errorf("gofnd: [.CreateDirs][1]")
		}
	}

	return
}
