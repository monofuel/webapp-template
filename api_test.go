package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
	"time"
)

var host string
var client *http.Client = &http.Client{}

func TestIndex(t *testing.T) {
	go main()
	//give it a few seconds to start
	time.Sleep(500 * time.Millisecond)

	host = fmt.Sprintf("http://localhost:%v", webPort)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	req, err := http.NewRequest("GET", host, nil)
	req.WithContext(ctx)
	resp, err := client.Do(req)
	if err != nil {
		t.Error(err)
	}
	if resp.StatusCode != 200 {
		t.Errorf("bad status code:%d\n", resp.StatusCode)
	}

	defer resp.Body.Close()
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	index := buf.String()
	fmt.Println(index)
}

func TestRandom(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	endpoint := fmt.Sprintf("%v/v1/test/random", host)
	req, err := http.NewRequest("GET", endpoint, nil)
	req.WithContext(ctx)
	resp, err := client.Do(req)
	if err != nil {
		t.Error(err)
	}
	if resp.StatusCode != 200 {
		t.Errorf("bad status code:%d\n", resp.StatusCode)
	}

	defer resp.Body.Close()
	var body struct {
		Result int `json:"result"`
	}
	json.NewDecoder(resp.Body).Decode(&body)
	if body.Result == 0 {
		t.Errorf("bad random api number\n")
	}
	fmt.Printf("%v responded with %d", endpoint, body.Result)
}
