package main

import (
	"flag"
	"fmt"
	"net/http"
	"os/exec"
	"strings"
	"time"
)

func main() {
	port := 20508
	flag.IntVar(&port, "port", port, fmt.Sprintf("Set Port. Default is %d", port))
	flag.Parse()
	fmt.Printf("Start ms-sh service http://127.0.0.1:%d \n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s := r.FormValue("s")
		w.Header().Set("Content-Type", "text/plain;charset=utf-8")
		w.Write([]byte(Command(s)))
	}))
	if err != nil {
		fmt.Print("Error:" + err.Error())
	}
}

func Command(s string) (rs string) {
	defer func() {
		if e := recover(); e != nil {
			rs = "Exec Fail Error: " + e.(error).Error()
		}
		fmt.Println(time.Now().Format(time.RFC3339) + " " + rs)
	}()
	if s == "" {
		return "Exec Fail Error: Params `s` is nil"
	}
	l := strings.Split(s, " ")
	var cmd *exec.Cmd
	if len(l) == 1 {
		cmd = exec.Command(l[0])
	} else {
		cmd = exec.Command(l[0], l[1:]...)
	}
	bytes, err := cmd.CombinedOutput()
	if err != nil {
		return "Exec Fail Error:" + err.Error()
	}
	return "Exec Success... \nCommand: " + s + " \nResult: \n" + string(bytes)
}
