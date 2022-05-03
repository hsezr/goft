package goft

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func LoadConfigFile() []byte {
	dir, _ := os.Getwd()
	fmt.Println(dir)
	file := dir + "/application.yaml"
	b, err := ioutil.ReadFile(file)
	if err != nil {
		log.Println(err)
		return nil
	}

	return b
}