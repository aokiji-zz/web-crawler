package main

import (
	"fmt"
	"runtime"
	"web-crawler/db"
	"web-crawler/src/links"
)

func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Alloc = %v MiB", m.Alloc/1024/1024)
	fmt.Printf("\tTotalAlloc = %v MiB", m.TotalAlloc/1024/1024)
	fmt.Printf("\tSys = %v MiB", m.Sys/1024/1024)
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func main() {
	PrintMemUsage()
	links.VisitedLink("https://aprendagolang.com.br")
	db.FindMany()
	PrintMemUsage()
}
