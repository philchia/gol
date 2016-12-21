package internal

import "unsafe"

// JoinStrings join slice strings with ""
func JoinStrings(strs ...string) string {
	ln := 0
	for i := 0; i < len(strs); i++ {
		ln += len(strs[i])
	}
	bts := make([]byte, ln)
	ln = 0
	for _, str := range strs {
		ln += copy(bts[ln:], str)
	}

	return Bytes2str(bts)
}

// Str2bytes trans string to bytes
func Str2bytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

// Bytes2str trans bytes to string
func Bytes2str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
