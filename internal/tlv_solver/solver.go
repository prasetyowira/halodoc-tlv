package tlv_solver

import (
	"fmt"
	"strings"
)

func ProcessInput(inputData []*UserInput) {
	for _, input := range inputData {
		input.solve()
	}
}

func (i *UserInput) solve() {
	switch i.Types {
		case UPPRCS:
			value := strings.ToUpper(i.Value)
			result := &UserOutput{
				Types: i.Types,
				Value: value,
			}
			result.print()
			break
		case REPLCE:
			result := &UserOutput{
				Types: i.Types,
				Value: "THIS STRING",
			}
			result.print()
			break
		default:
			fmt.Println("Type not valid")
	}
}

func (o *UserOutput) print() {
	fmt.Printf("%s-%s\n", o.Types, o.Value)
}