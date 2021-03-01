package helper

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func loadFile() {
	stringNotification("Loading urls...")
	data, err := ioutil.ReadFile(inFile)
	errFatalNotification(err)
	addresses = strings.Split(string(data), "\n")
	stringNotification("Successfully completed.")
}

func saveFile() {
	stringNotification("Saving urls...")
	file, err := os.Create(outFile)
	errFatalNotification(err)
	for _, url := range urls {
		writeUrlToFile(file, &url)
	}
	errNotification(file.Close())
	stringNotification("Successfully completed.")
}

func writeStringToFile(file *os.File, s string) {
	_, err := file.WriteString(s)
	if err != nil {
		panic(err)
	}
}

func writeUrlToFile(file *os.File, url *Url) {
	writeStringToFile(file, url.ip)
	if len(url.port) > 0 {
		writeStringToFile(file, fmt.Sprintf(":%s", url.port))
	}
	writeStringToFile(file, "\n")
}
