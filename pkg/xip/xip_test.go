package xip

import (
	"fmt"
	"testing"
)

func TestGetLocalIp(t *testing.T) {
	fmt.Println(GetLocalIp())
}
