package exec

import (
	"fmt"
	"log"
	"os/exec"
)

func Capture_output() {
	out, err := exec.Command("ls", "-l").Output()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(out))
}
