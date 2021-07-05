package main

import (
    "io/ioutil"
    "github.com/bjartek/go-with-the-flow/gwtf"
) 

func main() {

    filename := "../scripts/block.cdc"
   	fb, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	gwtf := gwtf.NewGoWithTheFlowEmulator()

	gwtf.ScriptFromFile(filename, fb).Run()

}
