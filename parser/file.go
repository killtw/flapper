package parser

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

type File struct {
	Path string
	Name string
	Dir string
}

func NewFile(path string) File {
	info, err := os.Stat(path)

	if err != nil {
		log.Fatal("Could not get file info")
	}

	return File{
		Path: path,
		Name: info.Name(),
	}
}

func (f *File) mkDir(name string) *File {
	f.Dir = filepath.Join("/Users/killtw/Downloads", name)

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
