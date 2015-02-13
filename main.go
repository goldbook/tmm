package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3" // 直接呼ばれないライブラリの書き方
	"github.com/mitchellh/cli"
)

type Node struct {
	Id        int64
	ParentId  int64
	Body      string `sql:"type:varchar(100);"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Start struct{}

func (f *Start) Run(args []string) int {
	ui := cli.BasicUi{os.Stdin, os.Stdout, os.Stdout}

	var input, _ = ui.Ask("dbファイル名: ")
	var fn = input + ".tmm"

	// DBファイルの未作成ならmigrate
	if _, err := os.Stat(fn); os.IsNotExist(err) {
		fmt.Printf("DBファイルを作成: %s \n", fn)
		db, _ := gorm.Open("sqlite3", fn)
		db.CreateTable(&Node{})

		var body, _ = ui.Ask("ルートノードの本文:")
		db.Create(&Node{Id: 1, Body: body, CreatedAt: time.Now(), UpdatedAt: time.Now()})
	}
	db, _ := gorm.Open("sqlite3", fn)

	for {
		fmt.Println("コマンド: tree, cd, ls, add")
		var command, _ = ui.Ask("tmm>")
		switch command {
		case "tree":
			fmt.Println(command)
		case "cd":
			fmt.Println(command)
		case "ls":
			fmt.Println(command)
		case "add":
			fmt.Println(command)
			// var node Node
			// db.First(&node)
			// db.Create(&Node{ParentId: node.Id, Body: body, CreatedAt: time.Now(), UpdatedAt: time.Now()})
		default:
			fmt.Println(command)
		}
	}

	return 0
}

func (f *Start) Help() string {
	return `start tmm to edit text mind map.
and edit or create mindmap at current dir.`
}

func (f *Start) Synopsis() string {
	return "start tmm."
}

func main() {
	c := cli.NewCLI("tmm - Text MindMap", "1.0.0")
	c.Args = os.Args[1:]
	c.Commands = map[string]cli.CommandFactory{
		"start": func() (cli.Command, error) {
			return &Start{}, nil
		},
	}

	exitStatus, err := c.Run()
	if err != nil {
		log.Println(err)
	}

	os.Exit(exitStatus)
}
