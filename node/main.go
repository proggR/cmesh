package main
import (
  "fmt"
  "hash/fnv"
  "node/iam"
  stateProvider "node/state/providers/mock"
  // iam "node/iam/providers/mock"
  // "node/events/providers/mock"
  // "node/state/providers/mock"
  // "node/consensus/providers/mock"
  // "node/assembly/providers/wasmi"
)

var ssiKey string = "i am the walrus"
var invalidKey string = "i am one of the walruses"

func main() {
    fmt.Println("Node Started")
    // iam.Provider()
    iam_test()
}

func iam_test(){
  fmt.Println("Client: Beginning IRMA Session Handshake")

  fmt.Println("Client: Test One: Valid Complete Walkthrough")

  var callString string = iam.DIDSession()
  fmt.Println(fmt.Sprintf("Session Request Call ID: %s",callString))
  if callString == "" {
    return
  }

  fmt.Println("Client: Answering Call")
  fmt.Println("Client: Fake Answer")
  var fakeConfirmationString string = iam.DIDSessionAnswer(0,callString,9001)
  fmt.Println(fmt.Sprintf("Session Fake Answer Confirmation ID: %s",fakeConfirmationString))

  fmt.Println("Client: Valid Answer")
  var confirmationString string = iam.DIDSessionAnswer(0,callString,expectedAnswerSig(callString))
  fmt.Println(fmt.Sprintf("Session Answer Confirmation ID: %s",confirmationString))
  if confirmationString == "" {
    return
  }

  fmt.Println("Client: Call Answered")
  fmt.Println("Client: Consenting to Answer Confirmation")

  fmt.Println("Client: Fake Consent")
  var fakeConsentString string = iam.DIDSessionConsent(0,callString,confirmationString, 9001)
  fmt.Println(fmt.Sprintf("Session Fake Consent Confirmation ID: %s",fakeConsentString))

  fmt.Println("Client: Valid Consent")
  var consentString string = iam.DIDSessionConsent(0,callString,confirmationString, signConsent(confirmationString))
  fmt.Println(fmt.Sprintf("Client: Consented Session ID: %s",consentString))

  if consentString == "" {
    return
  }

  fmt.Println("Client: Call Consented")

  state_test(consentString)

  iam.DIDSessionHangup()
}

func state_test(consentString string){
  jwt := iam.JWT{Public:consentString}
  fmt.Println("Client: Running State Read Check With JWT")
  stateProvider.Provider.Read(jwt, "0x001", "hello_world", []byte{111,112,113,114}, "ping_world")
  fmt.Println("Client: Running State Write Check With JWT")
  stateProvider.Provider.Write(jwt, "0x001", "hello_world", []byte{11,12,13,14}, "pong_world")
  fmt.Println("Client: Running State Write Check With JWT")
  stateProvider.Provider.Write(jwt, "0x001", "hello_world", []byte{11,12,13,14}, "pong_world")
}

func expectedAnswerSig(callString string) uint32{
    return hash(ssiKey+":"+callString)
}

func signConsent(confirmString string) uint32 {
  return hash(ssiKey+":"+confirmString)
}


func hash(s string) uint32 {
    h := fnv.New32a()
    h.Write([]byte(s))
    return h.Sum32()
}
