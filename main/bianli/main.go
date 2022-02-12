package main

import "fmt"

func main() {
	var str string = "hello,world!中国"
	str2 := []rune(str)
	for i := 0; i < len(str2); i++ {
		fmt.Printf("%c\n", str2[i])
	}
	fmt.Println("---------------------")
	for _, value := range str {
		fmt.Printf("value:%c\n", value)
	}
	
}