package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"os/exec"
)

func main() {
	http.HandleFunc("/", LampStatus)
	http.HandleFunc("/off", LampOff)
	http.HandleFunc("/on", LampOn)
	http.HandleFunc("/pic", LampPic)
	http.HandleFunc("/set-by-schedule", LampSetBySchedule)
	http.ListenAndServe(":6969", nil)
}

func LampStatus(w http.ResponseWriter, r *http.Request) {
	out, err := exec.Command("/home/dietpi/Code/lamp/lamp.py").Output()
	if err != nil {
		renderPage(w, err.Error())
		return
	}
	renderPage(w, fmt.Sprintf("lamp is: %s\n", out))
}

func LampOnOff(w http.ResponseWriter, r *http.Request, state string) {
	err := exec.Command("/home/dietpi/Code/lamp/lamp.py", state).Run()
	if err != nil {
		renderPage(w, err.Error())
		return
	}
	renderPage(w, fmt.Sprintf("lamp has been turned: %s\n", state))
}

func LampOn(w http.ResponseWriter, r *http.Request) {
	LampOnOff(w, r, "on")
}

func LampOff(w http.ResponseWriter, r *http.Request) {
	LampOnOff(w, r, "off")
}

func LampPic(w http.ResponseWriter, r *http.Request) {
	out, err := exec.Command("fswebcam", "-S", "10", "-").Output()
	if err != nil || len(out) == 0 {
		errMsg := "can't get image"
		if err != nil {
			errMsg = errMsg + "\n" + err.Error()
		}
		renderPage(w, errMsg)
		return
	}
	renderPage(w, fmt.Sprintf(`
		<div>
		<img src="data:image/png;base64, %s"/>
		</div>
		`, base64.StdEncoding.EncodeToString(out)),
	)
}

func LampSetBySchedule(w http.ResponseWriter, r *http.Request) {
	out, err := exec.Command("/home/dietpi/Code/lamp/lamp.py", "set-by-schedule").Output()
	if err != nil {
		renderPage(w, err.Error())
		return
	}
	renderPage(w, string(out))
}

func renderPage(w http.ResponseWriter, s string) {
	fmt.Fprintf(w, `
	<!DOCTYPE html>
	<html>
	<body>
	%s
	<hr>
	<div>
	<a href="off">turn off</a> |
	<a href="on">turn on</a> |
	<a href="set-by-schedule">set by schedule</a> |
	<a href="pic">view pic</a> |
	<a href="/">view status</a>
	</div>
	</body>
	</html>
	`, s)
}
