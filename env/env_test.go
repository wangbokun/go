package env

import (
	"testing"
	"os"
	"github.com/wangbokun/go/log"
)

func TestLoadDotEnvFile(t *testing.T){

	log.Debug("%s","testing")
	err := Load()
	if err != nil {
		log.Error("%s",err)
	}
	log.Info("%s",os.Getenv("test"))
}