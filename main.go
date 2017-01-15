package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/rotblauer/bildRoam/bildRoam"
	"io"
	"os/exec"
	"os/user"
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

//http://stackoverflow.com/questions/10510691/how-to-check-whether-a-file-or-directory-denoted-by-a-path-exists-in-golang
// exists returns whether the given file or directory exists or not
func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

func main() {
	var output string
	var targetDir string
	var myName string
	var uniqify bool

	//get user's home dir so default to iphotos lib
	usr, err := user.Current()
	if err != nil {
		fmt.Println("but whom", err)
	}
	hd := usr.HomeDir
	defaultiPhotosLibPath := "Pictures/Photos Library.photoslibrary/Masters"
	defaultiPhotosLibPath, e := filepath.Abs(filepath.Join(hd, defaultiPhotosLibPath))
	if e != nil {
		fmt.Println("had a hard time setting default iphotos lib path")
	}

	//set default tag name
	uname := usr.Username
	if uname == "" {
		uname = usr.Name
	}
	if uname == "" {
		uname = filepath.Base(hd)
	}

	flag.StringVar(&output, "out", "iwazhere.csv", "specify the output .csv file")
	flag.StringVar(&targetDir, "dir", defaultiPhotosLibPath, "directory with photos")
	flag.StringVar(&myName, "name", uname, "your tag name")
	flag.BoolVar(&uniqify, "uniq", true, "make the output csv pipe through sort -u after")

	flag.Parse()

	if exists, _ := pathExists(targetDir); !exists {
		fmt.Println("You said the photos would be here:", targetDir)
		fmt.Println("And there just ain't no such place.")
		return
	}

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
