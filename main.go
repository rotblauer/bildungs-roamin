package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/rotblauer/bildRoam/bildRoam"
	"io"
	"os/exec"
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
	var targetDir string
	var myName string
	var uniqify bool

	// go run main.go -dir ~/Desktop/locs/ -out ~/Desktop/locs/whois.csv -name jl

	flag.StringVar(&output, "out", "iwazhere.csv", "specify the output .csv file")
	flag.StringVar(&targetDir, "dir", "/Users/ia/Pictures/Photos Library.photoslibrary/Masters", "directory with photos")
	flag.StringVar(&myName, "name", "ia", "your tag name")
	flag.BoolVar(&uniqify, "uniq", true, "make the output csv pipe through sort -u after")

	flag.Parse()

	file, err := os.Create(output)
	if err != nil {
		fmt.Println("Can't create file.", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)

	filepath.Walk(targetDir, func(path string, info os.FileInfo, err error) error {
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

	if uniqify {
		uniq := exec.Command("sort", "-u", output)
		upipe, e := uniq.StdoutPipe()
		if e != nil {
			fmt.Println(upipe)
		}

		// open the tmp out file for writing
		tmpout, err := os.Create("./sortedtmp.csv")
		if err != nil {
			panic(err)
		}
		defer tmpout.Close()

		writer := bufio.NewWriter(tmpout)

		es := uniq.Start()
		if es != nil {
			fmt.Println(es)
		}

		go io.Copy(writer, upipe)
		uniq.Wait()

		// now once our sorting is complete, we'll make tmpout the real output
		os.Rename("./sortedtmp.csv", output)
	}
}
