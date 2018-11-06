package types

import(
	"os"
	"io/ioutil"
)

//判断文件目录是否存在
func IsExist(file string) bool{

	_,error	:=	os.Stat(file)

	if error != nil {
		return false
	}
	return true
}

//return true 则是dir，反之file
func DirOrFile(name string) bool{
	return os.IsDir(name)
}


// create one file
func Create(name string) (*os.File, error) {
	return os.Create(name)
}

// remove one file
func Remove(name string) error {
	return os.Remove(name)
}

// get filepath base name
func Basename(fp string) string {
	return path.Base(fp)
}

// get filepath dir name
func Dir(fp string) string {
	return path.Dir(fp)
}

// rename file name
func Rename(src string, target string) error {
	return os.Rename(src, target)
}

// delete file
func Unlink(fp string) error {
	return os.Remove(fp)
}


//获取文件夹下的list[文件/文件夹]

// Name() string       // base name of the file
// Size() int64        // length in bytes for regular files; system-dependent for others
// Mode() FileMode     // file mode bits
// ModTime() time.Time // modification time
// IsDir() bool        // abbreviation for Mode().IsDir()
// Sys() interface{}   // underlying data source (can return nil)

func Dirlist(dirName string)([]byte,error)	{
	
	content,err := ioutil.ReadDir(dirName)

	if err != nil	{
		return err
	}
	return content
}


//读取整儿文件内容返回bytes
func FileToByte(fileName string)([]byte,error)	{

	content,err := ioutil.ReadFile(fileName)

	if err != nil	{
		return err
	}
	return content
}


// list dirs under dirPath
func DirsUnder(dirPath string) ([]string, error) {
	if !IsExist(dirPath) {
		return []string{}, nil
	}

	fs, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return []string{}, err
	}

	sz := len(fs)
	if sz == 0 {
		return []string{}, nil
	}

	ret := make([]string, 0, sz)
	for i := 0; i < sz; i++ {
		if fs[i].IsDir() {
			name := fs[i].Name()
			if name != "." && name != ".." {
				ret = append(ret, name)
			}
		}
	}

	return ret, nil
}

// list files under dirPath
func FilesUnder(dirPath string) ([]string, error) {
	if !IsExist(dirPath) {
		return []string{}, nil
	}

	fs, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return []string{}, err
	}

	sz := len(fs)
	if sz == 0 {
		return []string{}, nil
	}

	ret := make([]string, 0, sz)
	for i := 0; i < sz; i++ {
		if !fs[i].IsDir() {
			ret = append(ret, fs[i].Name())
		}
	}

	return ret, nil
}

// get file modified time
func FileMTime(fp string) (int64, error) {
	f, e := os.Stat(fp)
	if e != nil {
		return 0, e
	}
	return f.ModTime().Unix(), nil
}

// get file size as how many bytes
func FileSize(fp string) (int64, error) {
	f, e := os.Stat(fp)
	if e != nil {
		return 0, e
	}
	return f.Size(), nil
}

// Search a file in paths.
// this is often used in search config file in /etc ~/
func SearchFile(filename string, paths ...string) (fullPath string, err error) {
	for _, path := range paths {
		if fullPath = filepath.Join(path, filename); IsExist(fullPath) {
			return
		}
	}
	err = fmt.Errorf("%s not found in paths", fullPath)
	return
}