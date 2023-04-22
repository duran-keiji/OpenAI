package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"io"
	"io/ioutil"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// -------------------------- リクエストの構造体 --------------------------
type RequestBody struct {
	Model    string    `json:"model"`   
	Messages []Messages `json:"messages"`
}

type Messages struct {
	Role    string `json:"role"`   
	Content string `json:"content"`
}
// ---------------------------------------------------------------------

// -------------------------- レスポンスの構造体 --------------------------
type Response struct {
	ID      string   `json:"id"`     
	Object  string   `json:"object"` 
	Created int64    `json:"created"`
	Choices []Choice `json:"choices"`
	Usage   Usage    `json:"usage"`  
}

type Choice struct {
	Index        int64   `json:"index"`        
	Message      Message `json:"message"`      
	FinishReason string  `json:"finish_reason"`
}

type Message struct {
	Role    string `json:"role"`   
	Content string `json:"content"`
}

type Usage struct {
	PromptTokens     int64 `json:"prompt_tokens"`    
	CompletionTokens int64 `json:"completion_tokens"`
	TotalTokens      int64 `json:"total_tokens"`     
}
// ---------------------------------------------------------------------

func get_respons(request events.APIGatewayProxyRequest) (string) {
	url := "https://api.openai.com/v1/chat/completions"
	apiKey := os.Getenv("API_KEY")
	searchQuery := request.QueryStringParameters["q"]

	// リクエストパラメータ
	requestBody := RequestBody{
		Model: "gpt-3.5-turbo",
		Messages: []Messages{
			Messages{
				Role: "user",
				Content: searchQuery,
			},
		},
	}

	// requestBody構造体をjsonに変換
	requestBodyJSON, err := json.Marshal(requestBody)
	if err != nil {
		panic(err)
	}

	// HTTPリクエストの作成
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBodyJSON))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	// HTTPリクエストを送信
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(resp.Body)

	// レスポンスの中身を格納
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	return string(body)
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	respons := get_respons(request)

	return events.APIGatewayProxyResponse{
		Body:       fmt.Sprintf(string(respons)),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}