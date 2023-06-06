package exec

import (
	"fmt"
	"os/exec"
)

func Run_multiple_args() {
	prg := "echo"

	arg1 := "there"
	arg2 := "are slurms"
	arg3 := "in SlurmLand"

	cmd := exec.Command(prg, arg1, arg2, arg3)

	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Print(string(stdout))
}
