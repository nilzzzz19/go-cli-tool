package main

import (
	"flag"
	"fmt"
	"os"

	todo "github.com/nilzzzz19/go-cli-tool"
)

const (
	todoFile = ".todos.json"
)

func main() {
	//flags
	add := flag.Bool("add", false, "Add a todo item") //returns the address of the bool var

	//parsing in cli
	flag.Parse()

	todos := &todo.Todos{} //gets the ds

	if err := todos.Load(todoFile); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	switch {
	case *add:
		todos.Add("sample task")
		err := todos.Store(todoFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
	default:
		fmt.Fprintln(os.Stderr)
		os.Exit(1)
	}
}
