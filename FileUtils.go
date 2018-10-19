package utils

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

// 判断文件夹是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

//创建文件夹
func PathMkdir(path string) error {
	exist, err := PathExists(path)
	if err != nil {
		return err
	}
	if !exist {
		// 创建文件夹
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}

//获取文件后缀名
func GetFileExt(fileName string) string {
	if len(fileName) > 0 {
		return filepath.Ext(fileName)
	} else {
		return ""
	}
}

//获取目录下的文件（可以通过后缀过滤）
func GetAllFile(filePath string, filterExt string, notContain bool) ([]os.FileInfo, error) {
	rd, err := ioutil.ReadDir(filePath)
	if err == nil {
		var files = make([]os.FileInfo, 0)
		for _, f := range rd {
			if f.IsDir() == false {
				if len(filterExt) > 0 {
					if notContain {
						if GetFileExt(f.Name()) != filterExt {
							files = append(files, f)
						}
					} else {
						if GetFileExt(f.Name()) == filterExt {
							files = append(files, f)
						}
					}
				} else {
					files = append(files, f)
				}
			}
		}
		return files, nil
	} else {
		return nil, err
	}

}

//获取文件的MD5
func GetFileMd5(filePath string) string {
	md5File, _ := os.Open(filePath)
	md5h := md5.New()
	defer md5File.Close()
	io.Copy(md5h, md5File)
	return hex.EncodeToString(md5h.Sum([]byte("")))
}
