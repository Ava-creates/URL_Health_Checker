package main

import (
"os"
"bufio"
)
type Job struct{
	ID int
	URL string
}
func main(){
	data, err := os.Open("./urls.txt")
	if err!=nil{
		panic("AAAA")
	}
	i:=1
	results := make(chan check, 100)
	go startWriter(results, "output.json")
	wp := Newworkerpool(8, 100, results)
	wp.Start()
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		url := scanner.Text()
		// fmt.Println(url)
		j:= Job{ID:i, URL:url}
		wp.Submit(j)
		i+=1
	}
	wp.Shutdown()  
	close(results)

	// c := make(chan check)
	// url := "https://github.com"
	// go checker_(url, c)
	// z:= <-c
	// fmt.Println(z)
}