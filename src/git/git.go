package git

import (
	"fmt"
	"path/filepath"

	"github.com/go-git/go-billy/v5"
	"github.com/go-git/go-billy/v5/memfs"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/storage/memory"
)

func CloneRepo(url string) (*git.Repository, billy.Filesystem, error) {
	fs := memfs.New()
	storer := memory.NewStorage()
	repo, err := git.Clone(storer, fs, &git.CloneOptions{
		URL: url,
	})
	return repo, fs, err
}

func ListFiles(fs billy.Filesystem, dir string) error {
	fileInfos, err := fs.ReadDir(dir)
	if err != nil {
		return err
	}

	for _, fileInfo := range fileInfos {
		filePath := filepath.Join(dir, fileInfo.Name())
		if fileInfo.IsDir() {
			err = ListFiles(fs, filePath)
			if err != nil {
				return err
			}
		} else {
			fmt.Println(filePath)
		}
	}

	return nil
}
