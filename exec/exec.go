package exec

import (
	"os/exec"
)


func Command(command string) (status bool){

	_, err := exec.Command("bash", "-c", cmd).Output()

	if err != nil {
		fmt.Printf("Failed to execute!")
		return false
	}

	return true
}