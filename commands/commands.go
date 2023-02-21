package commands

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)

func GetInput() (string, error) {

	fmt.Print("=> ")
	str, err := reader.ReadString('\n')

	if err != nil {
		return "", err
	}

	s := strings.Replace(str, "\r\n", "", 1)

	return s, nil
}
