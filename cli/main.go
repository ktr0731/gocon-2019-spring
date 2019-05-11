package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/ktr0731/go-fuzzyfinder"
)

func main() {
	var slice []string
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		slice = append(slice, s.Text())
	}
	idx, err := fuzzyfinder.Find(slice, func(i int) string {
		return slice[i]
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to find: %s\n", err)
		os.Exit(1)
	}
	fmt.Println(slice[idx])
}
