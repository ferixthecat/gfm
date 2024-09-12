package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	// to read args by flag, must parse first
	flag.Parse()

	err := execCommand()
	if err != nil {
		log.Fatal(err)
	}
}

func execCommand() error {
	switch flag.Arg(0) {
	case "ls":
		files, err := os.ReadDir(".")
		if err != nil {
			return err
		}

		for _, file := range files {
			fmt.Printf("%s  ", file.Name())
		}
		fmt.Println()
	case "rm":
		path := getArg(1, "Enter path")
		return os.RemoveAll(path)
	case "pwd":
		currentPath, err := os.Getwd()
		if err != nil {
			return err
		}
		fmt.Println(currentPath)
	case "mv":
		path1 := getArg(1, "Enter path to move")
		path2 := getArg(2, "Enter path for final des")

		return os.Rename(path1, path2)
	case "mkdir":
		path := getArg(1, "Enter dir name")
		return os.MkdirAll(path, 0777)
	case "cat":
		path := getArg(1, "Enter path")
		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		os.Stdout.Write(data)
	case "touch":
		path := getArg(1, "Enter path")

		_, err := os.Create(path)
		return err
	case "cp":
		path1 := getArg(1, "Enter path to move")
		path2 := getArg(2, "Enter path for final des")
		data, err := os.ReadFile(path1)
		if err != nil {
			return err
		}
		return os.WriteFile(path2, data, 0777)
	case "echo":
		content := getArg(1, "Enter content")
		path := getArg(2, "Enter path")

		return os.WriteFile(path, []byte(content), 0777)
	}

	return nil
}

func getArg(indexArg int, mess string) string {
	val := flag.Arg(indexArg)
	if val == "" {
		log.Fatal(mess)
	}
	return val
}
