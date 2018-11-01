package file

import (
	"io/ioutil"
)


//读取整儿文件内容返回bytes
func FileToByte(fileName string)([]byte,error)	{

	content,err := ioutil.ReadFile(fileName)

	if err != nil	{
		return err
	}
	return content
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
