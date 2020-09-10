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
