package sanitizefilename

import (
	"runtime"
	"strings"
	"unicode"
)

var simulation string

func sanitizeWindows(filename string) string {
	sanitized := []rune{}

	// Cannot end with whitespace or "."
	filename = strings.TrimRightFunc(filename, func(r rune) bool {
		return unicode.IsSpace(r)
	})

	// Efficiency return - trimmed filename is empty
	if filename == "" {
		return ""
	}

	// Illegal filename characters, printable characters
	illRunes := map[rune]int{'<': 1, '>': 1, ':': 1, '"': 1, '/': 1, '\\': 1, '|': 1, '?': 1, '*': 1}

	for _, existingRune := range filename {
		if existingRune <= 31 || existingRune == 127 || !unicode.IsPrint(existingRune) || illRunes[existingRune] == 1 {
			continue
		}

		sanitized = append(sanitized, existingRune)
	}

	// Efficiency return - sanitized filename is empty
	if len(sanitized) == 0 {
		return ""
	}

	// Reserved names
	illNames := map[string]int{
		"CON": 1, "PRN": 1, "AUX": 1, "NUL": 1,
		"COM1": 1, "COM2": 1, "COM3": 1, "COM4": 1, "COM5": 1, "COM6": 1, "COM7": 1, "COM8": 1, "COM9": 1,
		"LPT1": 1, "LPT2": 1, "LPT3": 1, "LPT4": 1, "LPT5": 1, "LPT6": 1, "LPT7": 1, "LPT8": 1, "LPT9": 1,
	}
	sanitizedStr := string(sanitized)
	parts := strings.Split(sanitizedStr, ".") // An attempt to consider reserved names with extensions. While this isn't "illegal", things get weird, so best to avoid.

	if illNames[strings.ToUpper(parts[0])] == 1 {
		parts[0] = ""
		sanitizedStr = strings.Join(parts, ".")
	}

	return sanitizedStr
}

func sanitizeLinixAndUnix(filename string) string {
	sanitized := []rune{}

	// Illegal filename characters, printable characters (technically, only rune == 0 and rune == '/' are illegal,
	// but come on... rune <= 31, rune == 127, and other non-printable runes shouldn't really be in filenames, should they?)
	for _, existingRune := range filename {
		if existingRune <= 31 || existingRune == 127 || existingRune == '/' || !unicode.IsPrint(existingRune) {
			continue
		}

		sanitized = append(sanitized, existingRune)
	}

	// Efficiency return - sanitized filename is empty
	if len(sanitized) == 0 {
		return ""
	}

	sanitizedStr := string(sanitized)

	// Reserved names
	if sanitizedStr == "." || sanitizedStr == ".." {
		return ""
	}

	return sanitizedStr
}

// Sanitizes the given string under the assumption it represents a filename.
// Sanitation is OS-agnostic.
func Sanitize(filename string) string {
	if simulation == "windows" {
		return sanitizeWindows(filename)
	} else if simulation != "" {
		return sanitizeLinixAndUnix(filename)
	} else if runtime.GOOS == "windows" {
		return sanitizeWindows(filename)
	} else {
		return sanitizeLinixAndUnix(filename)
	}
}
