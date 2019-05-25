package main

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"errors"
	"encoding/json"
	"bytes"
	// "github.com/hyperledger/fabric/core/chaincode/lib/cid"
	// "fmt"
	// "time"
	// "strconv"
	// "strings"
)

//=================================================================================================================================
//   SaveIndividualDetails : Save details of an individual
//=================================================================================================================================
func (t *ChainCode) saveIndividualDetails(stub shim.ChaincodeStubInterface, args []string) (string, error ) {

	if len(args)<1{
		logger.Error("Inappropriate number of arguments")
		shim.Error("Inappropriate number of arguments")
	}

	//Buffer to concatenate strings
	var buf bytes.Buffer

	//individualDetails
	var individualDetails EmpDir

	err := json.Unmarshal([]byte(args[0]), & individualDetails)
	if err != nil {
		logger.Error(err)
		return "", errors.New("Error parsing individual info")
	}

	if individualDetails.Submitted == true {
		return "", errors.New("Error: Details are already in submitted state")
	}

	//Fill Univs
	var universities Univs 
	for _, education := range individualDetails.EducationDetails {
		buf.WriteString(education.University)
		buf.WriteString(education.Enrollment)
		univId := buf.String()
		buf.Reset()
		
		universities.UnivID = univId
		universities.PersonalDetails = individualDetails.PersonalDetails
		universities.EducationDetails = education
		universities.Visible = false

		bytes, err := json.Marshal(universities)

		if err != nil {
			logger.Error(err)
			return "", errors.New("Error converting university record")
		}
	
		err = stub.PutState(univId, bytes)
		if err != nil {
			logger.Error(err)
			return "", errors.New("Error storing university record")
		}
	}

	//Fill Employers
	var employers Employers 
	for _, employment := range individualDetails.EmploymentDetails {
		buf.WriteString(employment.EmployerCode)
		buf.WriteString(employment.EmployeeCode)
		employmentId := buf.String()
		buf.Reset()

		employers.EmploymentID = employmentId
		employers.PersonalDetails = individualDetails.PersonalDetails
		employers.EmploymentDetails = employment
		employers.Visible = false
		
		bytes, err := json.Marshal(employers)

		if err != nil {
			logger.Error(err)
			return "", errors.New("Error converting employment record")
		}
	
		err = stub.PutState(employmentId, bytes)
		if err != nil {
			logger.Error(err)
			return "", errors.New("Error storing employment record")
		}
	}	

	bytes, err := json.Marshal(individualDetails)

	if err != nil {
		logger.Error(err)
		return "", errors.New("Error converting individual record")
	}

	err = stub.PutState(individualDetails.UID, bytes)
	if err != nil {
		logger.Error(err)
		return "", errors.New("Error storing individual record")
	}

	return "true", nil
}


//=================================================================================================================================
//   SubmitIndividualDetails : Save details of an individual
//=================================================================================================================================
func (t *ChainCode) submitIndividualDetails(stub shim.ChaincodeStubInterface, args []string) (string, error ) {

	if len(args)<1{
		logger.Error("Inappropriate number of arguments")
		shim.Error("Inappropriate number of arguments")
	}

	//individualDetails
	var individualDetails EmpDir

	err := json.Unmarshal([]byte(args[0]), & individualDetails)
	if err != nil {
		logger.Error(err)
		return "", errors.New("Error parsing individual info")
	}

	if individualDetails.Submitted == true {
		return "", errors.New("Error: Details are already in submitted state")
	}

	individualDetails.Submitted = true;
	bytes, err := json.Marshal(individualDetails)

	if err != nil {
		logger.Error(err)
		return "", errors.New("Error converting individual record")
	}

	err = stub.PutState(individualDetails.UID, bytes)
	if err != nil {
		logger.Error(err)
		return "", errors.New("Error storing individual record")
	}

	return "true", nil
}


//=================================================================================================================================
//   AllowUniversity : Allow university to see details of an individual
//=================================================================================================================================
func (t *ChainCode) allowUniversity(stub shim.ChaincodeStubInterface, args []string) (string, error ) {

	if len(args)<2{
		logger.Error("Inappropriate number of arguments")
		shim.Error("Inappropriate number of arguments")
	}

	//StudentDetails
	var student Univs
    //Buffer to concatenate strings
	var buf bytes.Buffer

	buf.WriteString(args[0])
	buf.WriteString(args[1])
	univId := buf.String()
	buf.Reset()

	results, err1 := stub.GetState(univId)
	if err1 != nil {
		return "", errors.New(err1.Error())
	}

	if results != nil {
		err2 := json.Unmarshal(results, &student)
		if err2 != nil {
			return "", errors.New("Error parsing student's info")
		}
	}

    //Allow university to see info
	student.Visible = true

	bytes, err3 := json.Marshal(student)

	if err3 != nil {
		logger.Error(err3)
		return "", errors.New("Error converting student's record")
	}

	err4 := stub.PutState(univId, bytes)
	if err4 != nil {
		logger.Error(err4)
		return "", errors.New("Error storing student's record")
	}

	return "true", nil
}


//=================================================================================================================================
//   AllowEmployer : Allow employer to see details of an individual
//=================================================================================================================================
func (t *ChainCode) allowEmployer(stub shim.ChaincodeStubInterface, args []string) (string, error ) {

	if len(args)<2{
		logger.Error("Inappropriate number of arguments")
		shim.Error("Inappropriate number of arguments")
	}

	//StudentDetails
	var employee Employers
    //Buffer to concatenate strings
	var buf bytes.Buffer

	buf.WriteString(args[0])
	buf.WriteString(args[1])
	employmentId := buf.String()
	buf.Reset()

	results, err1 := stub.GetState(employmentId)
	if err1 != nil {
		return "", errors.New(err1.Error())
	}

	if results != nil {
		err2 := json.Unmarshal(results, &employee)
		if err2 != nil {
			return "", errors.New("Error parsing employee's info")
		}
	}

    //Allow university to see info
	employee.Visible = true

	bytes, err3 := json.Marshal(employee)

	if err3 != nil {
		logger.Error(err3)
		return "", errors.New("Error converting employee's record")
	}

	err4 := stub.PutState(employmentId, bytes)
	if err4 != nil {
		logger.Error(err4)
		return "", errors.New("Error storing employee's record")
	}

	return "true", nil
}