package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"howett.net/plist"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
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
    Hash	string	`pist:"hash"`
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

	fmt.Println("<html>")
	fmt.Println("<head>")
	fmt.Println("<title>Pocket Planes Global Event</title>")
	fmt.Println("</head>")
	fmt.Println("<body>")
	fmt.Println("<h1>Pocket Planes Global Event</h1>")
	fmt.Println(`<div> class="event">`)
	fmt.Println("<p>The current event is a ")
	fmt.Print(currentEvent.Name)
	fmt.Print(" in ")
	cityName := getCityName(currentEvent.CityID)
	fmt.Println(cityName)
	fmt.Println("</p>")

	// Data table
	fmt.Println("<table>")

	// Start Time
	fmt.Println("<tr>")
	fmt.Println("<th>Start Time</th>")
	fmt.Print("<td>")
	fmt.Print(readableTime(currentEvent.StartTime))
	fmt.Println("</td>")
	fmt.Println("</tr>")

	// End Time
	fmt.Println("<tr>")
	fmt.Println("<th>End Time</th>")
	fmt.Print("<td>")
	fmt.Print(readableTime(currentEvent.EndTime))
	fmt.Println("</td>")
	fmt.Println("</tr>")

	firstReward := getHumanReadableReward(currentEvent.Reward1)
	fmt.Println(firstReward)

	fmt.Println("</table>")
	fmt.Println("</div>")
	fmt.Println("</body>")
	fmt.Println("</html>")
}

func getCityName(id int) string {
	idString := strconv.Itoa(id)
	cityDataFile, err := os.Open("cityInfo.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer cityDataFile.Close()

	r := csv.NewReader(cityDataFile)

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if record[0] == idString {
			return record[1]
		}
	}
	log.Fatal("Failed to look up city name")
	return "UNKNOWN"
}

func getHumanReadableReward(rewardString string) string {
	re := regexp.MustCompile("(part|b|plane):([0-9]+)")
	humanReadable := ""
	matches := re.FindAllStringSubmatch(rewardString, -1)

	for _, match := range matches {
		rewardType := match[1]
		rewardData := match[2]

		if humanReadable != "" {
			humanReadable += " + "
		}

		if rewardType == "b" {
			// Bux
			humanReadable += rewardData + " bux"
		} else if rewardType == "plane" {
			// Entire plane
			humanReadable += getPlaneName(rewardData)
		} else if rewardType == "part" {
			humanReadable += getPlaneName(rewardData) + " part"
		}
	}
	return humanReadable
}

func getPlaneName(id string) string {
	planeDataFile, err := os.Open("planeInfo.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer planeDataFile.Close()

	r := csv.NewReader(planeDataFile)

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if record[0] == id {
			return record[1]
		}
	}
	log.Fatal("Failed to look up plane name")
	return "UNKNOWN"
}

func readableTime(unixTime float64) string {
	return "time"
}
