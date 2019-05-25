package main

import (
	// "fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	res "github.com/hyperledger/fabric/protos/peer"
	"os"
	// "github.com/hyperledger/fabric/core/chaincode/lib/cid"
	// "strconv"
)

//=======================================================================================
//   Logger : Helps in logging information
//=======================================================================================
var logger = shim.NewLogger("empDir")

//=======================================================================================
//   main() : Entry point for chaincode
//=======================================================================================

func main() {
	logger.SetLevel(shim.LogDebug)
	logLevel, _ := shim.LogLevel(os.Getenv("SHIM_LOGGING_LEVEL"))
	shim.SetLoggingLevel(logLevel)
	err := shim.Start(new(ChainCode))
	if err != nil {
		logger.Error("Error starting patentchain chaincode: %s", err)
	}
}

//=================================================================================================================================
//   Init : Trigger at the time of chaincode instantiation/upgradation
//=================================================================================================================================
func (t *ChainCode) Init(stub shim.ChaincodeStubInterface) res.Response {
	     return shim.Success(nil)
}

//=================================================================================================================================
//   Invoke : Perform transactions
//=================================================================================================================================
func (t *ChainCode) Invoke(stub shim.ChaincodeStubInterface) res.Response {
	function, args := stub.GetFunctionAndParameters()
	if len(args)<1{
		logger.Error("Inappropriate number of arguments")
		shim.Error("Inappropriate number of arguments")
	}

	var result string
	var err error
	
	switch function {
	case "saveIndividualDetails":
		result, err = t.saveIndividualDetails(stub, args)      
	case "submitIndividualDetails":
		result, err = t.submitIndividualDetails(stub, args)  		
	case "allowUniversity":
		result, err = t.allowUniversity(stub, args)  
	case "allowEmployer":
		result, err = t.allowEmployer(stub, args) 		
	case "getIndividualDetails":
		result, err = t.getIndividualDetails(stub, args[0])
	case "getStudentDetails":
		result, err = t.getStudentDetails(stub, args)
	case "getEmployeeDetails":
		result, err = t.getEmployeeDetails(stub, args)				
	default:
		logger.Error("Invalid function")
		return shim.Error("Invalid function")
	}

	if err != nil {
		return shim.Error(err.Error())
    }

	return shim.Success([]byte(result))
}


