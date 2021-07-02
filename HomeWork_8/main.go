package main

import (
	"fmt"
	"Second/foo"
)
	//type Config struct{
	//Port string
	//Name string
	//Count int
	//}
	func main(){
		c := foo.NewConfig()
		fmt.Println(c)
		//var port = flag.String("port","8080","Port number")
		//var name = flag.String("name", "Go student", "Name for hello")
		//var count = flag.Int("Count",1,"Number of repeats")
		//flag.Parse()

		//config := Config{
		//Port: *port,
		//Name: *name,
		//Count: *count,
	}
	//fmt.Println(config)
	//runServer(config)
	//c := Foo.NewConfig()
	//fmt.Println(c)
}


