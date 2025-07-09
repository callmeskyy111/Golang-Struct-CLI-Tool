//Struct Project - Getting user input and saving to a json file

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/go-project/note"
)

func getNoteData()(string ,string){
	title:=getUserInput("Note Title:")

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
	err = userNote.Save()
	if err != nil{
		fmt.Println("Saving the note failed 🔴")
		return
	}
	fmt.Println("Successfully saved the note ✅")
}

func getUserInput(prompt string)string{
fmt.Printf("%v ",prompt)
 //var val string
// fmt.Scanln(&val) //Scanln has it's limitations
reader:=bufio.NewReader(os.Stdin)
text,err:= reader.ReadString('\n')
if err != nil{
	return ""
}
text = strings.TrimSuffix(text,"\n")
text = strings.TrimSuffix(text,"\r")
return text
}