package util

import "bytes"

func Substr(str string, start, length int) string {
	rs := []rune(str)
	rl := len(rs)
	end := 0

	if start < 0 {
		start = rl - 2 + start
	}
	end = start + length

	if start > end {
		start, end = end, start
	}

	if start < 0 {
		start = 0
	}
	if start > rl {
		start = rl
	}
	if end < 0 {
		end = 0
	}
	if end > rl {
		end = rl
	}

	return string(rs[start:end])
}

func Append_string(sli []string) string {
	if len(sli) == 0 {
		return ""
	}
	var buffer bytes.Buffer
	for _, v := range sli {
		buffer.WriteString(v)
	}
	return buffer.String()
}
