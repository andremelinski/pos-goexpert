package usecases

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sync"
)


type StressTestInput struct{
	URL string
	Requests uint64
	Concurrency uint64
}



func NewStressURL(){
}

var (
	wg sync.WaitGroup
) 

func Aqui(){
	conc := 2
	call := 5
	// controla a quantidade que ocorre em concorrencia
	// esse channel suporta ate o valor que esta em conc dentro dele
	errorControlCh := make(chan struct{}, 1)
	httpResp := make(chan int, conc)

	// for{ 
	for i := 0; i < call; i++ {
		wg.Add(1)
		go callURL("http://google.com",   errorControlCh, httpResp)
	}

	defer wg.Wait()

	go func(){
		for {
			select{
			case msg := <- httpResp:
				// if ok{
					fmt.Printf("Http status %d\n", msg)
				// }
				break
			}
		}
	}()
	writeFile()
}

// func retryErrorFileUpload(errorFileUpload chan struct{}){
// 	for{
// 		select {
// 		case fileName, ok := <- errorFileUpload:
// 			if ok {
// 				wg.Add(1)
// 				// para retentar, deve sinalizar 
// 				uploadControl <- struct{}{}
			
// 				go callURL( fileName,  "s3-sqs-ts-bucket-dev-321123", uploadControl, errorFileUpload)
// 			}
// 		}
// 	}
// }


func callURL(url string, errorControlCh chan struct{}, httpCh chan int){
	resp, err := 	http.DefaultClient.Get(url)

	if err != nil {
		// <-errorControlCh
		fmt.Println(err)
	}
		// fmt.Println(resp)
	// fmt.Println(resp.StatusCode)
	httpCh <- resp.StatusCode
	
	defer wg.Done()
}

type HTTPStats struct {
	Total int `json:"total"`
	Code200 int `json:"200,omitempty"`
	Code400 int `json:"400,omitempty"`
	Code500 int `json:"500,omitempty"`
	Time int `json:"time"`
}
type DateStats map[string]map[string]HTTPStats

func writeFile(){
  // Open the JSON file
    file, err := os.OpenFile("data.txt", os.O_APPEND|os.O_CREATE | os.O_WRONLY, 0644)
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer file.Close()

	newStats := HTTPStats{
		Total: 15,
		Code200: 3,
		Time: 3000,
	}

	data :=  DateStats{
		"30-06-2024": {
			"http://example.com": newStats,
		},
	}

	updatedJsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}

	str := string(updatedJsonData)+",\n"

	_, err = file.Write([]byte(str))
	
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}
}