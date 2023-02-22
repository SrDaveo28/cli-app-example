package commands

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)

func GetInput() (string, error) {

	str, err := reader.ReadString('\n')

	if err != nil {
		return "", err
	}
	return str, nil
}
func ShowInConsole(expenses []float32) {
	builder := strings.Builder{}

	for _, expense := range expenses {
		builder.WriteString(fmt.Sprintf("Expense: %6.2f\n", expense))
	}

	fmt.Println(builder.String())
}

func Export() {}
