package main

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fatih/color"
)

//d1 nano 311016
//d2 nano 232787

//duration http 1000p 2484766
//duration fasthttp 1000p 2455488

func main() {

	wait := make(chan bool, 1)

	// urlStr := "http://localhost:8000/"

	// url, err := url.Parse(urlStr)

	// if err != nil {
	// 	log.Println(err)
	// }

	// //creating the proxyURLs

	// o := 12000
	// oPTR := &o
	// mutex := sync.Mutex{}

	// for j := 12000; j < 18000; j++ {

	// 	go func() {

	// 		mutex.Lock()

	// 		proxyStr := "http://localhost:" + strconv.Itoa(*oPTR)
	// 		mutex.Unlock()
	// 		log.Println(proxyStr)

	// 		proxyURL, err := url.Parse(proxyStr)

	// 		if err != nil {
	// 			log.Println(err)
	// 		}
	// 		transport := &http.Transport{
	// 			Proxy: http.ProxyURL(proxyURL),
	// 		}

	// 		client := &http.Client{
	// 			Transport: transport,
	// 		}
	// 		request, err := http.NewRequest("GET", url.String(), nil)
	// 		if err != nil {
	// 			log.Println(err)
	// 		}

	// 		_, err2 := client.Do(request)
	// 		if err2 != nil {
	// 			log.Println(err2)
	// 		} else {
	// 			fmt.Println("Req Sent")
	// 		}

	// 		//getting the response
	// 		// data, err := ioutil.ReadAll(response.Body)
	// 		// if err != nil {
	// 		// 	log.Println(err)
	// 		// }
	// 		// //printing the response
	// 		// log.Println(string(data))
	// 		// defer response.Body.Close()

	// 	}()

	// 	if *oPTR == (18000 - 1) {
	// 		<-wait

	// 	}
	// 	mutex.Lock()
	// 	*oPTR++
	// 	mutex.Unlock()
	// }

	//creating the URL to be loaded through the proxy

	//adding the proxy settings to the Transport object

	//adding the Transport object to the http Client

	//generating the HTTP GET request

	//calling the URL

	NumberOfRequests := 3800

	log.Println(NumberOfRequests)

	var NumberOfResponseErrors int
	NumberOfResponseErrorsPtr := &NumberOfResponseErrors

	var NumberOfReadErrors int
	NumberOfReadErrorsPtr := &NumberOfReadErrors
	time1 := time.Now().UnixNano()

	mutex := sync.Mutex{}

	for i := 0; i < NumberOfRequests; i++ {

		//time.Sleep(1 * time.Millisecond)

		go func() {

			_, err := http.Get("http://localhost:8000")
			if err != nil {

				mutex.Lock()

				*NumberOfResponseErrorsPtr++

				mutex.Unlock()
				//	color.Red("Error: ", err)
				return

			}

			// BodyBytes, err2 := ioutil.ReadAll(resp.Body)
			// defer resp.Body.Close()
			// if err2 != nil {

			// 	mutex.Lock()
			// 	*NumberOfReadErrorsPtr++
			// 	mutex.Unlock()
			// 	//color.Red("Error reading body: ", err2)
			// 	return
			// }
			// color.Yellow(string(BodyBytes) + "\n")
		}()

		if i == (NumberOfRequests - 1) {

			time2 := time.Now().UnixNano()
			time.Sleep(10 * time.Second)
			duration := (time2 - time1) % 1000000000
			color.Green("Duration: ", duration)
			color.HiCyan("Number of Resp errors: ", *NumberOfResponseErrorsPtr)
			color.HiCyan("Number of Read Errors: ", *NumberOfReadErrorsPtr)
			<-wait

		}
	}

}
