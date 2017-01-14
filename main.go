package main

import (
	"fmt"
	"github.com/rotblauer/bildRoam/bildRoam"
	// "log"
	"encoding/csv"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

var myName = "ia"
var headerLine = []string{"name", "time", "lat", "long"}

func floatToString(input float64) string {
	// to convert a float number to a string
	return strconv.FormatFloat(input, 'f', 6, 64)
}

func main() {

	// // fil := "./iphoto-export/IMG_4019.JPG"
	// // fil := "IMG_4019.JPG"
	// // fil := "01b TITLE-2.png"
	// // fil := "/Users/ia/Pictures/iphoto-export/01b TITLE-2.png"
	// // filp, _ := filepath.Abs(fil)
	// lat, lng, ti, err := bildRoam.GetLatLngTime(fil)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(lat, lng, ti)

	file, err := os.Create("iwazhere.csv")
	if err != nil {
		fmt.Println("Can't create file.", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)

	iphotosDir := "/Users/ia/Pictures/iphoto-export"
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
