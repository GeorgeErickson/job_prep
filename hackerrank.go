package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	s := bufio.NewScanner(os.Stdin)

	l := ""
	for i := 0; i < 2; i++ {

		l = s.Text()
	}
	fmt.Println(l)
	sum := 0
	for _, v := range strings.Fields(l) {
		i, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		sum += i
	}

	fmt.Println(sum)

}
