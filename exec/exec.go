package exec

import (
	"os/exec"
	"strings"
	"fmt"
    "bytes"
    "time"
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

func Cmd(name string, arg ...string) *exec.Cmd {
    return exec.Command(name, arg...)
}

func CmdOut(name string, arg ...string) ([]byte, error) {
    cmd := exec.Command(name, arg...)
    out, err := cmd.CombinedOutput()
    out = bytes.TrimSpace(out)
    if err != nil {
        return nil, fmt.Errorf(err.Error() + ":" + string(out))
    }
    return out, nil
}

func CmdOutWithTimeout(command string, timeout int) ([]byte, error) {
    cmd := exec.Command("/bin/bash", "-c", command)
    done := make(chan error)
    var stdout, stderr bytes.Buffer
    cmd.Stdout, cmd.Stderr = &stdout, &stderr

    cmd.Start()
    go func() {
        done <- cmd.Wait()
    }()

    select {
    case <-time.After(time.Duration(timeout) * time.Second):
        err := cmd.Process.Kill() // timeout
        if err != nil {
            return stdout.Bytes(), fmt.Errorf(stderr.String() + err.Error())
        }
        return stdout.Bytes(), fmt.Errorf(stderr.String()+`TIMEOUT %d,Process "%s" has been killed`, timeout, strings.Join(cmd.Args, " "))
    case err := <-done:
        if err != nil {
            return nil, fmt.Errorf(err.Error() + ":" + stderr.String())
        }
        return stdout.Bytes(), err
    }
}
