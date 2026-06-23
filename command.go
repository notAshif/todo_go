package main

import (
	"flag"
	"log"
	"os"
	"strconv"
	"strings"
)

type CommandFlag struct {
	add string
	rm  int
	ed  string
	tg  int
	ls  bool
}

func newCommand() *CommandFlag {
	cf := CommandFlag{}

	flag.StringVar(&cf.add, "Add", "", "Add a new todo specifying title.")
	flag.StringVar(&cf.ed, "ed", "", "Updated or edit todo by index specifying title. eg:- id:new_title")
	flag.IntVar(&cf.tg, "tg", -1, "Specify new todon to toggle.")
	flag.IntVar(&cf.rm, "rm", -1, "Remove existing todo by index.")
	flag.BoolVar(&cf.ls, "ls", false, "Lists all todo's.")

	flag.Parse()

	return &cf
	
}


// Execute the command
func (cf *CommandFlag) Execute(t *Todos){
	switch {
	case cf.ls:
		t.print()
	case cf.add != "":
		t.add(cf.add)
	case cf.rm != -1:
		t.delete(cf.rm)
	case cf.tg != -1:
		t.toggle(cf.tg)

	case cf.ed != "":
		part := strings.SplitN(cf.ed, ":", 2)

		if len(part) != 2{
			log.Printf("INDALID INDEX: PLEASE USE id:new_title");
			os.Exit(1)
		}
		
		i, err := strconv.Atoi(part[0])
		if err != nil {
			log.Printf("INVALID INDEX FOR EDIT ❌ %s", err)
			os.Exit(1)
		}

		t.update(i, part[1])

	default:
		log.Printf("INVALID COMMAND ❌");
	}
}