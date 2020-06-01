package main

import (
	"fmt"
	"os"
	"os/user"
	"compiler/reader"

)

func main(){
	
	user, err := user.Current()

	if err != nil{
		panic(err)
	}

	fmt.Printf("Please type the command \n")

	repl.Start(os.Stdin,os.Stdout)


}
