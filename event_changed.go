package main

import (
	"bytes"
	"howett.net/plist"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type currentEventData struct {
    Name	string	`plist:"name"`
    CityID	int	`plist:"city_id"`
    StartTime	float64	`plist:"start_time"`
    EndTime	float64	`plist:"end_time"`
    Reward1	string	`plist:"reward_1"`
    Reward2	string	`plist:"reward_10"`
    Reward3	string	`plist:"reward_100"`
    Reward4	string	`plist:"reward_1000"`
    MinJobs	string	`plist:"min_jobs"`
    Hash	string	`plist:"hash"`
}

func main() {
	// Load event data over the internet
	resp, err := http.Get("http://ppupdate.nimblebit.com/event/current")

	if err != nil {
		log.Fatal("Error loading current event", err)
	}

	plistBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	buffReader := bytes.NewReader(plistBytes)

	var currentEvent currentEventData
	decoder := plist.NewDecoder(buffReader)
	err = decoder.Decode(&currentEvent)
	if err != nil {
		log.Fatal(err)
	}

	var oldHash string = ""

	// Load old hash
	oldHashData, err := ioutil.ReadFile("old-hash.txt")
	if err == nil {
		oldHash = string(oldHashData)
	}

	// Overwrite it with current hash
	hashFile, err := os.Create("old-hash.txt")
	if err != nil {
		log.Fatal(err)
	}

	hashFile.WriteString(currentEvent.Hash)
	hashFile.Close()

	// Compare to new hash
	if oldHash == currentEvent.Hash {
		// Don't enerate site again
		os.Exit(1)
	}
}
