package listener

import "fmt"

// Listen is the entry point that wires together config, subscription, and output.
func Listen(cfg Config) error {
	fmt.Println("Listening with config:")
	fmt.Printf("%+v\n", cfg)
	return nil
}
