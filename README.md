# Concurrent URL Health Checker
 
A concurrent URL health checking system built in Go. Reads a list of URLs, distributes health checks across a configurable worker pool, and writes structured results to `output.json`.
 
## Architecture
 
```
[urls.txt] --> [main.go] --> [Job Queue] --> [Worker Pool] --> [Writer] --> [output.json]
                             (chan Job)      (goroutines)     (goroutine)
```
 
- **Job Queue** — a buffered channel that decouples job submission from job processing. `main.go` submits jobs without needing to know if a worker is ready
- **Worker Pool** — N goroutines running concurrently, each pulling jobs off the queue and calling `checker_` to make an HTTP request
- **Writer** — a single dedicated goroutine that receives results and appends JSON lines to `output.json`, avoiding file race conditions without needing a mutex
## Project Structure
 
```
URL_checker/
├── main.go           # Reads urls.txt, wires everything together, shuts down cleanly
├── checker.go        # Makes HTTP GET request, returns a check result struct
├── worker.go         # Worker pool, graceful shutdown
├── writer.go         # Single writer goroutine, writes results to output.json
├── urls.txt          # Input — one URL per line
└── README.md
```
 
## Usage
 
```bash
go run .
```
 
Results are written to `output.json`. Each line is a JSON object:
```json
{"URL":"https://google.com","STATUS":200,"LATENCY":42,"HEALTHY":true}
{"URL":"https://httpstat.us/500","STATUS":500,"LATENCY":120,"HEALTHY":false}
{"URL":"https://darby.thisdomaindoesnotexist.xyz","STATUS":-1,"LATENCY":0,"HEALTHY":false}
```
 
