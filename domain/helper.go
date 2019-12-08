package domain

// NeuroNet
// Help functions
// Copyright © 2019 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	"unsafe"
)

func uint64ToBytes(i uint64) []byte {
	x := (*[8]byte)(unsafe.Pointer(&i))
	out := make([]byte, 0, 8)
	out = append(out, x[:]...)
	return out
}

func bytesToUint64(b []byte) uint64 {
	var x [8]byte
	copy(x[:], b[:])
	return *(*uint64)(unsafe.Pointer(&x))
}

func uint32ToBytes(i uint32) []byte {
	x := (*[4]byte)(unsafe.Pointer(&i))
	out := make([]byte, 0, 4)
	out = append(out, x[:]...)
	return out
}

func bytesToUint32(b []byte) uint32 {
	var x [4]byte
	copy(x[:], b[:])
	return *(*uint32)(unsafe.Pointer(&x))
}
