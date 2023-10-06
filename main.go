package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
 
)

var cliName string = "Pokedex"

func printPrompt() {
	fmt.Println(cliName, "> ")
}


func printUnknown(text string){
	fmt.Println(text, ": command not found")
}


func commandHelp() error {
	fmt.Printf("Welcome to the %s! These are avaliable commands: \n", cliName)
	fmt.Println("help	- Show avaliable commands")
	fmt.Println("exit	- Exit the program")
    return nil
	
}
func commandExit() error {
    fmt.Printf("Exiting %s...\n", cliName)
    os.Exit(0)
    return nil
}

func cleanInput(text string) string {
    output := strings.TrimSpace(text)
    output = strings.ToLower(output)
    return output
}

type cliCommand struct{
	name string
	description string
	callBack func() error
}


func main(){
	commands := map[string]cliCommand{
        "help": {
            name:        "help",
            description: "Displays a help message",
            callBack:    commandHelp,
        },
        "exit": {
            name:        "exit",
            description: "Exit the Pokedex",
            callBack:    commandExit,
        },
    }
    reader := bufio.NewScanner(os.Stdin)
    printPrompt()

    for reader.Scan(){
        text:= cleanInput(reader.Text())
            if command, ok := commands[text]; ok {
                command.callBack()
        } else {
            printUnknown(text)
        }
        printPrompt()
    }

}
