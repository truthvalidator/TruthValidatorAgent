package utils

// Package utils provides utility functions for the AI search system.
// This file contains string manipulation utilities.

import (
	"encoding/json"
	"io"
	"strings"
	"unicode"
)

// HasTails checks if a string ends with any of the given suffixes.
//
// Parameters:
//   str: The string to check
//   tails: Variadic list of suffixes to check against
//
// Returns:
//   bool: True if string ends with any suffix, false otherwise
//
// Example:
//   HasTails("file.txt", ".txt", ".csv") // returns true
func HasTails(str string, tails ...string) bool {
	for _, tail := range tails {
		if strings.HasSuffix(str, tail) {
			return true
		}
	}
	return false
}

// CleanStr removes non-graphic and non-whitespace control characters from a string.
//
// Parameters:
//   s: The string to clean
//
// Returns:
//   string: Cleaned string with only graphic and whitespace characters
//
// Example:
//   CleanStr("Hello\x00World") // returns "HelloWorld"
func CleanStr(s string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsGraphic(r) {
			return r
		}
		if r == '\n' || r == '\r' || r == '\t' {
			return r
		}
		return -1
	}, s)
}

// StrMaxLen truncates a string to a maximum length in runes.
//
// Parameters:
//   str: The string to truncate
//   maxLen: Maximum number of runes to keep
//
// Returns:
//   string: Truncated string if needed
//
// Example:
//   StrMaxLen("你好世界", 2) // returns "你好"
func StrMaxLen(str string, maxLen int) string {
	ss := []rune(str)
	if len(ss) <= maxLen {
		return str
	}
	return string(ss[:maxLen])
}

// StrMaxLenSmart truncates a string with intelligent length adjustment based on character type.
// Latin characters get longer max length than CJK characters.
//
// Parameters:
//   str: The string to truncate
//   maxLen: Base maximum length in runes
//   tail: String to append when truncated
//
// Returns:
//   string: Truncated string with tail if needed
//
// Example:
//   StrMaxLenSmart("Hello World", 5, "...") // returns "Hello..."
//   StrMaxLenSmart("你好世界", 2, "...") // returns "你好..."
func StrMaxLenSmart(str string, maxLen int, tail string) string {
	// Adjust max length based on character type
	diffRate := StrByteRuneDiffRate(str)
	if diffRate < 0.2 {  // Mostly Latin
		maxLen = int(4 * float64(maxLen))
	} else if diffRate < 0.4 {  // Mixed
		maxLen = int(2 * float64(maxLen))
	}
	ss := []rune(str)
	if len(ss) <= maxLen {
		return str
	}
	return string(ss[:maxLen]) + tail
}

// StrByteRuneDiffRate calculates the difference rate between byte and rune length.
// Used to determine if a string is mostly Latin (low diff) or CJK (high diff).
//
// Parameters:
//   str: The string to analyze
//
// Returns:
//   float64: Difference rate (0-1)
//
// Example:
//   StrByteRuneDiffRate("Hello") // returns ~0
//   StrByteRuneDiffRate("你好") // returns ~0.5
func StrByteRuneDiffRate(str string) float64 {
	bl := float64(len([]byte(str)))
	rl := float64(len([]rune(str)))
	return (1.0 * (bl - rl)) / bl
}

// StrContains checks if a string exists in a list of strings.
//
// Parameters:
//   c: The string to search for
//   ss: Variadic list of strings to search in
//
// Returns:
//   bool: True if string is found, false otherwise
//
// Example:
//   StrContains("a", "a", "b", "c") // returns true
func StrContains(c string, ss ...string) bool {
	if ss == nil {
		return false
	}

	for _, s := range ss {
		if s == c {
			return true
		}
	}
	return false
}

// IndexOfRunes finds the index of a rune in a rune slice.
//
// Parameters:
//   rs: Slice of runes to search
//   r: Rune to find
//
// Returns:
//   int: Index of rune, or -1 if not found
//
// Example:
//   IndexOfRunes([]rune("abc"), 'b') // returns 1
func IndexOfRunes(rs []rune, r rune) int {
	for idx, r1 := range rs {
		if r == r1 {
			return idx
		}
	}
	return -1
}

// HasSensitiveWords checks if a string contains any sensitive words.
// Note: sensitiveWords list is currently empty.
//
// Parameters:
//   str: The string to check
//
// Returns:
//   bool: True if sensitive word found, false otherwise
//
// Example:
//   HasSensitiveWords("password") // returns false (unless in sensitiveWords)
func HasSensitiveWords(str string) bool {
	for _, s := range sensitiveWords {
		if strings.Contains(str, s) {
			return true
		}
	}
	return false
}

// sensitiveWords contains words to filter out (currently empty)
var sensitiveWords = []string{}

// writeSplitByRune is a helper function for WriteSplitByRune that handles splitting by a single rune.
func writeSplitByRune(w io.Writer, buf *strings.Builder, delta string, subStr rune) (newDelta string, needContinue bool) {
	rs := []rune(delta)
	idx := IndexOfRunes(rs, subStr)
	if idx < 0 {
		newDelta = delta
		return
	}
	toPrint := buf.String() + string(rs[:idx+1])
	w.Write([]byte(toPrint))
	buf.Reset()
	needContinue = true
	newDelta = string(rs[idx+1:])
	return
}

// WriteSplitByRune writes content to a writer, splitting and flushing when any of the given runes are encountered.
//
// Parameters:
//   w: io.Writer to write to
//   buf: strings.Builder buffer holding content
//   delta: New content to process
//   subStrs: Runes that trigger a flush
//
// Returns:
//   bool: True if a flush occurred, false otherwise
//
// Example:
//   var buf strings.Builder
//   WriteSplitByRune(os.Stdout, &buf, "a\nb", '\n') // prints "a\n"
func WriteSplitByRune(w io.Writer, buf *strings.Builder, delta string, subStrs ...rune) (needContinue bool) {
	for _, subString := range subStrs {
		if delta, needContinue = writeSplitByRune(w, buf, delta, subString); needContinue {
			buf.WriteString(delta)
			return
		}
	}
	buf.WriteString(delta)
	return
}

// Json converts any value to its JSON string representation.
// Silently ignores marshaling errors (returns empty string on error).
//
// Parameters:
//   v: The value to marshal
//
// Returns:
//   string: JSON string representation
//
// Example:
//   Json(map[string]int{"a":1}) // returns `{"a":1}`
func Json(v any) string {
	bs, _ := json.Marshal(v)
	return string(bs)
}
