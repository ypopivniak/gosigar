// Copyright (c) 2012 VMware, Inc.

package gosigar

import (
	"bytes"
)

// byteListToString converts the raw byte arrays we get into a string. This is a bit of a process, as byte strings are normally []uint8
func byteListToString(raw []int8) string {
	byteList := make([]byte, len(raw))

	for pos, singleByte := range raw {
		byteList[pos] = byte(singleByte)
		if singleByte == 0 {
			break
		}
	}

	return string(bytes.Trim(byteList, "\x00"))
}

func chop(buf []byte) []byte {
	return buf[0 : len(buf)-1]
}

func convertBytesToString(arr []byte) string {
	n := bytes.IndexByte(arr, 0)
	if n == -1 {
		return string(arr[:])
	}
	return string(arr[:n])
}
