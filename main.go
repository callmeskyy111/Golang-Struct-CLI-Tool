//Struct Project - Getting user input

package main

import (
	"fmt"

	"example.com/go-project/note"
)

func getNoteData()(string ,string){
	title:=getUserInput("Not Title:")

	content:=getUserInput("Note Content:")
	return  title, content
}

func main() {
	title, content := getNoteData()
	userNote,err:= note.New(title,content)
	if err != nil{
		fmt.Println("🔴",err)
		return 
	}
	userNote.Display()
	
}

func getUserInput(prompt string)string{
fmt.Print(prompt)
var val string
fmt.Scanln(&val) //todo: 93 - Solution to Scanln
return val
}