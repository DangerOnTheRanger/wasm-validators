package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/wasmerio/wasmer-go/wasmer"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s [ip]\n", os.Args[0])
		os.Exit(1)
	}

	ip := os.Args[1]

	wasmBytes, err := ioutil.ReadFile("ip_validator.wasm")
	if err != nil {
		log.Fatal(err.Error())
	}

	engine := wasmer.NewEngine()
	store := wasmer.NewStore(engine)

	module, err := wasmer.NewModule(store, wasmBytes)
	if err != nil {
		log.Fatal(err.Error())
	}

	importObject := wasmer.NewImportObject()
	instance, err := wasmer.NewInstance(module, importObject)

	if err != nil {
		log.Fatal(err.Error())
	}

	// TODO(DangerOnTheRanger): deallocate this properly
	allocate, err := instance.Exports.GetFunction("allocate")
	if err != nil {
		log.Fatal(err.Error())
	}

	allocateResult, err := allocate(len(ip))
	if err != nil {
		log.Fatal(err.Error())
	}
	allocatePtr := allocateResult.(int32)

	memory, err := instance.Exports.GetMemory("memory")
	if err != nil {
		log.Fatal(err.Error())
	}
	memSlice := memory.Data()[allocatePtr:]
	ipLen := len(ip)
	for i := 0; i < ipLen; i++ {
		memSlice[i] = ip[i]
	}
	memSlice[ipLen] = 0

	validateIP, err := instance.Exports.GetFunction("validate_ip")
	if err != nil {
		log.Fatal(err.Error())
	}

	rawResult, err := validateIP(allocatePtr)
	if err != nil {
		log.Fatal(err.Error())
	}
	isValidIP := rawResult.(int32) != 0

	if isValidIP {
		fmt.Printf("%s is a valid IPv6 or IPv4 address\n", ip)
	} else {
		fmt.Printf("%s is not a valid IPv6 or IPv4 address\n", ip)
	}
}
