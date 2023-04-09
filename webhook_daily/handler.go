package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"time"

	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	Name string `json:"name"`
}

type Field struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline bool   `json:"inline"`
}

type Embed struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Fields      []Field `json:"fields"`
}

type RequestJson struct {
	Embeds []Embed `json:"embeds"`
}

func total(date time.Time) string {
	url := fmt.Sprintf(
		"https://paylogger.an.r.appspot.com/total?year=%d&month=%d&date=%d",
		date.Year(),
		date.Month(),
		date.Day(),
	)

	fmt.Printf("url is %s", url)

	res, err := http.Get(
		url,
	)
	if err != nil {
		return "-"
	}
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)

	return string(body)

}

func test(ctx context.Context, name MyEvent) (string, error) {
	webHook := "https://discord.com/api/webhooks/1066294249043263578/Qp62DbrZN7SVkA2syjHQm4SjJVVE1FtwkavTkk3b9Lsd4ZMD-IZZExXJn8sWYHGffOfL"

	data := RequestJson{
		[]Embed{
			{
				Title:       fmt.Sprintf("Today's Amount is %s", total(time.Now())),
				Description: "You used these money.",
				Fields: []Field{
					{
						Name:  "Test",
						Value: "500",
					},
				},
			},
		},
	}

	// encode json
	data_json, _ := json.Marshal(data)
	fmt.Printf("[+] %s\n", string(data_json))

	// send json
	res, err := http.Post(
		webHook,
		"application/json",
		bytes.NewBuffer(data_json),
	)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	if err != nil {
		fmt.Println("[!] " + err.Error())
	} else {
		fmt.Println("[*] " + res.Status)
	}

	return "Success", nil
}

func main() {
	lambda.Start(test)
}
