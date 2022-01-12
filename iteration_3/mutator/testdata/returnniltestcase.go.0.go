package testdata

import (
	"errors"
	"fmt"
)

func main() {
	err := func() error {
		thisErrWontBeRemoved := errors.New("test")
		fmt.Printf("%s", thisErrWontBeRemoved.Error())
		return nil
	}()

	fmt.Printf("%s", err.Error())
}
