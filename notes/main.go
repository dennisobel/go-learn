package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/notes/note"
	"example.com/notes/todo"
	// "golang.org/x/text/cases"
)

func add[T int|float64|string](a,b T)T{
	return a + b
}

type saver interface {
	Save() error
}

// type displayer interface {
// 	Display()
// }

// type outputable interface {
// 	Save() error
// 	Display()
// }

type outputtable interface {
	saver
	Display()
}

func main() {
	result := add(1,2)
	fmt.Println(result)
	printSomething(1)
	printSomething(1.5)
	printSomething("Hello")
	title, content := getNoteData()
	todoText := getUserInput("TODO Text: ")

	todo, err := todo.New(todoText)
	printSomething(todo)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(todo)
	if err != nil {
		fmt.Println(err)
		return
	}

	outputData(userNote)
}

func printSomething(value interface{}) {
	intVal, ok := value.(int)

	if ok {
		fmt.Println("Integer: ", intVal)
		return
	}

	floatVal, ok := value.(float64)

	if ok {
		fmt.Println("Float: ", floatVal)
		return
	}

	stringVal, ok := value.(string)

	if ok {
		fmt.Println("String: ", stringVal)
		return
	}

	// switch value.(type) {
	// case int:
	// 	fmt.Println("Integer:", value)
	// case float64:
	// 	fmt.Println("Float:", value)
	// case string:
	// 	fmt.Println(value)
	// }
}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
}

func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("Saving the note failed.")
		return err
	}
	fmt.Println("Saving the note succeeded!")
	return nil
}

func getNoteData() (string, string) {
	title := getUserInput("Note title: ")

	content := getUserInput("Note content: ")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
