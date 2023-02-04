package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"os"
)

func disableDNS(w http.ResponseWriter, r *http.Request) {
	mycwd, err := os.Getwd()
	if err != nil {
        log.Fatal(err)
    }
	execmd := fmt.Sprintf("%s/block.sh", mycwd)
	cmd := exec.Command(execmd)
	err = cmd.Run()
	if err != nil {
        fmt.Fprintf(w, "error calling block.sh")
    }
	fmt.Fprintf(w, "DNS Blocking Enabled.")
}

func enableDNS(w http.ResponseWriter, r *http.Request) {
	mycwd, err := os.Getwd()
	if err != nil {
        fmt.Fprintf(w, "Couldn't get current directory.")
    }
	execmd := fmt.Sprintf("%s/allow.sh", mycwd)
	cmd := exec.Command(execmd)
	err = cmd.Run()
	if err != nil {
        fmt.Fprintf(w, "error calling allow.sh")
    }
	fmt.Fprintf(w, "DNS Blocking Disabled.")
}

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)
	http.HandleFunc("/block", disableDNS)
	http.HandleFunc("/allow", enableDNS)

	log.Print("Listening on :3001...")
	err := http.ListenAndServe(":3001", nil)
	if err != nil {
		log.Fatal(err)
	}
}