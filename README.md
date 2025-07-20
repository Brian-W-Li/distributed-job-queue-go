# distributed-job-queue-go
A concurrent job queue system in Go with HTTP endpoints and worker simulation

-HTTP server built with Go's net/http package
can submit a job with 'payload' and 'jobType' via form
worker goroutines retrieve jobs from shared queue
you can also check the current status of the job queue

# getting started

go run main.go
then visit http://localhost:8080/ for job submission form
and (todo) visit http://localhost:8080/status for current status
