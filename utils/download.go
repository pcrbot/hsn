package utils

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"github.com/dustin/go-humanize"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

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

func Download(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header["User-Agent"] = []string{"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.116 Safari/537.36 Edg/83.0.478.61"}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if strings.Contains(resp.Header.Get("Content-Encoding"), "gzip") {
		buffer := bytes.NewBuffer(body)
		r, _ := gzip.NewReader(buffer)
		defer r.Close()
		unCom, err := ioutil.ReadAll(r)
		return unCom, err
	}
	return body, nil
}

type WriteCounter struct {
	Data  []byte
	Total uint64
}

func (wc *WriteCounter) Write(p []byte) (int, error) {
	n := len(p)
	wc.Total += uint64(n)
	wc.Data = append(wc.Data, p...)
	wc.PrintProgress()
	return n, nil
}

func (wc WriteCounter) PrintProgress() {
	fmt.Printf("\r%s", strings.Repeat(" ", 35))
	fmt.Printf("\rDownloading... %s complete", humanize.Bytes(wc.Total))
}

func DownloadFileWithProgress(filepath string, url string) error {
	out, err := os.Create(filepath + ".tmp")
	if err != nil {
		return err
	}
	resp, err := http.Get(url)
	if err != nil {
		out.Close()
		return err
	}
	defer resp.Body.Close()
	counter := &WriteCounter{}
	counter.Data = []byte{}
	if _, err = io.Copy(out, io.TeeReader(resp.Body, counter)); err != nil {
		_ = out.Close()
		return err
	}
	_ = out.Close()
	fmt.Print("\n")
	if err = os.Rename(filepath+".tmp", filepath); err != nil {
		return err
	}
	return nil
}

func DownloadWithProgress(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	counter := &WriteCounter{}

	_ = io.TeeReader(resp.Body, counter)
	return counter.Data, nil
}
