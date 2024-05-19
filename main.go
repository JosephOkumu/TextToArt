package main

import (
	"texttoart/art"
	"fmt"
	"os"
)

// main function reads a banner file, creates a map of ASCII art, validates user input,
// and prints the corresponding ASCII art to the console.
func main() {
	// Read banner file
	bannerFile, err := art.ReadBannerFile("standard.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	// Create map of ASCII art
	ASCIIArtMap := art.MapCreator(string(bannerFile))

	// Validate user input
	inputText, err := art.ValidateInput(os.Args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	// Check for empty input or newline character
	if inputText == "" {
		return
	}
	if inputText == "\\n" {
		fmt.Println()
		return
	}

	// Get ASCII art corresponding to input text
	asciiArt, err := art.ArtRetriever(inputText, ASCIIArtMap)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(asciiArt)
}
