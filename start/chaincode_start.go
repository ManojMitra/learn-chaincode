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
	"encoding/json"
	"errors"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

// Pre Auth form structure
type preAuthForm struct {
	preAuthID     string `json:"preAuthID"`
	preAuthStatus string `json:"preAuthStatus"`

	providerName          string `json:"providerName"`
	providerAddr          string `json:"providerAddr"`
	providerCityZip       string `json:"providerCityZip"`
	providerPhone         string `json:"providerPhone"`
	providerFax           string `json:"providerFax"`
	providerContactPerson string `json:"providerContactPerson"`

	memName string `json:"memName"`
	memID   string `json:"memID"`
	memDOB  string `json:"memDOB"`
	memDOR  string `json:"memDOR"`

	/*srvReq       string `json:"srvReq"`
	srvDOS       string `json:"srvDOS"`
	srvDiagnosis string `json:"srvDiagnosis"`
	srvCptCode   string `json:"srvCptCode"`
	srvIcdCode   string `json:"srvIcdCode"`
	srvFacility  string `json:"srvFacility"`
	srvPhone     string `json:"srvPhone"`
	srvAddr      string `json:"srvAddr"`
	srvCityZip   string `json:"srvCityZip"`

	payerLOS         string `json:"payerLOS"`
	payerProvTIN     string `json:"payerProvTIN"`
	payerDOS         string `json:"payerDOS"`
	payerBillTIN     string `json:"payerBillTIN"`
	payerAmtAuth     string `json:"payerAmtAuth"`
	payerDiag        string `json:"payerDiag"`
	payerAllowedProc string `json:"payerAllowedProc"`
	payerComment     string `json:"payerComment"`
	payerAddDocReqd  string `json:"payerAddDocReqd"`*/
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
	fmt.Printf("Init called by MM")
	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting 1")
	}
	return t.writeDummyRec(stub)
	//return nil, nil
}

func (t *SimpleChaincode) writeDummyRec(stub shim.ChaincodeStubInterface) ([]byte, error) {
	preAuth1 := preAuthForm{"PA001", "Submitted", "John Smith", "XYZ Capitol avenue NY", "22322", "112-223-22222", "112-223-22223", "Susan Smith", "Johnson", "MNM11231122", "02-22-1986", "04-02-2016"}
	theJSON1, _ := json.Marshal(preAuth1)
	err := stub.PutState("PA001", theJSON1)

	fmt.Println("Wrote details with: " + preAuth1.preAuthID + " - " + preAuth1.preAuthStatus + " - " + preAuth1.providerName + " - " + preAuth1.providerAddr + " - " + preAuth1.providerCityZip + " - " + preAuth1.providerPhone + " - " + preAuth1.providerFax + " - " + preAuth1.providerContactPerson + " - " + preAuth1.memName + " - " + preAuth1.memID + " - " + preAuth1.memDOB + " - " + preAuth1.memDOR)

	if err != nil {
		return nil, err
	}

	/*preAuth1 = preAuthForm{"PA002", "Submitted", "Steven Foss", "ABC Capitol avenue NY", "22321", "112-223-33333", "112-223-33334", "Susan Smith", "Jim", "MNM11231124", "02-22-1986", "04-02-2016"}
	theJSON2, _ := json.Marshal(preAuth1)
	err = stub.PutState("PA002", theJSON2)
	if err != nil {
		return nil, err
	}

	fmt.Println("Wrote details with: " + preAuth1.preAuthID + " - " + preAuth1.preAuthStatus + " - " + preAuth1.providerName + " - " + preAuth1.providerAddr + " - " + preAuth1.providerCityZip + " - " + preAuth1.providerPhone + " - " + preAuth1.providerFax + " - " + preAuth1.providerContactPerson + " - " + preAuth1.memName + " - " + preAuth1.memID + " - " + preAuth1.memDOB + " - " + preAuth1.memDOR)

	preAuth1 = preAuthForm{"PA003", "Submitted", "Steven Foss", "ABC Capitol avenue NY", "22323", "112-223-33333", "112-223-33334", "Susan Smith", "Robert", "MNM11231125", "02-22-1986", "04-02-2016"}
	theJSON3, _ := json.Marshal(preAuth1)
	err = stub.PutState("PA003", theJSON3)
	if err != nil {
		return nil, err
	}
	fmt.Println("Wrote details with: " + preAuth1.preAuthID + " - " + preAuth1.preAuthStatus + " - " + preAuth1.providerName + " - " + preAuth1.providerAddr + " - " + preAuth1.providerCityZip + " - " + preAuth1.providerPhone + " - " + preAuth1.providerFax + " - " + preAuth1.providerContactPerson + " - " + preAuth1.memName + " - " + preAuth1.memID + " - " + preAuth1.memDOB + " - " + preAuth1.memDOR)

	preAuth1 = preAuthForm{"PA004", "Submitted", "Tad Harison", "ABC Capitol avenue NY", "22323", "112-223-33344", "112-223-33345", "Robert Smith", "Kim", "MNM11231126", "02-22-1986", "04-02-2016"}
	theJSON4, _ := json.Marshal(preAuth1)
	err = stub.PutState("PA004", theJSON4)
	if err != nil {
		return nil, err
	}
	fmt.Println("Wrote details with: " + preAuth1.preAuthID + " - " + preAuth1.preAuthStatus + " - " + preAuth1.providerName + " - " + preAuth1.providerAddr + " - " + preAuth1.providerCityZip + " - " + preAuth1.providerPhone + " - " + preAuth1.providerFax + " - " + preAuth1.providerContactPerson + " - " + preAuth1.memName + " - " + preAuth1.memID + " - " + preAuth1.memDOB + " - " + preAuth1.memDOR)

	preAuth1 = preAuthForm{"PA005", "Submitted", "Albert", "ABC Capitol avenue NY", "22323", "112-223-33355", "112-223-33356", "Robert Smith", "Rose", "MNM11231127", "02-22-1986", "04-02-2016"}
	theJSON5, _ := json.Marshal(preAuth1)
	err = stub.PutState("PA005", theJSON5)
	if err != nil {
		return nil, err
	}
	fmt.Println("Wrote details with: " + preAuth1.preAuthID + " - " + preAuth1.preAuthStatus + " - " + preAuth1.providerName + " - " + preAuth1.providerAddr + " - " + preAuth1.providerCityZip + " - " + preAuth1.providerPhone + " - " + preAuth1.providerFax + " - " + preAuth1.providerContactPerson + " - " + preAuth1.memName + " - " + preAuth1.memID + " - " + preAuth1.memDOB + " - " + preAuth1.memDOR)
	*/
	return nil, nil
}

// Invoke is our entry point to invoke a chaincode function
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("invoke is running MM" + function)
	fmt.Println(args)

	// Handle different functions
	if function == "init" {
		return t.Init(stub, "init", args)
	} else if function == "write" {
		return t.write(stub, args)
	} else if function == "read" {
		return t.read(stub, args)
	}

	return nil, errors.New("Received unknown function invocation: " + function)
}

func (t *SimpleChaincode) write(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	fmt.Printf("write called by MM")
	if len(args) != 4 {
		return nil, errors.New("Incorrect number of arguments. Expecting 4")
	}

	/*key := "PT" + strconv.Itoa(rand.Intn(10000000))
	fmt.Println("Key is: " + key)

	pt := Patient{args[0], args[1], args[2], args[3]}

	fmt.Println("- start patient details with " + pt.Name + " - " + pt.MemberID + " - " + " - " + pt.DOB + " - " + pt.DOR)
	jsonAsBytes, _ := json.Marshal(pt)
	fmt.Printf("%+v\n", string(jsonAsBytes))

	err := stub.PutState(key, jsonAsBytes)

	if err != nil {
		return nil, err
	}*/
	return nil, nil
}

func (t *SimpleChaincode) read(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	fmt.Println("read called by MM")
	var key, jsonResp string
	var err error

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting name of the var to query")
	}

	key = args[0]
	fmt.Println("value received in read is : " + key)

	valAsbytes, err := stub.GetState(key)
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + key + "\"}"
		return nil, errors.New(jsonResp)
	}

	if len(valAsbytes) > 0 {
		var p preAuthForm
		err1 := json.Unmarshal(valAsbytes, &p)
		if err1 != nil {
			jsonResp = "{\"Error\":\"Failed to get object }"
			fmt.Printf("Error starting Simple chaincode: %s", err1)
			return nil, errors.New(jsonResp)
		}

		fmt.Println(string(valAsbytes))

		fmt.Println("Reading details of: " + p.preAuthID + " - " + p.preAuthStatus + " - " + p.providerName + " - " + p.providerAddr + " - " + p.providerCityZip + " - " + p.providerPhone + " - " + p.providerFax + " - " + p.providerContactPerson + " - " + p.memName + " - " + p.memID + " - " + p.memDOB + " - " + p.memDOR)

	} else {
		fmt.Printf("Response length is 0 : " + string(valAsbytes))
	}

	return valAsbytes, nil
}

// Query is our entry point for queries
func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("query is running MM " + function)

	// Handle different functions
	if function == "read" { //read a variable
		fmt.Println("hi there " + function)
		return t.read(stub, args)
	}

	return nil, errors.New("Received unknown function query: " + function)
}
