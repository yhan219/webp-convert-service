package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os/exec"
	"strings"
)

func main() {
	const bindAddress = ":80"
	http.HandleFunc("/", requestHandler)
	fmt.Println("Http server listening on", bindAddress)
	_ = http.ListenAndServe(bindAddress, nil)
}

func requestHandler(response http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/" {
		response.WriteHeader(http.StatusNotFound)
		return
	}
	var imgUrl string
	segments := make([]string, 0)
	query := request.URL.Query()
	for q := range query {
		key := q
		element := request.FormValue(key)
		if key == "url" {
			imgUrl = element
		} else if element == "true" {
			segments = append(segments, fmt.Sprintf("-%v", key))
		} else {
			segments = append(segments, fmt.Sprintf("-%v", key), fmt.Sprintf("%v", element))
		}
	}
	if imgUrl=="" {
		response.WriteHeader(http.StatusNotFound)
		return
	}
	u, err := url.QueryUnescape(imgUrl)
	if err != nil {
		response.WriteHeader(http.StatusNotFound)
		log.Println("imageUrl is error")
		return
	}
	resp, err := http.Get(strings.TrimRight(u, "/"))
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()

	segments = append(segments, "-o", "-", "--", "-")
	cmd := exec.Command("cwebp", segments...)
	cmd.Stdin = io.Reader(resp.Body)
	cmd.Stdout = response
	_ = cmd.Start()
	defer cmd.Wait()

	response.Header().Set("Content-Type", "image/webp")
	response.WriteHeader(http.StatusOK)
}
