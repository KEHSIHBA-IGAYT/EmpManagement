package main

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"fmt"
	"errors"
	"bytes"
	// "github.com/hyperledger/fabric/core/chaincode/lib/cid"
	"encoding/json"
	// "time"
	// "strconv"
	// "strings"
)

//================================================================================
//   GetIndividualDetails : Get details of an individual
//================================================================================
func (t *ChainCode) getIndividualDetails(stub shim.ChaincodeStubInterface, key string) (string, error) {

	if len(key)<1{
		logger.Error("Individual's UID is not appropriate")
		shim.Error("Individual's UID is not appropriate")
	}

	queryString := "{\"selector\":{\"uid\":\"" + key + "\"}}"

	fmt.Println("queryString")
	fmt.Println(queryString)


	queryResults, err1 := getQueryResultForQueryString(stub, queryString)
	if err1 != nil {
		return "", errors.New(err1.Error())
	}
	return string(queryResults), nil

}

//================================================================================
//   GetStudentDetails : Get details of the student for the university
//================================================================================
func (t *ChainCode) getStudentDetails(stub shim.ChaincodeStubInterface, args []string) (string, error) {

	if len(args)<2{
		logger.Error("Invalid number of arguments")
		shim.Error("Invalid number of arguments")
	}

	var buf bytes.Buffer

	buf.WriteString(args[0])
	buf.WriteString(args[1])
	univId := buf.String()
	buf.Reset()

	fmt.Println(univId)

    results, err1 := stub.GetState(univId)
	if err1 != nil {
		return "", errors.New(err1.Error())
	}

	var univInfo Univs

	if results != nil {
		err2 := json.Unmarshal(results, &univInfo)
		if err2 != nil {
			return "", errors.New("Error parsing student's info")
		}
	}

	if univInfo.Visible == false {
		return "", errors.New("Student has not authorized you to see the confidential information")
	}

	fmt.Println(univInfo)

	return string(results), nil
}

//================================================================================
//   GetEmployeeDetails : Get details of the employee for the university
//================================================================================
func (t *ChainCode) getEmployeeDetails(stub shim.ChaincodeStubInterface, args []string) (string, error) {

	if len(args)<2{
		logger.Error("Invalid number of arguments")
		shim.Error("Invalid number of arguments")
	}

	var buf bytes.Buffer

	buf.WriteString(args[0])
	buf.WriteString(args[1])
	employmentId := buf.String()
	buf.Reset()

	fmt.Println(employmentId)

    results, err1 := stub.GetState(employmentId)
	if err1 != nil {
		return "", errors.New(err1.Error())
	}

	var employeeInfo Employers

	if results != nil {
		err2 := json.Unmarshal(results, &employeeInfo)
		if err2 != nil {
			return "", errors.New("Error parsing employee's info")
		}
	}

	if employeeInfo.Visible == false {
		return "", errors.New("Employee has not authorized you to see the confidential informaton")
	}

	fmt.Println(string(results))
	return string(results), nil
}