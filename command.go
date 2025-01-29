package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type cmdFlags struct {
	Add    string
	Del    int
	Edit   string
	Toggle int
	List   bool
}

func NewCmdFlags() *cmdFlags {
	cf := cmdFlags{}
	flag.StringVar(&cf.Add, "add", "", "Add a new todo specify title")
	flag.StringVar(&cf.Edit, "edit", "", "Edit a todo by index & specify a newtitle. id:new_title")
	flag.IntVar(&cf.Del, "delete", -1, "Specify Index to be deleted")
	flag.IntVar(&cf.Toggle, "toggle", -1, "Specify Index to be toggled")
	flag.BoolVar(&cf.List, "list", false, "List all todos")
	flag.Parse()

	return &cf
}

func (cf *cmdFlags) Execute(todos *Todos) {
	switch {
	case cf.List:
		todos.Print()
	case cf.Add != "":
		todos.Add(cf.Add)
	case cf.Edit != "":
		parts := strings.SplitN(cf.Edit, ":", 2)
		if len(parts) != 2 {
			fmt.Println("Error , invalid format for edit. Please use id:new_title")
			os.Exit(1)
		}

		index,err :=strconv.Atoi(parts[0])
	
	   if err!= nil {
		 fmt.Print("Error:invalid index for edit")
		 os.Exit(1)
	   }

	   todos.Edit(index,parts[1])
    case cf.Toggle!=-1:
		todos.Toggle(cf.Toggle)
	case cf.Del!=-1:
		todos.Delete(cf.Del)
	default:
		fmt.Println("Invalid Command")
	}

}
