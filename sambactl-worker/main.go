package main

import (
	"fmt"
	"flag"
	"log"
	"os"
	"os/exec"
	"encoding/json"
	"bytes"
	"strings"
)

type User struct {
	Username	string
}

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		log.Fatalln("no command")
	}
	if args[0] == "list" {
		cmd := exec.Command("pdbedit", "-L")
		var outBuffer bytes.Buffer
		cmd.Stdout = &outBuffer
		err := cmd.Run()
		if err != nil {
			fmt.Fprintln(os.Stderr, "pdbedit failed")
			os.Exit(1)
		}
		out := outBuffer.String()
		userLines := strings.Split(out, "\n")
		users := make([]User, 0, len(userLines))
		for _, s := range userLines {
			if len(s) == 0 {
				continue
			}
			parts := strings.Split(s, ":")
			u := User{Username: parts[0]}
			users = append(users, u)
		}
		data, err := json.Marshal(users)
		if err != nil {
			fmt.Fprintln(os.Stderr, "json marshal failed: %s", err.Error())
		}
		os.Stdout.Write(data)
	} else if args[0] == "passwd" {
		// XXX: 2nd arg is dummy
		cmd := exec.Command("mysmbpasswd", "-D0", args[1], args[2])
		var out bytes.Buffer
		cmd.Stdout = &out
		cmd.Run()
	} else if args[0] == "adduser" {
		name := args[1]
		pass := args[2]
		cmd := exec.Command("smb_adduser", name, pass)
		var out bytes.Buffer
		cmd.Stdout = &out
		cmd.Run()
	} else if args[0] == "deluser" {
		name := args[1]
		cmd := exec.Command("smb_deluser", name)
		var out bytes.Buffer
		cmd.Stdout = &out
		cmd.Run()
	}
}
