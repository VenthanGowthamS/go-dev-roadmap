package main

import (
	"errors"
	"fmt"
)

type UserError struct {
	id  int
	msg string
}

func (r *UserError) Error() string {
	return fmt.Sprintf("UserError: ID=%d, message:%s", r.id, r.msg)
}

func main() {

	var err error = &UserError{id: 2, msg: "User Not found"}

	var ue *UserError
	//err contains a value that implements the error interface — like your *UserError

	//You pass a pointer to a pointer (&ue) so Go can assign the value to ue if it matches

	//🔸 “Is the error err (which is an interface) of type *UserError (a custom struct)?”

	// And if yes, then it assigns the actual *UserError object to ue.

	if errors.As(err, &ue) {
		fmt.Println("custom error found \n", ue.id, "\n", ue.msg)
	} else {
		fmt.Println("it is not a custom error")
	}

	var err2 error = errors.New("not custom error")

	if errors.As(err2, &ue) {
		fmt.Println("custom error found ")
	} else {
		fmt.Println("it is not a custom error")
	}
}
