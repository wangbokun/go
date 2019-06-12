package exec

import (
	"os/exec"
	"fmt"
)


func Command(command string) (status bool){

	_, err := exec.Command("bash", "-c", command).Output()

	if err != nil {
		fmt.Printf("Failed to execute!",err)
		return false
	}

	return true
}