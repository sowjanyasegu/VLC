
/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

/*
 * The sample smart contract for documentation topic:
 * Writing Your First Blockchain Application
 */

package main

/* Imports
 * 4 utility libraries for formatting, handling bytes, reading and writing JSON, and string manipulation
 * 2 specific Hyperledger Fabric specific libraries Smart Contracts
 */
import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/core/chaincode/shim/ext/cid"
	sc "github.com/hyperledger/fabric/protos/peer"
)

// Define the Smart Contract structure for Car Registration
type StatutoryCarRegistration struct {
}

// Define the car structure, with 4 properties.  Structure tags are used by encoding/json library
type Car struct {
    CarId  string `json:"carId"`
	Model  string `json:"model"`
	EngineNumber string `json:"engineNumber"`
	ChassisNumber string `json:"chassisNumber"`
	Colour string `json:"colour"`
	YearOfManufacture string `json:"yearOfManufacture"`
	Owner  string `json:"owner"`
	Dealer string `json:"dealer"`
	InsurancePolicyNumber string `json:"insurancePolicyNumber"`
	RegistrationNumber string `json:"registrationNumber"`
	CurrentState string `json:"currentState"`
}

/*
 * The Init method is called when the Smart Contract "fabcar" is instantiated by the blockchain network
 * Best practice is to have any Ledger initialization in separate function -- see initLedger()
 */
func (s *StatutoryCarRegistration) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

/*
 * The Invoke method is called as a result of an application request to run the Smart Contract "fabcar"
 * The calling application program has also specified the particular smart contract function to be called, with arguments
 */
func (s *StatutoryCarRegistration) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {

	// Retrieve the requested Smart Contract function and arguments
	function, args := APIstub.GetFunctionAndParameters()

	mspId, err := cid.GetMSPID(APIstub)
	fmt.Println("MSPId : ",mspId)
	if err != nil {
		return  shim.Error("Invalid MSP.")
	}
	// Route to the appropriate handler function to interact with the ledger appropriately
	if function == "queryCar" {
		return s.queryCar(APIstub, args)
	} else if function == "initLedger" {
		return s.initLedger(APIstub)
	} else if function == "createCar" && mspId == "ManufacturerMSP"{
		return s.createCar(APIstub, args)
	} else if function == "queryAllCars" {
		return s.queryAllCars(APIstub)
	} else if function == "purchaseCar" && mspId == "DealerMSP"{
		return s.purchaseCar(APIstub, args)
	} else if function == "getCarHistory" {
		return s.getCarHistory(APIstub, args)
	} else if function == "deliverCar" && mspId == "ManufacturerMSP"{
		return s.deliverCar(APIstub, args)
	} else if function == "sellCar" {
		return s.sellCar(APIstub, args)
	} else if function == "insuredCar" && mspId == "InsurerMSP" {
		return s.insuredCar(APIstub, args)
	} else if function == "registeredCar" && mspId == "RTOMSP" {
		return s.registeredCar(APIstub, args)
	} else if function == "deliverCarToCustomer" && mspId == "DealerMSP" {
		return s.deliverCarToCustomer(APIstub, args)
	} else if function == "sellCar" {
		return s.sellCar(APIstub, args)
	} else if function == "queryCarByEngineNumber" {
		return s.queryCarByEngineNumber(APIstub, args)
	} 


	return shim.Error("Invalid Smart Contract function name.") 
}


func (s *StatutoryCarRegistration) initLedger(APIstub shim.ChaincodeStubInterface) sc.Response {
	cars := []Car{
		Car{CarId: "1",EngineNumber: "7456001",ChassisNumber: "1HGCM82633A004352",Model: "Maruti Suzuki Ertiga",Colour:"RED",YearOfManufacture: "2019",CurrentState: "ReadyForSale"},
        Car{CarId: "2",EngineNumber: "8450210",ChassisNumber: "1HGBH41JXMN109186",Model: "New Maruti Suzuki Swift",Colour:"SILVER",YearOfManufacture: "2020",CurrentState: "ReadyForSale"},
        Car{CarId: "3",EngineNumber: "1436709",ChassisNumber: "1GNCS18Z3M0115610",Model: "New Maruti Suzuki Swift",Colour:"ORANGE",YearOfManufacture: "2018",CurrentState: "ReadyForSale"},
        Car{CarId: "4",EngineNumber: "1480235",ChassisNumber: "1HGYT13MBYC918237",Model: "Maruti Suzuki Ertiga",Colour:"GOLD",YearOfManufacture: "2019",CurrentState: "ReadyForSale"},
        Car{CarId: "5",EngineNumber: "3602154",ChassisNumber: "1GBTS14X5M0764610",Model: "Maruti Suzuki Baleno",Colour:"SILVER",YearOfManufacture: "2020",CurrentState: "ReadyForSale"},
        Car{CarId: "6",EngineNumber: "8002145",ChassisNumber: "15BVT64X5M0867510",Model: "Maruti Suzuki Alto 800",Colour:"BLACK",YearOfManufacture: "2017",CurrentState: "ReadyForSale"},
        Car{CarId: "7",EngineNumber: "1020354",ChassisNumber: "19UB3F77GA0011989",Model: "Maruti Suzuki Dzire",Colour:"SILVER",YearOfManufacture: "2020",CurrentState: "ReadyForSale"},
        Car{CarId: "8",EngineNumber: "5630012",ChassisNumber: "1HGGY12537G284352",Model: "Maruti Suzuki Ignis",Colour:"BURGANDI",YearOfManufacture: "2019",CurrentState: "ReadyForSale"},
        Car{CarId: "9",EngineNumber: "5689084",ChassisNumber: "19RV3G87MH4784100",Model: "Maruti Suzuki Dzire",Colour:"YELLOW",YearOfManufacture: "2018",CurrentState: "ReadyForSale"},
        Car{CarId: "10",EngineNumber: "4853210",ChassisNumber: "1HBYH41UYXG843186",Model: "Maruti Suzuki Baleno",Colour:"GOLD",YearOfManufacture: "2020",CurrentState: "ReadyForSale"},
	}

	i := 0
	indexName := "CarId~EngineNumber~ChassisNumber~Model"

	for i < len(cars) {
		fmt.Println("i is ", i)
		carAsBytes, _ := json.Marshal(cars[i])
		APIstub.PutState(cars[i].CarId, carAsBytes)
		colorNameIndexKey, err := APIstub.CreateCompositeKey(indexName, []string{cars[i].CarId, cars[i].EngineNumber, cars[i].ChassisNumber, cars[i].Model})
	    if err != nil {
			return shim.Error(err.Error())
		}
		value := []byte{0*00}
		APIstub.PutState(colorNameIndexKey,value)
		fmt.Println("Added", cars[i])
		i = i + 1
	}
	return shim.Success(nil)
}

func (s *StatutoryCarRegistration) createCar(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 6 {
		return shim.Error("Incorrect number of arguments. Expecting 5")
	}
	var car = Car{CarId: args[5], EngineNumber: args[0], ChassisNumber: args[1], Model: args[2], Colour: args[3], YearOfManufacture: args[4], CurrentState: "ReadyForSale"}

	carAsBytes, _ := json.Marshal(car)
	APIstub.PutState(args[5], carAsBytes)
	indexName := "CarId~EngineNumber~ChassisNumber~Model"
    colorNameIndexKey, err := APIstub.CreateCompositeKey(indexName, []string{car.CarId, car.EngineNumber, car.ChassisNumber, car.Model})
	    if err != nil {
			return shim.Error(err.Error())
		}
		value := []byte{0*00}
		APIstub.PutState(colorNameIndexKey,value)
	return shim.Success(carAsBytes)
}

func (s *StatutoryCarRegistration) purchaseCar(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2 CarId, DealerName")
	}

	carAsBytes, _ := APIstub.GetState(args[0])
	car := Car{}

	json.Unmarshal(carAsBytes, &car)
	car.Dealer = args[1]
	car.CurrentState = "Booked";

	carAsBytes, _ = json.Marshal(car)
	APIstub.PutState(args[0], carAsBytes)

	return shim.Success(nil)
}

func (s *StatutoryCarRegistration) deliverCar(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting only one argument CarId")
	}

	carAsBytes, _ := APIstub.GetState(args[0])
	car := Car{}

	json.Unmarshal(carAsBytes, &car)
	car.CurrentState = "DeliveredToDealer"

	carAsBytes, _ = json.Marshal(car)
	APIstub.PutState(args[0], carAsBytes)

	return shim.Success(nil)
}

func (s *StatutoryCarRegistration) sellCar(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2 CarId, BuyerName")
	}

	carAsBytes, _ := APIstub.GetState(args[0])
	car := Car{}

	json.Unmarshal(carAsBytes, &car)
	car.Owner =  args[1]
	car.CurrentState = "Purchased"

	carAsBytes, _ = json.Marshal(car)
	APIstub.PutState(args[0], carAsBytes)

	return shim.Success(nil)
}

func (s *StatutoryCarRegistration) insuredCar(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1 argument, CarId")
	}

	carAsBytes, _ := APIstub.GetState(args[0])
	car := Car{}

	json.Unmarshal(carAsBytes, &car)
	car.CurrentState = "Insured"

	carAsBytes, _ = json.Marshal(car)
	APIstub.PutState(args[0], carAsBytes)

	return shim.Success(nil)
}

func (s *StatutoryCarRegistration) registeredCar(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1 argument, CarId")
	}

	carAsBytes, _ := APIstub.GetState(args[0])
	car := Car{}

	json.Unmarshal(carAsBytes, &car)
	car.CurrentState = "Registered"

	carAsBytes, _ = json.Marshal(car)
	APIstub.PutState(args[0], carAsBytes)

	return shim.Success(nil)
}

func (s *StatutoryCarRegistration) deliverCarToCustomer(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1 argument, CarId")
	}

	carAsBytes, _ := APIstub.GetState(args[0])
	car := Car{}

	json.Unmarshal(carAsBytes, &car)
	car.CurrentState = "DeliveredToCustomer"

	carAsBytes, _ = json.Marshal(car)
	APIstub.PutState(args[0], carAsBytes)

	return shim.Success(nil)
}

func (s *StatutoryCarRegistration) queryCar(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1, CarId")
	} 

	carAsBytes, _ := APIstub.GetState(args[0])
	return shim.Success(carAsBytes)
}

func (s *StatutoryCarRegistration) queryCarByEngineNumber(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1, EngineNumber")
			} 
	queryString := fmt.Sprintf("{\"selector\":{\"EngineNumber\":\"%s\"}}", args[0])

	resultsIterator, err := APIstub.GetQueryResult(queryString)

	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryResults
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- queryCarByEngineNumber:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
	
}

func (s *StatutoryCarRegistration) queryAllCars(APIstub shim.ChaincodeStubInterface) sc.Response {

	startKey := "1"
	endKey := "999"

	resultsIterator, err := APIstub.GetStateByRange(startKey, endKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryResults
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- queryAllCars:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}

// add  below code in the end

func (s *StatutoryCarRegistration) getCarHistory(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1, CarId")
	}
	carID := args[0]

	fmt.Printf("- start getCarHistory for : %s\n", carID)

	resultsIterator, err := APIstub.GetHistoryForKey(carID)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"TxId\":")
		buffer.WriteString("\"")
		buffer.WriteString(response.TxId)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Value\":")
		if response.IsDelete {
			buffer.WriteString("null")
		} else {
			buffer.WriteString(string(response.Value))
		}

		buffer.WriteString(", \"Timestamp\":")
		buffer.WriteString("\"")
		buffer.WriteString(time.Unix(response.Timestamp.Seconds, int64(response.Timestamp.Nanos)).String())
		buffer.WriteString("\"")

		buffer.WriteString(", \"IsDelete\":")
		buffer.WriteString("\"")
		buffer.WriteString(strconv.FormatBool(response.IsDelete))
		buffer.WriteString("\"")
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- getCarHistory returning:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())

}

func (s *StatutoryCarRegistration) deleteCar(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
		var jsonResp string
	var carJSON Car
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1 argument, CarId")
	}
	carID := args[0]
	// to maintain the color~name index, we need to read the marble first and get its color
	valAsbytes, err := APIstub.GetState(carID) //get the marble from chaincode state
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + carID + "\"}"
		return shim.Error(jsonResp)
	} else if valAsbytes == nil {
		jsonResp = "{\"Error\":\"car does not exist: " + carID + "\"}"
		return shim.Error(jsonResp)
	}
	err = json.Unmarshal([]byte(valAsbytes), &carJSON)
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to decode JSON of: " + carID + "\"}"
		return shim.Error(jsonResp)
	}

	err = APIstub.DelState(carID) //remove the marble from chaincode state
	if err != nil {
		return shim.Error("Failed to delete state:" + err.Error())
	}
	return shim.Success(nil)
}

// The main function is only relevant in unit test mode. Only included here for completeness.
func main() {

	// Create a new Smart Contract -Madhav
	err := shim.Start(new(StatutoryCarRegistration))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}

