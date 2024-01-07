package main

import (
	"flag"
)

func stats(email string) {
	print(email)
}

func main() {
	var folder string
	var email string
	print(folder)
	flag.StringVar(&folder, "add", "", "Add a new folder to scan for Git Repositories")
	flag.StringVar(&email, "email", "your@email.com", "the email to scan")
	flag.Parse()

	if folder != "" {
		print(folder)
		scan(folder)
		return
	}

	stats(email)
}
