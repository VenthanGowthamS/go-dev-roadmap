package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", safehandler)
	fmt.Println("🌐 Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func safehandler(w http.ResponseWriter, r *http.Request) {
	defer func() {

		/* 		Runs recover() and assigns its return to err.

		If err is not nil, it goes inside the block. */

		if err := recover(); err != nil {

			http.Error(w, "website crash avoided", http.StatusInternalServerError)
			fmt.Println("Recovered from panic", err)
		}

	}()

	panic("something happeneded")
}

/*
Flow Explanation for defer, panic, and recover:

1. The main function runs — this is the program's entry point.

2. The main function calls safeHandler() (or another function where panic may occur).

3. Inside safeHandler, a deferred anonymous function is registered:

   defer func() {
       if r := recover(); r != nil {
           fmt.Println("Recovered from panic:", r)
       }
   }()

   - This deferred function is registered immediately but will only run when safeHandler returns — either normally or due to a panic.

4. Code runs normally until a panic occurs.

5. When panic happens:

   - Go starts unwinding the call stack, cleaning up from the most recent function call backwards.

   - During unwinding, Go executes any defer blocks in last-in-first-out order.

6. The deferred recovery function runs during this unwinding:

   - recover() captures the panic value.

   - If a panic was present (r != nil), it prints the recovery message.

7. Because recover() handled the panic, the program does NOT crash and continues gracefully.

*/
