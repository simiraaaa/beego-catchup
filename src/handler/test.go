package handler

import (
	"io/ioutil"
	"net/http"
)

// Test ... 適当なページを表示
func Test(w http.ResponseWriter, r *http.Request) {
	res, err := http.Get(`https://ja.wikipedia.org/wiki/%E3%83%86%E3%82%B9%E3%83%88`)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
}
