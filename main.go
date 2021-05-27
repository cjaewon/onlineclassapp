package main

import (
	"embed"
	"os/exec"
	"runtime"

	"github.com/webview/webview"
)

//go:embed web/public/*
var public embed.FS

func main() {

	debug := true
	w := webview.New(debug)

	defer w.Destroy()

	w.SetTitle("시간표앱")
	w.SetSize(800, 600, webview.HintNone)

	bytes, err := public.ReadFile("web/public/index.html")
	if err != nil {
		panic(err)
	}

	w.Navigate(`data:text/html` + string(bytes))

	bytes, err = public.ReadFile("web/public/build/bundle.js")
	if err != nil {
		panic(err)
	}

	w.Eval(string(bytes))

	bytes, err = public.ReadFile("web/public/build/bundle.css")
	if err != nil {
		panic(err)
	}

	w.Eval(`
	const style = document.createElement('style');
	document.head.append(style);
	style.textContent = '` + string(bytes) + `'`)

	w.Bind("openLink", openLink)

	w.Run()
}

func openLink(link string) error {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start"}
	case "darwin":
		cmd = "open"
	default: // "linux", "freebsd", "openbsd", "netbsd"
		cmd = "xdg-open"
	}
	args = append(args, link)
	return exec.Command(cmd, args...).Start()
}
