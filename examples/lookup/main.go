// Copyright 2016 David Lavieri.  All rights reserved.
// Use of this source code is governed by a MIT License
// License that can be found in the LICENSE file.

package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/falmar/goradix"
)

func main() {

	radix := goradix.New()

	exampleData(radix)

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Text to Search: ")
	for scanner.Scan() {

		textToLook := scanner.Text()

		node, err := radix.LookUp(textToLook)

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Found node Value:", node.Get())
		}

		fmt.Print("Text to Search: ")
	}
}

func exampleData(radix *goradix.Radix) {
	text := []string{"test", "toaster", "toasting", "slow", "slowly"}
	for i, v := range text {
		fmt.Println("Inserted:", v, "Value: ", i)
		radix.Insert(v, i)
	}
}