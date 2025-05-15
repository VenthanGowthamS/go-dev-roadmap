package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
)

func main() {
	// Start pprof server
	go func() {
		err := http.ListenAndServe("localhost:6060", nil)
		if err != nil {
			log.Fatal("Failed to start pprof server:", err)
		}
	}()

	// Simulate some work (optional)
	file, err := os.Create("testfile.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	//defer file.Close()
	fmt.Fprintln(file, "Hello, file!")

	// Keep the program running
	select {}
}

/* PS C:\Users\Venthan Gowtham> go tool pprof http://localhost:6060/debug/pprof/heap
Fetching profile over HTTP from http://localhost:6060/debug/pprof/heap
Saved profile in C:\Users\Venthan Gowtham\pprof\pprof.exception-handling.exe.alloc_objects.alloc_space.inuse_objects.inuse_space.002.pb.gz
File: exception-handling.exe
Build ID: C:\Users\VENTHA~1\AppData\Local\Temp\go-build3565374245\b001\exe\exception-handling.exe2025-05-15 09:08:35.9203214 +0800 +08
Type: inuse_space
Time: 2025-05-15 09:09:04 +08
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) top
Showing nodes accounting for 513kB, 100% of 513kB total
      flat  flat%   sum%        cum   cum%
     513kB   100%   100%      513kB   100%  runtime.allocm
         0     0%   100%      513kB   100%  runtime.mstart
         0     0%   100%      513kB   100%  runtime.mstart0
         0     0%   100%      513kB   100%  runtime.mstart1
         0     0%   100%      513kB   100%  runtime.newm
         0     0%   100%      513kB   100%  runtime.resetspinning
         0     0%   100%      513kB   100%  runtime.schedule
         0     0%   100%      513kB   100%  runtime.startm
         0     0%   100%      513kB   100%  runtime.wakep
(pprof)
*/
