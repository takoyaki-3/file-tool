package filetool

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
	"path/filepath"
)

type FileInfo struct {
	Hash  string `csv:"hash"`
	Path  string `csv:"path"`
	Name  string `csv:"name"`
	Size  int64  `csv:"size"`
	IsDir bool   `csv:"is_dir"`
}

func FileMd5(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}

	hashInBytes := hash.Sum(nil)
	return hex.EncodeToString(hashInBytes), nil
}

type DirWalkOption struct {
	SkipDirList   []string
	IsIncludeHash bool
}

func DirWalk(dirPath string, option DirWalkOption) (error, []FileInfo) {

	fileInfos := []FileInfo{}

	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// ディレクトリをスキップするか否か
		if info.IsDir() {
			for _, skipDir := range option.SkipDirList {
				if skipDir == info.Name() {
					return filepath.SkipDir
				}
			}
		}

		// ハッシュ値の処理
		hash := ""
		if option.IsIncludeHash {
			hash, err = FileMd5(path)
		}
		fileInfos = append(fileInfos, FileInfo{
			Hash:  hash,
			Path:  path,
			Size:  info.Size(),
			IsDir: info.IsDir(),
			Name:  info.Name(),
		})
		return nil
	})

	return err, fileInfos
}
