package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Channel struct {
	Items []Item `xml:"item"`
}

type Item struct {
	Title string `xml:"title"`
	Link  string `xml:"link"`
}

func getRSSData() ([]string, error) {
	resp, err := http.Get("http://pstgu.yss.su/iu9/mobiledev/lab4_yandex_map/?x=var01")
	fmt.Println(resp)
	err_lst := []string{}
	if err != nil {
		return err_lst, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err_lst, err
	}
	fmt.Println(string(body))

	var rss []string
	err = json.Unmarshal(body, &rss)
	if err != nil {
		return err_lst, err
	}
	fmt.Println(rss)
	return rss, nil
}

func getRSSPage(w http.ResponseWriter, r *http.Request) {
	rss, err := getRSSData()
	fmt.Println(rss)
	if err != nil {
		http.Error(w, "Error fetching RSS data", http.StatusInternalServerError)
		return
	}
	for _, item := range rss {
		fmt.Fprintf(w, "%s\n", item)
	}
}

func main() {
	http.HandleFunc("/", getRSSPage)

	err := http.ListenAndServe("127.0.0.1:3333", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
