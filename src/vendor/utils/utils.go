package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func openbrowser(url string) {
	err := exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	if err != nil {
		log.Fatal(err)
	}
}

func closeBrowser() {
	exec.Command("taskkill", "/im", "chrome.exe", "/f").Run()
}

func doQuestions(bookmark_name *string) (string, string, string) {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("What category this bookmark belongs to?")
	scanner.Scan()
	category := scanner.Text()

	fmt.Printf("Do you like the bookmarks name? %s (y/n)\n", *bookmark_name)
	scanner.Scan()
	answer := scanner.Text()

	if answer == "n" {
		fmt.Println("What name do you want to give to this bookmark?")
		scanner.Scan()
		*bookmark_name = scanner.Text()
	}

	fmt.Println("Write a description for this bookmark")
	scanner.Scan()
	description := scanner.Text()

	return category, *bookmark_name, description
}

func createFileHeader(file os.File) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Write the title of this category")
	scanner.Scan()
	title := scanner.Text()

	fmt.Println("Write a description for this category")
	scanner.Scan()
	description := scanner.Text()

	headerText := fmt.Sprintf("---\ntitle: %s\ndescription: %s\n---", title, description)

	file.WriteString(fmt.Sprintf("%s\n", headerText))
}

func createFile(file_name string) {
	if !fileExists(file_name) {
		file, err := os.Create(fmt.Sprintf("./md/%s.md", file_name))
		createFileHeader(*file)
		if err != nil {
			log.Fatal(err)
		}
		file.Close()
	}
}

func fileExists(file_name string) bool {
	if _, err := os.Stat(fmt.Sprintf("./md/%s.md", file_name)); os.IsNotExist(err) {
		return false
	}
	return true
}

func writeToFile(file_name string, bookmark_name string, bookmark_url string, description string) {
	file, err := os.OpenFile(fmt.Sprintf("./md/%s.md", file_name), os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf("- ## [%s](%s)\n\t%s\n", bookmark_name, bookmark_url, description))

	if err != nil {
		panic(err)
	}

}

func FolderProcess(folder_name string, name string, url string) {
	fmt.Printf("------ Folder: %s ------\n", folder_name)
	openbrowser(url)
	category, bookmark_name, description := doQuestions(&name)
	createFile(category)
	writeToFile(category, bookmark_name, url, description)
	closeBrowser()
}
