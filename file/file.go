package file

import(
	"os"
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