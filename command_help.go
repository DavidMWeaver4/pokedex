package main
import "fmt"
func helpCommand(cfg *config, args []string) error{
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for name, description:= range getCommands(cfg){
		fmt.Printf("%s: %s\n", name, description.description)
	}
	return nil
}
