package igstory

import (
	"strconv"
	"testing"
)

func TestGetAllStories(t *testing.T) {
	trays, err := GetAllStories(config)
	if err != nil {
		t.Error(err)
		return
	}

	for _, tray := range trays {
		printTray(tray)
	}

	if len(trays) > 0 {
		GetUserStories(strconv.Itoa(trays[0].Id), config)
	}
}
