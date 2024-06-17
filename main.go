package main

import (
	"fmt"
)

func main() {

	// task defined in heere will be executed in worker pool
	taskLists := []Task{
		&EmailTask{
			EmailTo:     "john.doe@unkwnonmail.com",
			Subject:     "Hi John",
			MessageBody: "How are you John?",
		},
		&DownloadTask{
			URL: "https://download.john.doe.com",
		},
		&DownloadTask{
			URL: "http://fresh.download.com",
		},
	}

	wp := WorkerPool{
		Tasks:       taskLists,
		concurrency: 10, // worker
	}

	// run the worker
	wp.Run()

	// show the result
	for _, res := range wp.Results {
		fmt.Printf("response: %v\n", res.TextResult)
	}

	fmt.Println("All task already executed")
}
