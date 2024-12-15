package main

import (
	"log"
	"os"
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/rajatnai/task/cmd"
	"github.com/rajatnai/task/db"
)

func main() {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "task.db")
	err := db.Init(dbPath)
	if err != nil {
		log.Fatal("Error in the connection with database")
		os.Exit(1)
	}
	cmd.RootCmd.Execute()
}
