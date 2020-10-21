package testhelper

import (
	"io/ioutil"
	"log"
)

func GetExpectFromFile(path string) string {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	return string(content)
}
