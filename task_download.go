package main

import "time"

type DownloadTask struct {
	URL string
}

func (d *DownloadTask) Proccess() string {
	// this is the function that implement Process from the interface
	time.Sleep(10 * time.Second) // simulate heavy or process that need more time
	return "Download From " + d.URL
}
