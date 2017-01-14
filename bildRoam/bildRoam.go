package bildRoam

import (
	// "errors"
	"github.com/rwcarlsen/goexif/exif"
	// "log"
	// "net/http"
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
		return 0, 0, time.Now(), err
	}

	x, err := exif.Decode(f) //-> *Exif, err
	if err != nil {
		return 0, 0, time.Now(), err
	}

	f.Close()

	//sometimes this throws an actual erro like seconds out of bounds at 82
	t, err = x.DateTime()
	if err != nil {
		// log.Println("datetime")
		// log.Fatal(err)
		// return 0, 0, t, err
		t = time.Now()
	}

	lat, lng, err = x.LatLong()
	if err != nil {
		// return 0, 0, t, err
	}

	return lat, lng, t, err
}
