package helper

import (
	"time"
)

const (
	pingTimeMax = time.Second
)

var (
	addresses []string
	inFile    = "in.txt"
	outFile   = "out.txt"
	urls      []Url
)
