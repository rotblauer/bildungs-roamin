package bildRoam

import (
	"errors"
	"github.com/rwcarlsen/goexif/exif"
	"log"
	"net/http"
	"os"
	"time"
)

var allowedContentTypes = []string{"image/png", "image/jpg", "image/jpeg", "image/tiff", "image/gif"}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

//GetLatLngTime get lat, lng, and time for a file with opening it too.
func GetLatLngTime(path string) (lat, lng float64, t time.Time, err error) {
	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		// log.Println("os opening")
		// log.Fatal(err)
		return 0, 0, time.Now(), err
	}

	// log.Println("Opening ", path)

	buff := make([]byte, 512) // docs tell that it take only first 512 bytes into consideration
	if _, err = f.Read(buff); err != nil {
		// log.Println(err) // do something with that error
		return 0, 0, time.Now(), err
	}
	contentType := http.DetectContentType(buff)
	if !stringInSlice(contentType, allowedContentTypes) {
		// log.Println("Unallowed content type: ", contentType)
		err = errors.New("Unallowed content type.")
		return 0, 0, time.Now(), err
	}
	// log.Println("Content type: ", contentType)

	x, err := exif.Decode(f) //-> *Exif, err
	if err != nil {
		// log.Println("exif decoding")
		// log.Println(err)
		return 0, 0, time.Now(), err
	}

	f.Close()

	t, err = x.DateTime()
	if err != nil {
		log.Println("datetime")
		log.Fatal(err)
		return 0, 0, time.Now(), err
	}

	lat, lng, err = x.LatLong()
	if err != nil {
		log.Println("latlng")
		log.Fatal(err)
		return 0, 0, t, err
	}

	return lat, lng, t, err
}
