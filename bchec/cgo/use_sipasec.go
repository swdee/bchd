package main

import(
	"github.com/swdee/bchd/bchec"
	"github.com/swdee/bchd/bchec/cgo/sipasec"
	"fmt"
)

func EC_Verify(k, s, h []byte) bool {
	return sipasec.EC_Verify(k, s, h) == 1
}

func init() {
	// dirty way of indicating bchd is compiled with sipa lib, could be made to
	// use bchdLog by patching bchd.go and doing a check in bchdMain() after the
	// logger has been started.
	fmt.Println("Using libsecp256k1.a by sipa for EC_Verify")

	bchec.EC_Verify = EC_Verify
}