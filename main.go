package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"
)

var lport string
var dip string

func main() {

	lport = os.Getenv("LPORT")
	dip = os.Getenv("DIP")

	if lport == "" {
		lport = os.Args[1]
	}

	if dip == "" {
		dip = os.Args[2]
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		b, _ := os.ReadFile("index.html")
		fmt.Fprint(w, string(b))
	})

	// Anything we don't do in Go, we pass to the old platform
	http.HandleFunc("/wordpress/", wordpress) //  ex. /wordpress/xyxshxpxchxy/photo3235

	fmt.Println("Listening on localhost:" + lport)
	fmt.Println("Forward to:" + dip)

	// Start the server
	//http.ListenAndServe(":"+getEnvVariable("PORT"), nil)
	http.ListenAndServe(":"+lport, nil)
}

func wordpress(w http.ResponseWriter, r *http.Request) {

	// change the request host to match the target
	vars := strings.Split(r.URL.Path, "/")

	str := vars[2]

	str = strings.Replace(str, "xch", "4", -1)
	str = strings.Replace(str, "xsh", "6", -1)
	str = strings.Replace(str, "xha", "8", -1)
	str = strings.Replace(str, "xse", "0", -1)
	str = strings.Replace(str, "xy", "1", -1)
	str = strings.Replace(str, "xd", "2", -1)
	str = strings.Replace(str, "xs", "3", -1)
	str = strings.Replace(str, "xp", "5", -1)
	str = strings.Replace(str, "xh", "7", -1)
	str = strings.Replace(str, "xn", "9", -1)

	//fmt.Println("Forward to: " + "http://" + dip + ":" + str)

	//u, _ := url.Parse("http://" + getEnvVariable("DIP") + ":" + vars[2])
	u, _ := url.Parse("http://" + dip + ":" + str)
	pro := httputil.NewSingleHostReverseProxy(u)

	//r.URL.Path = "/"

	pro.ServeHTTP(w, r)
}
