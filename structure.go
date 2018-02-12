package main
import "fmt"
import "github.com/hyperledger/fabric/core/chaincode/shim"
type SampleChaincode struct {
}


type student struct {
    name string `json:"name"`
    marks string `json:"marks"`
    dept  string `json:"dept"`
}

func (t *SampleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte,
 error) {
 return nil, nil
}
func (t *SampleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte,
 error) {
 return nil, nil

}
func (t *SampleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte,
 error) {
 return nil, nil
}
func main() {
 err := shim.Start(new(SampleChaincode))
 if err != nil {
 fmt.Println("Could not start SampleChaincode")
 } else {
 fmt.Println("SampleChaincode successfully started")
 }
}

//try to init with 2 elements in json array (this works!) next step doesn't work when i try to write more elements in write method
func (t *ShareInfoCode) init(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
     fmt.Println("Entering CreateLoanApplication")
 if len(args) < 2 {
 fmt.Println("Invalid number of args")
 return nil, errors.New("Expected at least two arguments for loan application creation")
 }
 var loanApplicationId = args[0]
 var loanApplicationInput = args[1]
 err := stub.PutState(loanApplicationId, []byte(loanApplicationInput))
 if err != nil {
 fmt.Println("Could not save loan application to ledger", err)
 return nil, err
 }
 fmt.Println("Successfully saved loan application")
 return nil, nil

    var name = args[0]
    var dept = args[1]
    var marks []string = strings.Split(args[2], "|")
   
    var r1 = student{name, marks, dept}
    var r2 = student{name, marks, dept}
    var res1 [2]student
    res1[0] = r1
    res1[1] = r2
    bytestostore11, _ := json.Marshal(res1)
    _ = stub.PutState(user+"-info", bytestostore11)



    return nil, nil

}

//invoke method body
//nothing seems to happen here, no error and state data is not updated
func (t *ShareInfoCode) Query(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
       var name = args[1]
    var marks string = args[2]
    var dept = args[3]
    var res1 = student{name, marks, dept}

    //get the current state data first
    bytesofdata, _ := stub.GetState(args[0] + "-info")
    var res2 = []student{}
    //unmarshal into struct array from json
    err = json.Unmarshal(bytesofdata, &res2)
    if err != nil {
        return nil, errors.New(err.Error() + "unable to unmarshall state data")
    }

    res2new := make([]student, len(res2)+1)
    if len(res2) > 0 {
        copy(res2new, res2[:len(res2)])
    }

    //add the new shareinfo data from user to the list
    res2new[len(res2)] = res1

    //unmarshal into new json to store in ledger
    bytestosave, _ := json.Marshal(res2new)

    //save to ledger state
    _= stub.PutState(user+"-info", bytestosave)

       return nil, nil
}