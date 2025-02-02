// helper function that deals with the ollama stuff
package main

import (
    "os/exec"
)

// checks if ollama exists on the system
func findOllama() bool {
    _, err := exec.LookPath("ollama")
    return err == nil
}

// starts ollama instance on a separate thread
func serveOllama() bool{
    cmd := exec.Command("ollama", "serve")
    if err := cmd.Start(); err != nil {
        return false
    }
    return true
}

// runs an ollama model in the command line
// errors if the model doesn't exist
func runOllama(model string) (bool, error) {
    cmd := exec.Command("ollama", "run", model)
    if err := cmd.Start(); err != nil {
        return false, err
    }
    return true, nil
}
