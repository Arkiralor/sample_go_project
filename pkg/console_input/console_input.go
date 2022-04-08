package console_input

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func NumberFromCon() int {
	reader := bufio.NewReader(os.Stdin)
	input_str, _ := reader.ReadString('\n')
	input_str = RemoveEOLMarkers(input_str)
	input_int, err := strconv.Atoi(input_str)

	if err != nil {
		log.Fatalf("Error Code: %v", err)
	}

	return input_int
}

func RemoveEOLMarkers(s string) string {
	s = strings.Replace(s, "\r", "", -1)
	s = strings.Replace(s, "\n", "", -1)

	return s
}
