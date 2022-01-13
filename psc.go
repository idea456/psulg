package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getInput(reader *bufio.Reader) string {
	t, _ := reader.ReadString('\n')
	return strings.TrimSpace(t)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("PSU Language | psuc 1.0.0\n")
	fmt.Print("Type exit to exit the program or press Ctrl-D.\n")

	fmt.Print(">>> ")
	line := getInput(reader)
	for ; !strings.EqualFold("exit", line); line = getInput(reader) {
		s := NewScanner(line)
		arr := s.Scan()
		fmt.Println(arr)
		parser := NewParser(arr)
		expr := parser.Parse()
		fmt.Printf("%#v\n", expr)
		fmt.Print(">>> ")
	}
	fmt.Print("Bai bai!\n")
}
