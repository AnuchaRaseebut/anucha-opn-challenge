package main

import (
	"fmt"
	"os"

	"anucha-challenge/donation"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a path to the file")
		return
	}

	filepath := os.Args[1]
	cfg := donation.GetEnv()
	if err := donation.ProcessDonation(cfg, filepath); err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("donation process successful")
}
