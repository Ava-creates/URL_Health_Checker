package main


import "sync"
type workerPool struct{
	num int
	jobs chan Job
	result chan check
	wg sync.WaitGroup
}

func Newworkerpool(num int, queue int , result chan check) *workerPool{
	return &workerPool{num: num, jobs: make(chan Job, queue), result: result}

}

func (wp *workerPool) Submit(job Job){
	wp.jobs <- job
}
func (wp *workerPool) Start(){
	for i := 0; i < wp.num; i++{
		wp.wg.Add(1)
        go wp.worker(i)
	}
}

func (wp *workerPool) worker(id int){
	defer wp.wg.Done()
    
    for job := range wp.jobs {
        result := checker_(job.URL)
        wp.result <- result
    }

}

func (wp *workerPool) Shutdown() {
    close(wp.jobs)  
    wp.wg.Wait()   
}