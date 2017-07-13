
/*
Employee Information
Block Chain Sample Demo Application
Author: Venkat
*/

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

//Employer - Structure for Employer details
type Employer struct {
	Name   string  `json:"name"`
	Company string `json:"company"`
	Designation string  `json:"designation"`
	Empid string     `json:"empid"`
	Details string     `json:"details"`
}

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

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

	err := stub.PutState("hello_world", []byte(args[0]))
	if err != nil {
		return nil, err
	}

	return nil, nil
}

// Invoke isur entry point to invoke a chaincode function
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("invoke is running " + function)

	// Handle different functions
	if function == "init" {
		return t.Init(stub, "init", args)
	} else if function == "write" {
		return t.write(stub, args)
	} else if function == "addemployer" {
		return t.addEmployer(stub, args)
	}
	fmt.Println("invoke did not find func: " + function)

	return nil, errors.New("Received unknown function invocation: " + function)
}

// Query is our entry point for queries
func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("query is running " + function)

	// Handle different functions
	if function == "read" { //read a variable
		return t.read(stub, args)
	} else if function == "reademployer" {
		return t.readEmployer(stub, args)
	}
	fmt.Println("query did not find func: " + function)

	return nil, errors.New("Received unknown function query: " + function)
}

// write - invoke function to write key/value pair
func (t *SimpleChaincode) write(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var key, value string
	var err error
	fmt.Println("running write()")

	if len(args) != 2 {
		return nil, errors.New("Incorrect number of arguments. Expecting 2. name of the key and value to set")
	}

	key = args[0] //rename for funsies
	value = args[1]
	err = stub.PutState(key, []byte(value)) //write the variable into the chaincode state
	if err != nil {
		return nil, err
	}
	return nil, nil
}

// read - query function to read key/value pair
func (t *SimpleChaincode) read(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var key, jsonResp string
	var err error

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting name of the key to query")
	}

	key = args[0]
	valAsbytes, err := stub.GetState(key)
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + key + "\"}"
		return nil, errors.New(jsonResp)
	}

	return valAsbytes, nil
}

func (t *SimpleChaincode) addEmployer(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	fmt.Println("adding employer information")
	if len(args) != 5 {
		return nil, errors.New("Incorrect Number of arguments.Expecting 5 for adding Employee")
	}
	//amt, err := strconv.ParseFloat(args[1], 64)
	

	employer := Employer{
		Name:   args[0],
		Company: args[1],
		Designation: args[2],
		Empid: args[3],
		Details: args[4],
	}

	bytes, err := json.Marshal(employer)
	if err != nil {
		fmt.Println("Error marshaling employer")
		return nil, errors.New("Error marshaling employer")
	}

	err = stub.PutState(employer.Empid, bytes)
	if err != nil {
		return nil, err
}
return nil, nil
}

func (t *SimpleChaincode) readEmployer(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	fmt.Println("read() is running")

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. expecting 1")
	}

	key := args[0] // name of Entity
	fmt.Println("key is ")
	fmt.Println(key)
	bytes, err := stub.GetState(args[0])
	fmt.Println(bytes)
	if err != nil {
		fmt.Println("Error retrieving " + key)
		return nil, errors.New("Error retrieving " + key)
	}
	/*
	employer := Employer{}
	err = json.Unmarshal(bytes, &employer)
	if err != nil {
		fmt.Println("Error Unmarshaling EmployerBytes")
		return nil, errors.New("Error Unmarshaling EmployerBytes")
	}
	
	bytes, err = json.Marshal(employer)
	if err != nil {
		fmt.Println("Error marshaling Employer")
		return nil, errors.New("Error marshaling Employer")
	}

	fmt.Println(bytes)
	*/
	return bytes, nil
}
