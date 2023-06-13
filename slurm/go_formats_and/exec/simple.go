package exec

import (
	"log"
	"os/exec"
)

func RunSimpleApp() {
	cmd := exec.Command("firefox")

	err := cmd.Run()

	if err != nil {
		log.Fatal(err)
	}
}
