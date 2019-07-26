package parser

import (
	"fmt"
	"github.com/killtw/flapper/config"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type File struct {
	Path string
	Name string
	Dir string
	home string
}

func NewFile(path string, conf config.Configure) File {
	info, err := os.Stat(path)

	if err != nil {
		log.Fatal("Could not get file info")
	}

	return File{
		Path: path,
		Name: info.Name(),
		home: conf.Home,
	}
}

func (f *File) mkDir(name string) *File {
	if strings.HasPrefix(f.home, "~/") {
		usr, _ := os.UserHomeDir()
		f.home = strings.ReplaceAll(f.home, "~/", usr + "/")
	}

	f.Dir = filepath.Join(f.home, name)

	_ = os.MkdirAll(f.Dir, 0755)

	return f
}

func (f *File) Move(dir string, name string) (*File, error) {
	f.mkDir(dir)

	if name == "" {
		name = f.Name
	} else {
		name = name + filepath.Ext(f.Name)
	}

	if err := os.Rename(f.Path, fmt.Sprintf("%s/%s", f.Dir, name)); err != nil {
		return nil, fmt.Errorf("move failed: %s", err)
	}

	return f, nil
}
