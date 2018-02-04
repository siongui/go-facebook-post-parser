package igstory

// Get all stories of users that a user follows

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

const UrlAllStories = `https://i.instagram.com/api/v1/feed/reels_tray/`

type RawReelsTray struct {
	Trays []Tray `json:"tray"`
}

type Tray struct {
	Id              int        `json:"id"`
	LatestReelMedia int        `json:"latest_reel_media"`
	User            TrayUser   `json:"user"`
	Items           []TrayItem `json:"items"`
}

type TrayUser struct {
	Username string `json:"username"`
}

type TrayItem struct {
	TakenAt         int                    `json:"taken_at"`
	DeviceTimestamp int                    `json:"device_timestamp"`
	VideoVersions   []TrayItemVideoVersion `json:"video_versions"`
	HasAudio        bool                   `json:"has_audio"`
}

type TrayItemVideoVersion struct {
	Url string `json:"url"`
	Id  string `json:"id"`
}

func GetAllStories(cfg map[string]string) (err error) {
	req, err := http.NewRequest("GET", UrlAllStories, nil)
	if err != nil {
		return
	}

	req.AddCookie(&http.Cookie{Name: "ds_user_id", Value: cfg["ds_user_id"]})
	req.AddCookie(&http.Cookie{Name: "sessionid", Value: cfg["sessionid"]})
	req.AddCookie(&http.Cookie{Name: "csrftoken", Value: cfg["csrftoken"]})

	req.Header.Set("User-Agent", cfg["User-Agent"])

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		err = errors.New("resp.StatusCode: " + strconv.Itoa(resp.StatusCode))
		return
	}

	dec := json.NewDecoder(resp.Body)
	t := RawReelsTray{}
	if err = dec.Decode(&t); err != nil {
		return
	}
	for _, tray := range t.Trays {
		printTray(tray)
	}

	return
}

func printTray(tray Tray) {
	fmt.Print(tray.Id)
	fmt.Print(" : ")
	fmt.Println(tray.User.Username)
	for _, item := range tray.Items {
		if len(item.VideoVersions) > 0 {
			fmt.Println(item.VideoVersions[0].Url)
		}
	}
}
