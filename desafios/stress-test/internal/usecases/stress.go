package usecases

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)


type StatusCode map[string]int

type HTTPStats struct {
	URL string `json:"url"`
	Total int64 `json:"total"`
	StatusCode []StatusCode
	ExecutionTimeMs int64 `json:"execution-time-miliseconds"`
}
type DateStats map[string]HTTPStats

type StressTestURL struct{
	URL string
	Requests int64
	Concurrency int64
	arr []StatusCode
	mapper map[string]int
	ExecutionMs int64
}

func NewStressURL(url string, req, conc int64) *StressTestURL{
	 return &StressTestURL{
		URL: url,
		Requests: req, 
		Concurrency: conc,
		arr: []StatusCode{},
		mapper: map[string]int{},
		ExecutionMs: 0,
	}
}

var (
	wg sync.WaitGroup
) 

type HttpInfo struct{
	HTTPStats int64
	callDuration int64
}

func(s *StressTestURL) Stress() (*string, error){
	httpResp := make(chan HttpInfo, s.Concurrency)
	
	for i := 0; i < int(s.Requests); i++ {
		wg.Add(1)
		go s.callURL(  httpResp)
	}

	defer wg.Wait()
	done := make(chan bool)

	i :=int64(0)
	go func(){
		for {
			select{
			case status, _ := <- httpResp:
				s.groupData(status)
				atomic.AddInt64(&i, 1)
				s.ExecutionMs += status.callDuration
				if i == s.Requests {
					done <- true
					return
				}
			}
		}
	}()
	<- done
	s.writeFile()
	return s.readFile()
}

func(s *StressTestURL) callURL( httpCh chan HttpInfo){
	start := time.Now()
	resp, err := 	http.DefaultClient.Get(s.URL)
	elapsed := time.Since(start).Milliseconds()
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	
	httpCh <- HttpInfo{
	int64(resp.StatusCode),
	elapsed,
	}
	defer wg.Done()
}


func(s *StressTestURL) groupData(status HttpInfo){
	strStatus := strconv.FormatInt(status.HTTPStats, 10)
	// s.mapper[strStatus] ==0 ocorre para um novo item no mapper e na 1 insercao. Erro: [{200:0}, {200:1}]
	if s.mapper[strStatus] ==0 {
		// cria uma posicao a frente para nao duplicar
		s.mapper[strStatus] = len(s.arr)+1
		statusCode:= StatusCode{
				strStatus: 1,
			}
			s.arr = append(s.arr, statusCode)
	}else{
		statusCodeIndexArr := s.mapper[strStatus]-1
		s.arr[statusCodeIndexArr][strStatus]++
	}
}


func(s *StressTestURL) writeFile()  error {
  // Open the JSON file
    file, err := os.OpenFile("data.txt", os.O_APPEND|os.O_CREATE | os.O_WRONLY, 0644)
    if err != nil {
        return errors.Join(errors.New("Error opening file:"), err)
    }
    defer file.Close()

	newStats := HTTPStats{
		URL: s.URL,
		Total: s.Requests,
		StatusCode: s.arr,
		ExecutionTimeMs: s.ExecutionMs,
	}

	updatedJsonData, err := json.MarshalIndent(newStats, "", "  ")
	if err != nil {
		return errors.Join(errors.New("Error encoding JSON:"), err)
	}

	str := string(updatedJsonData)+",\n"

	_, err = file.Write([]byte(str))
	
	if err != nil {
		return errors.Join(errors.New("Error writing into file:"), err)
	}
	return nil
}

func(s *StressTestURL) readFile() (*string, error){
	  file, err := os.OpenFile("data.txt", os.O_RDONLY, 0644)
    if err != nil {
        return nil, err
    }
    defer file.Close()
	byteContent, err := io.ReadAll(file)

	if err != nil {
        return nil, err
    }
	strByte := string(byteContent)
	return &strByte, err
}