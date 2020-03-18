package tlv_solver

import (
	"fmt"
	"os"
)

func InitializeApp(args []string) {
	if len(args) != 0 {
		input, err := GetUserInput(args)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		ProcessInput(input)
	} else {
		input, err := GetUserInputFromFile(os.Stdin)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		ProcessInput(input)
	}
}


