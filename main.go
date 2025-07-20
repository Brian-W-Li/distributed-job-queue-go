package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

type Job struct {
	ID      int
	Payload string `json:"payload"`
	JobType string
	//Created time.Time
}

var jobqueue = make(chan Job, 10)
var jobCounter int = 0
var mu sync.Mutex

func main() {
	for i := 0; i < 3; i++ {
		go func() {
			for job := range jobqueue {
				fmt.Printf("Worker processing job %d, %s: %s\n", job.ID, job.JobType, job.Payload)
				duration := DoWork(job)
				fmt.Printf("Worker Job %d, %s completed after %d seconds\n", job.ID, job.JobType, duration)
			}
		}()
	}

	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/enqueue", enqueueHandler)

	fmt.Println("Server starting")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))

}

// Takes in a responseWriter
func enqueueHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:

		payload := r.FormValue("payload")
		jobType := r.FormValue("jobType")
		// type requestPayload struct {
		// 	Payload string `json:"payload"`
		// }
		// var req requestPayload
		// json.NewDecoder(r.Body).Decode(&req) //reads the json body of the post
		//request and decodes the value into req.Payload
		//make sure to http request has payload:something

		mu.Lock()
		jobCounter++
		job := Job{
			ID:      jobCounter,
			Payload: payload,
			JobType: jobType,
			//Created: time.Now(),
		}
		mu.Unlock()

		jobqueue <- job

		http.Redirect(w, r, "/", http.StatusSeeOther)
	case http.MethodGet:
		fmt.Fprintln(w, "This is the GET version of /enqueue. Maybe you're just checking?")
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

}

func DoWork(job Job) int {

	var sleepTime time.Duration
	if job.JobType == "default" {
		sleepTime = 10 * time.Second
	} else if job.JobType == "encrypt" {
		sleepTime = 7 * time.Second
	} else if job.JobType == "compress" {
		sleepTime = 4 * time.Second
	} else if job.JobType == "hash" {
		sleepTime = 1 * time.Second
	}

	time.Sleep(sleepTime)

	return int(sleepTime.Seconds())
}
