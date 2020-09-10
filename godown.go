package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
)

// Download type struct
type Download struct {
	Url           string
	TargetPath    string
	TotalSections int
}

func main() {
	startTime := time.Now()

	download := Download{
		Url:           os.Args[1],
		TargetPath:    os.Args[2],
		TotalSections: 10,
	}

	err := download.Do()
	if err != nil {
		log.Fatalf("Error downloading file: %s\n", err)
	}
	fmt.Printf("Download completed in %v seconds\n", time.Since(startTime).Seconds())
}

// Do the actual download
func (download Download) Do() error {
	fmt.Println("Making connection...")
	req, err := download.getNewRequest("HEAD")
	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	fmt.Printf("Got %v\n", resp.StatusCode)
	if resp.StatusCode > 299 {
		return errors.New(fmt.Sprintf("Can't process, response is %v", resp.StatusCode))
	}

	size, err := strconv.Atoi(resp.Header.Get("Content-Length"))
	if err != nil {
		return err
	}

	fmt.Printf("Size is %v bytes.\n", size)

	// Make sections array with values for start and end of sections
	// If file is 100 bytes and section size is 10 then [[0, 10][11, 20]...[91, 99]]
	var sections = make([][2]int, download.TotalSections)
	eachSize := size / download.TotalSections

	for i := range sections {
		if i == 0 {
			sections[i][0] = 0
		} else {
			sections[i][0] = sections[i-1][1] + 1
		}
		if i < download.TotalSections-1 {
			sections[i][1] = sections[i][0] + eachSize
		} else {
			sections[i][1] = size - 1
		}
	}

	var wg sync.WaitGroup
	for i, section := range sections {
		// Increment waitgroup counter by 1
		wg.Add(1)
		// Capture the values because they will keep change
		i := i
		section := section
		go func() {
			// Waits for section to download, decrements counter by one
			defer wg.Done()
			err := download.downloadSection(i, section)
			if err != nil {
				panic(err)
			}
		}()
	}
	// Waits until all sections have been downloaded
	wg.Wait()
