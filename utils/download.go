package utils

import (
	"fmt"
	"github.com/spf13/afero"
	"io/ioutil"
	"net/http"
	"os"
)

var aferoFs = afero.NewMemMapFs()

func DownloadFile(url string, path string) (err error) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = ioutil.WriteFile(path, body, os.ModePerm)
	return
}

func Download(url string) (data []byte, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	data, err = ioutil.ReadAll(resp.Body)
	return
}
