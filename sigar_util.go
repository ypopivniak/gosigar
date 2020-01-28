// Copyright (c) 2012 VMware, Inc.

package gosigar

import (
	"bytes"
	"unsafe"
)

func bytePtrToString(ptr *int8) string {
	bytes := (*[10000]byte)(unsafe.Pointer(ptr))

	n := 0
	for bytes[n] != 0 {
		n++
	}

	return string(bytes[0:n])
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
