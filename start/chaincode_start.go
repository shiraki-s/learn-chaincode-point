/*
Copyright IBM Corp 2016 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

		 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"errors"
	"fmt"
	"strconv"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

// ============================================================================================================================
// Main
// ============================================================================================================================
func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}

// Init resets all the things
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	
	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting 1")
	}

    return nil, nil

}

// Invoke is our entry point to invoke a chaincode function
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("invoke is running " + function)

	// Handle different functions

	if function == "set" {													//initialize the chaincode state, used as reset
		return t.set(stub, args)
	}

	fmt.Println("invoke did not find func: " + function)					//error

	return nil, errors.New("Received unknown function invocation: " + function)
}

// Query is our entry point for queries
func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("query is running " + function)

	// Handle different functions

	if function == "get" {											//read a variable
		return t.get(stub, args)

	}
	fmt.Println("query did not find func: " + function)						//error

	return nil, errors.New("Received unknown function query: " + function)
}

func (t *SimpleChaincode) set(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
  var key string
  var value int
  var err error
  fmt.Println("running set()")

  if len(args) != 2 {
    return nil, errors.New("Incorrect number of arguments. Expecting 2 name of the key and value to set")
  }

  key = args[0]
  value, err = strconv.Atoi(args[1])
  
  if err != nil {
		return nil, errors.New("Expecting integer value for asset holding")
  }
  
  err = stub.PutState(key, []byte(strconv.Itoa(value)))
  
  if err != nil {
    return nil, err
  }
  return nil, nil
}

func (t *SimpleChaincode) get(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

  var key, jsonResp string
  var err error

  if len(args) != 1 {
    return nil, errors.New("Incorrenct number of arguments. Expecting name of the key to query")
  }

  key = args[0]
  valAsbytes, err := stub.GetState(key)
  if err != nil {
    jsonResp = "{\"Error\":\"Failed to get state for " + key + "\"}"
    return nil, errors.New(jsonResp)
  }

  return valAsbytes, nil
}