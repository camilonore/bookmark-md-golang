package main

import (
	"encoding/json"
	"log"
	"os"
	"types"
	"utils"
)

func main() {
	bookmarks_url := "C:\\Users\\USUARIO\\AppData\\Local\\Google\\Chrome\\User Data\\Default\\Bookmarks"
	bookmarks_file, err := os.Open(bookmarks_url)
	if err != nil {
		log.Fatal(err)
	}
	var bookmarks types.Data
	json.NewDecoder(bookmarks_file).Decode(&bookmarks)

	bookmarks_file.Close()

	roots := bookmarks.Roots

	for _, folder := range roots.BookmarkBar.Children {
		if folder.Type == "folder" {
			for _, folder_child := range folder.Children {
				utils.FolderProcess(folder.Name, folder_child.Name, folder_child.Url)
				if folder_child.Type == "folder" {
					for _, nestedFolder := range folder_child.Children {
						utils.FolderProcess(folder.Name, nestedFolder.Name, nestedFolder.Url)
					}
				}
			}
		}
	}
}
