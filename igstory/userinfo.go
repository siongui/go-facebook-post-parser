package igstory

// read user information, such as id,  via username
// see ``Instagram API -Get the userId - Stack Overflow``
// https://stackoverflow.com/a/44773079

import (
	"strings"
)

// no need to login or cookie to access this URL
const UrlUserInfo = `https://www.instagram.com/{{USERNAME}}/?__a=1`

func GetUserInfo(username string) string {
	url := strings.Replace(UrlUserInfo, "{{USERNAME}}", username, 1)
	return url
}
