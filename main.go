package main

import (
	"fmt"
	"os"
)

func main() {
    // check if ollama is on the system
    // if it isn't exit from the application
    // and send an error msg
    exists := findOllama()
    if !exists {
        fmt.Println("Ollama was not found. Please install.")
        os.Exit(1)
    }
    startInterface()
}
