package utils

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"regexp"

	"github.com/parnurzeal/gorequest"
)

// Getopt reads environment variables.
// If not found will return a supplied default value
func Getopt(name, dfault string) string {
	value := os.Getenv(name)
	if value == "" {
		value = dfault
	}
	return value
}

// Assert asserts there was no error, else log.Fatal
func Assert(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// GetSHA256 calculates a file's sha256sum
func GetSHA256(name string) string {

	dat, err := ioutil.ReadFile(name)
	Assert(err)

	h256 := sha256.New()
	_, err = h256.Write(dat)
	Assert(err)

	return fmt.Sprintf("%x", h256.Sum(nil))
}

// RunCommand runs cmd on file
func RunCommand(cmd string, args ...string) string {

	cmdOut, err := exec.Command(cmd, args...).Output()
	if len(cmdOut) == 0 {
		Assert(err)
	}

	return string(cmdOut)
}

func printStatus(resp gorequest.Response, body string, errs []error) {
	fmt.Println(resp.Status)
}

// RemoveDuplicates removes duplicate items from a list
func RemoveDuplicates(elements []string) []string {
	// Use map to record duplicates as we find them.
	encountered := map[string]bool{}
	result := []string{}

	for v := range elements {
		if encountered[elements[v]] == true {
			// Do not add duplicate.
		} else {
			// Record this element as an encountered element.
			encountered[elements[v]] = true
			// Append to result slice.
			result = append(result, elements[v])
		}
	}
	// Return the new slice.
	return result
}

// GetHashType returns the hash type (md5, sha1, sha256, sha512)
func GetHashType(hash string) (string, error) {
	var validMD5 = regexp.MustCompile(`^[a-fA-F\d]{32}$`)
	var validSHA1 = regexp.MustCompile(`^[a-fA-F\d]{40}$`)
	var validSHA256 = regexp.MustCompile(`^[a-fA-F\d]{64}$`)
	var validSHA512 = regexp.MustCompile(`^[a-fA-F\d]{128}$`)

	switch {
	case validMD5.MatchString(hash):
		return "md5", nil
	case validSHA1.MatchString(hash):
		return "sha1", nil
	case validSHA256.MatchString(hash):
		return "sha256", nil
	case validSHA512.MatchString(hash):
		return "sha512", nil
	default:
		return "", errors.New("This is not a valid hash.")
	}
}
