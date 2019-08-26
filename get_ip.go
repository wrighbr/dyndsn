package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)




func getIP() (string) {

	resp, err := http.Get("https://ipecho.net/plain")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	ipaddr := string(body)
	
	return ipaddr
}
