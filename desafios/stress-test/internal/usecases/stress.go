package usecases

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)


type StatusCode map[string]int

type HTTPStats struct {
	Total int64 `json:"total"`
	StatusCode []StatusCode
	ExecutionTimeMcs int64 `json:"execution-time-microseconds"`
}
type DateStats map[string]HTTPStats

type StressTestURL struct{
	URL string
	Requests int64
	Concurrency int64
	arr []StatusCode
	mapper map[string]int 
}

func NewStressURL(url string, req, conc int64) *StressTestURL{
	 return &StressTestURL{
		URL: "http://google.com",
		Requests: req, 
		Concurrency: conc,
		arr: []StatusCode{},
		mapper: map[string]int{},
	}
}

var (
	wg sync.WaitGroup
) 

func(s *StressTestURL) Aqui() error{
	httpResp := make(chan int64, s.Concurrency)
	
	start := time.Now()
	for i := 0; i < int(s.Requests); i++ {
		wg.Add(1)
		go s.callURL(  httpResp)
	}
	elapsed := time.Since(start).Microseconds()

	defer wg.Wait()
	done := make(chan bool)

	i :=int64(0)
	go func(){
		for {
			select{
			case status, _ := <- httpResp:
				s.groupData(status)
				atomic.AddInt64(&i, 1)
				if i == s.Requests {
					done <- true
					return
				}
			}
		}
	}()
	<- done
	return s.writeFile(elapsed)
}

func(s *StressTestURL) callURL( httpCh chan int64){
	resp, err := 	http.DefaultClient.Get(s.URL)
	if err != nil {
		fmt.Println(err)
	}
	httpCh <- int64(resp.StatusCode)
	defer wg.Done()
}


func(s *StressTestURL) groupData(status int64){
	strStatus := strconv.FormatInt(status, 10)
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


func(s *StressTestURL) writeFile(executionTime int64) error {
  // Open the JSON file
    file, err := os.OpenFile("data.txt", os.O_APPEND|os.O_CREATE | os.O_WRONLY, 0644)
    if err != nil {
        return errors.Join(errors.New("Error opening file:"), err)
    }
    defer file.Close()

	newStats := HTTPStats{
		Total: s.Requests,
		StatusCode: s.arr,
		ExecutionTimeMcs: executionTime,
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