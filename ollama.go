// helper function that deals with the ollama stuff
package main

import (
	"fmt"
	"os/exec"
)

var serveCmd *exec.Cmd

// checks if ollama exists on the system
func findOllama() bool {
    _, err := exec.LookPath("ollama")
    return err == nil
}

// starts ollama instance on a separate thread
func serveOllama() bool{
    serveCmd = exec.Command("ollama", "serve")
    if err := serveCmd.Start(); err != nil {
        return false
    }
    return true
}

// kills ollama serve process
func killOllama() {
    serveCmd.Process.Kill()
}

// runs an ollama model in the command line
// errors if the model doesn't exist
func runOllama(prompt, model string) (string, error) {
    promptCmd := exec.Command("echo", prompt)
    cmd := exec.Command("ollama", "run", model)

    pipe, err := promptCmd.StdoutPipe()
    if err != nil {
        fmt.Println("Error creating pipe: ", err)
        return "", err
    }
    cmd.Stdin = pipe

    if err := promptCmd.Start(); err != nil {
        fmt.Println("Error starting echo cmd", err)
        return "", err
    }
    out, err := cmd.CombinedOutput()
    if err != nil {
        fmt.Println("Error executing ollama cmd", err)
        return "", err
    }
    if err := promptCmd.Wait(); err != nil {
        fmt.Println("Error waiting for echo cmd", err)
        return "", err
    }
    fmt.Println(string(out))
    return string(out), nil
}
