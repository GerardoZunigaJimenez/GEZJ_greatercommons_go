package main

import (
	"fmt"
)

func main()  {
	s := `"Hello, 
	playground`
	s = "Hello Hawai"
	fmt.Println(s)
	fmt.Printf("%T\n",s)

	bs:= []byte(string (s))
	fmt.Println(bs)
	fmt.Printf("%T\n",bs)

	for i := 0; i < len(s); i++{
		fmt.Printf("%#U", s[i])
	}

	for i, v := range s{
		fmt.Println(i,v)
	}
}
