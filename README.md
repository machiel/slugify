# Overview

Slugify is a small library that turns strings in to slugs.

# License
Slugify is licensed under a MIT license.

# Installation
A simple `go get github.com/Machiel/slugify/` should suffice.

# Usage

## Example

	package main

	import (
		"fmt"
		"github.com/Machiel/slugify"
	)

	func main()
	{
		fmt.Println(slugify.Slugify("Hello, world!")) // Will print: hello-world
		fmt.Println(slugify.Slugify("💻  I love this computer! 💻 ")) // Will print: i-love-this-computer
	}
