package main

import (
	"fmt"

	"github.com/Machiel/slugify"
)

func main() {
	fmt.Println(slugify.Slugify("Hello, world!"))               // Will print: hello-world
	fmt.Println(slugify.Slugify("💻  I love this computer! 💻 ")) // Will print: i-love-this-computer
}
