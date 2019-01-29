package main

import "runtime"
import "runtime/pprof"
import "os"
import "time"

func main() {
	go leakyFunction()
	time.Sleep(2000 * time.Millisecond)
	f, _ := os.Create("./profile.pb.gz")
	defer f.Close()
	runtime.GC()
	pprof.WriteHeapProfile(f)
}

func leakyFunction() {
	s := make([]string, 3)
	for i := 0; i < 10000000; i++ {
		s = append(s, "magical pprof time, devops bh")
	}
}
