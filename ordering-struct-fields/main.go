package main

import (
	"fmt"
	"reflect"
	"runtime"
)

type staff struct {
	Name   string // 16 bytes
	ID     uint16 // 2 bytes
	Salary int8   // 1 byte
	Bonus  int32  // 4 bytes
}

func main() {
	typ := reflect.TypeOf(staff{})
	fmt.Printf("Struct is %d bytes long\n", typ.Size()) // printing the size of the struct

	n := typ.NumField()
	for i := 0; i < n; i++ {
		field := typ.Field(i)
		fmt.Printf("%s at offset %v, size=%d, align=%d\n",
			field.Name, field.Offset, field.Type.Size(),
			field.Type.Align())
	}

	// simulate if the struct is processed 100mio times
	allStaffs := []staff{}
	for i := 0; i < 100000000; i++ {
		allStaffs = append(allStaffs, staff{})
	}

	printMemUsage()
}

func printMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
