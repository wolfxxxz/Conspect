package exec

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
)

func Output_pipe() {
	cmd := exec.Command("echo", "piping slurms")

	stdout, err := cmd.StdoutPipe()

	if err != nil {
		log.Fatal(err)
	}

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	data, err := ioutil.ReadAll(stdout)

	if err != nil {
		log.Fatal(err)
	}

	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", string(data))
}
