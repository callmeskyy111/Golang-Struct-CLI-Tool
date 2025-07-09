👨🏻‍💻 Let’s walk through our Go project step-by-step. We're building a **CLI note-taking tool**
---

## 📁 Project Structure (2 files)

### `main.go`

Handles:

* User input,
* Note creation,
* Displaying and saving the note.

### `note/note.go`

Defines:

* The `Note` struct,
* Methods for displaying and saving,
* A constructor function `New()` for creating validated notes.

---

## 🧠 FILE 1: `main.go`

```go
package main
```

Declares this as the **main package**, the entry point of the program.

### ✅ Imports

```go
import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"example.com/go-project/note"
)
```

* `bufio`, `os`: for reading input from terminal.
* `strings`: to trim newline characters.
* `fmt`: for printing output.
* `note`: custom module for handling notes.

---

### 🔹 `getUserInput(prompt string) string`

This function prints a prompt and reads the user's typed input.

```go
reader := bufio.NewReader(os.Stdin)
text, err := reader.ReadString('\n')
```

* Uses `bufio.NewReader` for better line reading than `fmt.Scanln`.
* Trims `\n` or `\r` (cross-platform newline handling).

---

### 🔹 `getNoteData() (string, string)`

Prompts user for both `title` and `content`:

```go
title := getUserInput("Note Title:")
content := getUserInput("Note Content:")
return title, content
```

---

### 🔹 `main()`

This is the driver logic:

```go
title, content := getNoteData()
```

#### 🔸 Create Note

```go
userNote, err := note.New(title, content)
```

* Calls constructor from the `note` package.
* Validates empty fields.

#### 🔸 Display Note

```go
userNote.Display()
```

* Shows note title and content.

#### 🔸 Save Note

```go
err = userNote.Save()
```

* Writes the note to a `.json` file.

---

## 📁 FILE 2: `note/note.go`

```go
package note
```

This is a **custom package** named `note`.

---

### 🧱 Struct Definition

```go
type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}
```

A **custom type** `Note` with:

* Title (user-entered)
* Content (user-entered)
* CreatedAt (timestamp when note is created)

**JSON tags** make sure this struct is saved with the correct key names when encoded to JSON.

---

### 🔸 `func (n Note) Display()`

```go
fmt.Printf("📝 Your note Titled %v has the following Content:\n\n%v\n", n.Title, n.Content)
```

* Method on `Note`
* Displays nicely formatted note in terminal

---

### 🔸 `func (n Note) Save() error`

```go
fileName := strings.ReplaceAll(n.Title, " ", "_")
fileName = strings.ToLower(fileName) + ".json"
```

* Converts title like `My Note` to `my_note.json`

```go
json, err := json.Marshal(n)
```

* Marshals (converts) `Note` struct to JSON

```go
return os.WriteFile(fileName, json, 0644)
```

* Saves it to a file with **read-write** permission

---

### 🔸 `func New(title, content string) (Note, error)`

```go
if title == "" || content == "" {
	return Note{}, errors.New("invalid input")
}
```

* Validates that both fields are not empty.

```go
return Note{
	Title:     title,
	Content:   content,
	CreatedAt: time.Now(),
}, nil
```

* Returns a new note with current timestamp.

---

## 📂 Output Example

If the user enters:

```
Note Title: Grocery List
Note Content: Buy eggs, milk, and bread.
```

A file `grocery_list.json` is created:

```json
{
  "title": "Grocery List",
  "content": "Buy eggs, milk, and bread.",
  "created_at": "2025-07-03T15:04:05Z"
}
```

---

## 🔄 Flow Summary

```bash
User Input ➜ Struct Creation ➜ JSON File Write
```

1. We get title & content via terminal
2. Validate and construct `Note` struct
3. Marshal it to JSON
4. Save to file

---
