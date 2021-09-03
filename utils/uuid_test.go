package utils

import (
	"fmt"
	"testing"
)

func TestGenRandomXUIDByClientID(t *testing.T) {
	res, err := GenRandomXUIDByClientID(20)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)

}
