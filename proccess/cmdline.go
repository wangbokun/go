package process

import(
	"fmt"
	"github/wangbokun/go/file"
)

func CmdLine(pid int )(cmdLine string,err error){

	fileName := fmt.Sprintf("/proc/%s/cmdline",pid)

	if !file.IsExist(fileName) {
		
		content,err	:=	file.FileToByte(fileName)

		if err != nil {
			return err
		}
		return content
	}
}