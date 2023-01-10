package shell

import (
	"log"
	"os"
	"path/filepath"

	"github.com/otiai10/copy"
)

func IsExist(path string) bool {
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		return true
	}
	return false
}

func IsDir(path string) bool {
	src, err := os.Stat(path)
	if err != nil {
		return false
	}
	return src.IsDir()
}

func Join(elm ...string) string {
	return filepath.Join(elm...)
}

func Mkdir(path string) {
	os.MkdirAll(path, os.ModePerm)
}

func Move(source, target string) {
	if IsDir(target) {
		target = Join(target, filepath.Base(source))
	}

	if GlobalMakeStyle {
		log.Println("move", source, target)
	}

	if err := os.Rename(source, target); err != nil {
		log.Fatal(err)
	}
}

func Copy(source, target string) {
	if IsDir(target) {
		target = Join(target, filepath.Base(source))
	}

	if GlobalMakeStyle {
		log.Println("copy", source, target)
	}

	if err := copy.Copy(source, target); err != nil {
		log.Fatal(err)
	}
}

func Remove(paths ...string) {
	for _, path := range paths {
		if GlobalMakeStyle {
			log.Println("remove", path)
		}
		os.RemoveAll(path)
	}
}
