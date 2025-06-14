package main

import (
	"escuuta/srcYoutube"
	"fmt"
)

func main() {

	name, err := srcYoutube.GetTitle("https://www.youtube.com/watch?v=d5dj2XrkvQk")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Title:", name)

}
