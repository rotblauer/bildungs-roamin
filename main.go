package main

import (
	"fmt"
	"github.com/rotblauer/bildRoam/bildRoam"
	// "log"
	"os"
	"path/filepath"
)

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
		}

		// save to... csv?
		// fmt.Println(path, lat, lng, ti)
		// counter++
		return nil //nil
	})

}
