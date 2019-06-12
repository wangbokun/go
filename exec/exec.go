package exec

import (
	"os/exec"
	"strings"
	"fmt"
)


func Command(command string) (status bool){
	args := strings.Split(command, " ")
	_, err := exec.Command(args[0], args[1:]...).Output()

	if err != nil {
		fmt.Printf("Failed to execute!",err)
		return false
	}

	return true
}