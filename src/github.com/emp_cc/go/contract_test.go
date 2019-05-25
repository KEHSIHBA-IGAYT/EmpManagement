/*
SimpleLedger_Test uses the functions from mockstub.go.
Basic fuctions like create ledger entry, update ledger entry can be tested.
Composite key functionality also can be tested
How to try that is demonstrated in this file.

******Important******
Has to be read together with the SimpleLedger.go file

Note: History functionality is not implemented in mockstub.go
Guess we will have to deploy on the network and test for some more time :(

@author Sanjay SB
@version 1.0
*/
package main

import (
	"fmt"
	"testing"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

func TestSimpleLedger(t *testing.T) {
	// Various prints have been placed for testing purpose
	fmt.Println("Entered TestSimpleLedger Function")

	simpleChaincode := new(ChainCode)

	stub := shim.NewMockStub("mockstub", simpleChaincode)
	if stub == nil {
		t.Fatalf("MockStub creation failed")
	}

	var args1, args2 [][]byte // the invoke in mockstub takes 2D byte array as an input
	var response pb.Response
	var i int

	//===============
	//Create Battery
	//===============

	//enter the arguments as expected by the method
	args_array1 := []string{"{\"uid\":\"1234\",\"PersonalDetails\":{\"firstName\":\"abc\",\"lastName\":\"abc\",\"dob\":\"20/07/1995\",\"father\":\"abc\",\"mother\":\"abc\",\"address\":\"abc\",\"pan\":\"abc\",\"passport\":\"abc\"},\"skills\":[{\"SkillName\":\"React\",\"experience\":1,\"level\":2},{\"SkillName\":\"Angular\",\"experience\":1,\"level\":1}],\"educationDetails\":[{\"Course\":\"B.Tech\",\"university\":\"APJAK\",\"univcode\":\"U101\",\"college\":\"KEC\",\"enrollment\":\"101611110\",\"level\":2015,\"grades\":\"75%\",\"verified\":false}],\"employmentDetails\":[{\"employer\":\"Sopra\",\"employercode\":\"E101\",\"employeecode\":\"673413\",\"fromDate\":\"2015\",\"toDate\":\"2019\",\"location\":\"Noida\",\"verified\":false}],\"submitted\":false}"}
         
	//create the two dimensional byte array to pass to MockInvoke
	args1 = append(args1, []byte("saveIndividualDetails")) //function name, pass it as the first byte array

	for i = 0; i < len(args_array1); i++ {
		args1 = append(args1, []byte(args_array1[i]))
	}

	//Call the invoke function
	response = stub.MockInvoke("t123", args1)
  
	//Response handling
	printResponse("+++ Invoke Response +++", response)

	// //===============
	// //Update Battery
	// //===============

	//enter the arguments as expected by the above method
	args_array2 := []string{"U101","101611110"}

	//create the two dimensional byte array to pass to MockInvoke
	args2 = append(args2, []byte("getStudentDetails")) //function name, pass it as the first byte array

	for i = 0; i < len(args_array2); i++ {
		args2 = append(args2, []byte(args_array2[i]))
	}

	//Call the invoke function
	response = stub.MockInvoke("t124", args2)

	//Response handling
	printResponse("+++ Query Response +++", response)

	// /*//================
	// //Create Pacemaker
	// //================

	// //enter the arguments as expected by the method
	// args_array3 := []string{"2", "Pacemaker", "", "Manufacturer", ""}

	// //create the two dimensional byte array to pass to MockInvoke
	// args3 = append(args3, []byte("createItem"))

	// for i = 0; i < len(args_array3); i++ {
	// 	args3 = append(args3, []byte(args_array3[i]))
	// }

	// response = stub.MockInvoke("t125", args3)

	// printResponse("+++Create Pacemaker Response+++", response)

	// //====================
	// //Update Battery Again
	// //====================

	// //enter the arguments as expected by the above method
	// args_array4 := []string{"1", "Battery", "2", "Manufacturer", "QA Completed"}

	// //create the two dimensional byte array to pass to MockInvoke
	// args4 = append(args4, []byte("updateItem"))

	// for i = 0; i < len(args_array4); i++ {
	// 	args4 = append(args4, []byte(args_array4[i]))
	// }

	// response = stub.MockInvoke("t126", args4)

	// printResponse("+++Update Battery Response+++", response)

	// //===========================
	// //Query Battery and PaceMaker
	// //===========================*/

	// var queryargs_array1 [][]byte

	// //create the two dimensional byte array to pass to MockInvoke
	// queryargs_array1 = append(queryargs_array1, []byte("queryConsignment"))
	// queryargs_array1 = append(queryargs_array1, []byte("ContainerID"))

	// response = stub.MockInvoke("t300", queryargs_array1)

	// printResponse("+++Query for Item 2+++", response)

	//==================
	//Query By Location
	//==================

}

func printResponse(heading string, response pb.Response) {

	fmt.Println(heading)

	fmt.Printf("Message: ")
	fmt.Println(response.GetMessage())
	fmt.Printf("Payload: ")
	fmt.Println(string(response.GetPayload()))
	fmt.Printf("Status: ")
	fmt.Println(response.GetStatus())

	fmt.Println("+++++End of Response+++++")
}
