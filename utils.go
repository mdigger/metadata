package metadata

import (
	"crypto/rand"
	"fmt"
	"io"
	"regexp"
	"time"
)

// NewUUID returns the canonical string representation of a UUID:
//  xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx.
func NewUUID() string {
	var uuid [16]byte
	if _, err := io.ReadFull(rand.Reader, uuid[:]); err != nil {
		panic(err)
	}
	uuid[6] = (uuid[6] & 0x0f) | 0x40 // set version byte
	uuid[8] = (uuid[8] & 0x3f) | 0x80 // set high order byte 0b10{8,9,a,b}
	return fmt.Sprintf("%x-%x-%x-%x-%x",
		uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:])
}

// ParseDate return parsed date.
func ParseDate(date string) time.Time {
	if date != "" {
		for _, layout := range []string{"2006-01-02", "2006-01", "2006", time.RFC3339} {
			if parsedDate, err := time.Parse(layout, date); err == nil {
				return parsedDate
			}
		}
	}
	return time.Time{}
}

var reVersion = regexp.MustCompile(`^\d{1,3}(\.\d{1,3}){2,}$`)

func CheckVersion(ver string) bool {
	return reVersion.MatchString(ver)
}
