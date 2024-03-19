package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: client <command> [args...]")
		return
	}

	command := os.Args[1]
	args := os.Args[2:]

	switch command {
	case "store":
		storeFiles(args)
	case "update":
		updateFile(args)
	case "delete":
		deleteFile(args)
	default:
		fmt.Printf("Unknown command: %s\n", command)
	}
}

func storeFiles(files []string) {
	for _, file := range files {
		data, err := os.ReadFile(file)
		if err != nil {
			log.Printf("Failed to read file %s: %v", file, err)
			continue
		}

		resp, err := http.Post("http://localhost:8080/store", "text/plain", bytes.NewBuffer(data))
		if err != nil {
			log.Printf("Failed to store file %s: %v", file, err)
			continue
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			log.Printf("Failed to store file %s: %s", file, resp.Status)
		} else {
			fmt.Printf("File %s stored successfully\n", file)
		}
	}
}

func updateFile(args []string) {
	if len(args) != 2 {
		fmt.Println("Usage: client update <filename> <file>")
		return
	}

	filename := args[0]
	file := args[1]

	data, err := os.ReadFile(file)
	if err != nil {
		log.Printf("Failed to read file %s: %v", file, err)
		return
	}

	req, err := http.NewRequest(http.MethodPut, "http://localhost:8080/update?filename="+filename, bytes.NewBuffer(data))
	if err != nil {
		log.Printf("Failed to create request: %v", err)
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("Failed to update file %s: %v", filename, err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Failed to update file %s: %s", filename, resp.Status)
	} else {
		fmt.Printf("File %s updated successfully\n", filename)
	}
}

func deleteFile(args []string) {
	if len(args) != 1 {
		fmt.Println("Usage: client delete <filename>")
		return
	}

	filename := args[0]

	req, err := http.NewRequest(http.MethodDelete, "http://localhost:8080/delete?filename="+filename, nil)
	if err != nil {
		log.Printf("Failed to create request: %v", err)
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("Failed to delete file %s: %v", filename, err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Failed to delete file %s: %s", filename, resp.Status)
	} else {
		fmt.Printf("File %s deleted successfully\n", filename)
	}
}