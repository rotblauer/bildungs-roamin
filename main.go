package main

import (
	"flag"
	"fmt"

	"github.com/rotblauer/bildRoam/bildRoam"
	// "log"
	"encoding/csv"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

var headerLine = []string{"name", "time", "lat", "long"}

func floatToString(input float64) string {
	// to convert a float number to a string
	return strconv.FormatFloat(input, 'f', 6, 64)
}

func main() {
	var output string
	var iphotosDir string
	var myName string

	// go run main.go -dir ~/Desktop/locs/ -out ~/Desktop/locs/whois.csv -name jl

	flag.StringVar(&output, "out", "iwazhere.csv", "specify the output .csv file")
	flag.StringVar(&iphotosDir, "dir", "/Users/ia/Pictures/iphoto-export", "directory with photos")
	flag.StringVar(&myName, "name", "ia", "your tag name")

	flag.Parse()

	file, err := os.Create(output)
	if err != nil {
		fmt.Println("Can't create file.", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)

	filepath.Walk(iphotosDir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		lat, lng, ti, err := bildRoam.GetLatLngTime(path)
		if err != nil {
			// log.Println("ERROR: ", err, path)
		} else {

			fmt.Println("SUCCE: ", path, lat, lng, ti)
			err := writer.Write([]string{myName, ti.Format(time.UnixDate), floatToString(lat), floatToString(lng)})
			if err != nil {
				fmt.Println("error writng csv line", err)
			}
		}

		// save to... csv?
		// fmt.Println(path, lat, lng, ti)
		// counter++
		return nil //nil
	})

	defer writer.Flush()

}
