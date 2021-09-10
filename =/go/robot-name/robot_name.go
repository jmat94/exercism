package robotname

import (
	"fmt"
	"math/rand"
	"time"
)

type Robot string

var robotName = map[Robot]bool{}

func (r *Robot) Reset() {
	*r = ""
}

func getRandomNumber() string {
	return fmt.Sprintf("%03d", rand.Intn(1000))
}

func generateRandomLetter() string {
	return string([]byte{byte(90 - rand.Intn(26))})
}

func (r *Robot) Name() (string, error) {

	rand.Seed(time.Now().UTC().UnixNano())

	if *r == "" {
		*r = Robot(generateRandomLetter()) + Robot(generateRandomLetter()) + Robot(getRandomNumber())

		for robotName[*r] {
			*r = Robot(generateRandomLetter()) + Robot(generateRandomLetter()) + Robot(getRandomNumber())
		}

		robotName[*r] = true
	}

	return string(*r), nil
}
