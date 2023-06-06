package exec

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func Run_app_with_simple_args() {
	cmd := exec.Command("tr", "a-z", "A-Z")

	cmd.Stdin = strings.NewReader("Little slurm goes big")

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("translated phrase: %q\n", out.String())
}
