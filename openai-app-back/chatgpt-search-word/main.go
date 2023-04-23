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
	RequestModel    string    `json:"model"`   
	RequestMessages []RequestMessages `json:"messages"`
}

type RequestMessages struct {
	RequestRole    string `json:"role"`   
	RequestContent string `json:"content"`
}
// ---------------------------------------------------------------------

// -------------------------- レスポンスの構造体 --------------------------
type Response struct {
	ResponseID      string   `json:"id"`     
	ResponseObject  string   `json:"object"` 
	ResponseCreated int64    `json:"created"`
	ResponseChoices []ResponseChoices `json:"choices"`
	ResponseUsage   ResponseUsage    `json:"usage"`  
}

type ResponseChoices struct {
	ResponseIndex        int64   `json:"index"`        
	ResponseMessage      ResponseMessage `json:"message"`      
	ResponseFinishReason string  `json:"finish_reason"`
}

type ResponseMessage struct {
	ResponseRole    string `json:"role"`   
	ResponseContent string `json:"content"`
}

type ResponseUsage struct {
	ResponsePromptTokens     int64 `json:"prompt_tokens"`    
	ResponseCompletionTokens int64 `json:"completion_tokens"`
	ResponseTotalTokens      int64 `json:"total_tokens"`     
}
// ---------------------------------------------------------------------

func get_respons(request events.APIGatewayProxyRequest) (string) {
	url := "https://api.openai.com/v1/chat/completions"
	apiKey := os.Getenv("API_KEY")
	searchQuery := request.QueryStringParameters["q"]

	// リクエストパラメータ
	requestBody := RequestBody{
		RequestModel: "gpt-3.5-turbo",
		RequestMessages: []RequestMessages{
			RequestMessages{
				RequestRole: "user",
				RequestContent: searchQuery,
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
	fmt.Sprintf("body=", string(body))


	var response Response
	json.Unmarshal(body, &response)
	fmt.Sprintf("response=", response)


	message, err := json.Marshal(response.ResponseChoices[0].ResponseMessage.ResponseContent)
	if err != nil {
		panic(err)
	}	

	return string(message)
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	respons := get_respons(request)

    headers := map[string]string{
        "Content-Type":           "application/json",
        "Access-Control-Allow-Origin": "*",
        "Access-Control-Allow-Headers": "Content-Type,X-Amz-Date,Authorization,X-Api-Key,X-Amz-Security-Token",
        "Access-Control-Allow-Methods": "DELETE,GET,HEAD,OPTIONS,PATCH,POST,PUT",
		"Access-Control-Allow-Credential": "true",
    }

	return events.APIGatewayProxyResponse{
		Body:       fmt.Sprintf(string(respons)),
		StatusCode: 200,
		Headers: 	headers,
	}, nil
}

func main() {
	lambda.Start(handler)
}