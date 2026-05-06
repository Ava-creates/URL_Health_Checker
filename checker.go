package main

import (
	"net/http"
	"time"
)

type check struct {
	URL string
	STATUS int
	LATENCY int64
	HEALTHY bool
}

func checker_(url string) check{
	start := time.Now()
	resp, err := http.Get(url)

	end := time.Since(start)
	c := check{}
	c.URL = url
	c.HEALTHY = false
	c.LATENCY =  end.Milliseconds()
	if err==nil{
		c.STATUS = resp.StatusCode
		if resp.StatusCode>=200 && resp.StatusCode<300{
			c.HEALTHY = true}
	}else{
			c.STATUS = -1
			c.HEALTHY = false
	}
	return c
}

