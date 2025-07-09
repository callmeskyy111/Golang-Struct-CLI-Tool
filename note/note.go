package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title     string `json:"title"`
	Content   string `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func (n Note) Display(){
	fmt.Printf("📝 Your note Titled %v has the following Content:\n\n%v\n",n.Title, n.Content,)
}

func (n Note) Save()error{
//Encoding JSON content {...}
fileName := strings.ReplaceAll(n.Title, " ","_")
fileName = strings.ToLower(fileName) + ".json"
json,err:= json.Marshal(n)
if err != nil{
	return err
}
return os.WriteFile(fileName,json, 0644)
}

func New(Title, Content string)(Note, error){
// Validation
if Title =="" || Content == ""{
	return Note{},errors.New("invalid input")
}
return Note{
	Title:Title,
	Content: Content,
	CreatedAt: time.Now(),
},
nil // nil, if no error
}