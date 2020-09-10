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

