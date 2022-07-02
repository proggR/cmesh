package main
import (
  "fmt"
  "hash/fnv"
  iamProvider "node/iam/providers/mock"
  iamService "node/iam"
  stateProvider "node/state/providers/mock"
  // iam "node/iam/providers/mock"
  // "node/events/providers/mock"
  // "node/state/providers/mock"
  // "node/consensus/providers/mock"
  // "node/assembly/providers/wasmi"
)

var ssiKey string = "i am the walrus"
var invalidKey string = "i am one of the walruses"

var IAMService iamService.IAM
var IAMProvider iamProvider.IRMAProvider

var StateProvider stateProvider.StateProvider

func main() {
    fmt.Println("Node Started")
    iamServ := iam_bootstrap()
    stateServ := state_bootstrap(iamServ)
    if(stateServ.Initialized){
      fmt.Println("State Service Initialized. Beginning Node Tests.")
      iam_test()
    // eventServ := events_bootstrap(iamServ, stateServ)
    // assemblyServ := assembly_bootstrap(iamServ, stateServ, eventServ)
    // consensusServ := consensus_bootstrap(iamServ, stateServ, eventServ, assemblyServ)
    // stateServ.EstablishConsensus(consensusServ)
    // iamServ.EstablishConsensus(consensusServ)
    // eventServ.EstablishConsensus(consensusServ)
    // assemblyServ.EstablishConsensus(consensusServ)
    // iamProviderProvider()
    } else {
      fmt.Println("State Service Failed To Initialize.")
    }
}

func iam_bootstrap() iamService.IAM {
  iamp := &iamProvider.IRMAProvider{}

  IAMProvider = iamp.Construct()
  IAMService = iamService.IAM{}
  IAMService.IAMService(iamp)
  IAMProvider.IAMService = IAMService
  return IAMService
}

func state_bootstrap(iamServ iamService.IAM) stateProvider.StateProvider{
  stp := &stateProvider.StateProvider{}
  StateProvider = stp.Construct(iamServ)
  return StateProvider
}

func iam_test(){
  fmt.Println("Client: Beginning IRMA Session Handshake")

  fmt.Println("Client: Test One: Valid Complete Walkthrough")

  var callString string = IAMProvider.DIDSession()
  fmt.Println(fmt.Sprintf("Session Request Call ID: %s",callString))
  if callString == "" {
    return
  }

  fmt.Println("Client: Answering Call")
  fmt.Println("Client: Fake Answer")
  var fakeConfirmationString string = IAMProvider.DIDSessionAnswer(0,callString,9001)
  fmt.Println(fmt.Sprintf("Session Fake Answer Confirmation ID: %s",fakeConfirmationString))

  fmt.Println("Client: Valid Answer")
  var confirmationString string = IAMProvider.DIDSessionAnswer(0,callString,expectedAnswerSig(callString))
  fmt.Println(fmt.Sprintf("Session Answer Confirmation ID: %s",confirmationString))
  if confirmationString == "" {
    return
  }

  fmt.Println("Client: Call Answered")
  fmt.Println("Client: Consenting to Answer Confirmation")

  fmt.Println("Client: Fake Consent")
  var fakeConsentString string = IAMProvider.DIDSessionConsent(0,callString,confirmationString, 9001)
  fmt.Println(fmt.Sprintf("Session Fake Consent Confirmation ID: %s",fakeConsentString))

  fmt.Println("Client: Valid Consent")
  var consentString string = IAMProvider.DIDSessionConsent(0,callString,confirmationString, signConsent(confirmationString))
  fmt.Println(fmt.Sprintf("Client: Consented Session ID: %s",consentString))

  if consentString == "" {
    return
  }

  fmt.Println("Client: Call Consented")

  state_test(consentString)

  IAMProvider.DIDSessionHangup()
}

func state_test(consentString string){
  jwt := iamService.JWT{Public:consentString}
  fmt.Println("Client: Running State Read Check With JWT")
  stateProvider.Provider.Read(jwt, "0x001", "hello_world", []byte{111,112,113,114}, "ping_world")
  fmt.Println("Client: Running State Write Check With JWT")
  stateProvider.Provider.Write(jwt, "0x001", "hello_world", []byte{11,12,13,14}, "pong_world")
  fmt.Println("Client: Running State Write Check With JWT")
  stateProvider.Provider.Write(jwt, "0x001", "hello_world", []byte{11,12,13,14}, "pong_world")
  fmt.Println("Client: Running State Read Check With JWT")
  stateProvider.Provider.Read(jwt, "0x001", "hello_world", []byte{111,112,113,114}, "ping_world")
  fmt.Println("Client: Running State Read Check With JWT")
  stateProvider.Provider.Read(jwt, "0x001", "hello_world", []byte{111,121,131,141}, "ping_world")
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
