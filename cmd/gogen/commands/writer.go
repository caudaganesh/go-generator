package commands

import (
	"io"
	"log"
	"os"

	"github.com/caudaganesh/go-generator/constant"
	"github.com/caudaganesh/go-generator/runner"
)

func Write(oPath string, out io.Reader, action string) {
	if oPath != "" {
		oPath = oPath + constant.MapActionsToPrefix[action]
		w, err := os.OpenFile(oPath, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}
		_, err = io.Copy(w, out)
		w.Close()
		if err != nil {
			log.Fatal(err)
		}
		err = runner.GoImports(oPath)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		_, err := io.Copy(os.Stdout, out)
		if err != nil {
			log.Fatal(err)
		}
	}

}
