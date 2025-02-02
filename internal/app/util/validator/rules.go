package validator

import "regexp"

const (
	usernameRegexString = "^[a-z0-9_.]+$"
)

var usernameRegex = regexp.MustCompile(usernameRegexString)
