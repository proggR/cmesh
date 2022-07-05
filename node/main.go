package main
import (
  "fmt"
  "hash/fnv"
  routerService "node/router"
  iamService "node/iam"
  stateProvider "node/state/providers/mock"
   iamProvider "node/iam/providers/mock"
  // iam "node/iam/providers/mock"
  // "node/events/providers/mock"
  // "node/state/providers/mock"
  // "node/consensus/providers/mock"
  // "node/assembly/providers/wasmi"
)

var RouterService routerService.Router
var IAMService iamService.IAM
var IAMProvider iamProvider.IRMAProvider // deprecate to service, only exist to satisfy unported test
var StateProvider stateProvider.StateProvider // deprecate to service, only exist to satisfy unported test


var PortIAM chan iamService.IAM
var PortRouter chan routerService.Router
var PortState chan stateProvider.StateProvider

var ssiKey string = "i am the walrus"
var invalidKey string = "i am one of the walruses"

func main() {
    expectedPingbackString := "blah"
    fmt.Println("Node Started")
    IAMService := iamService.IAM{}
    IAMService = IAMService.IAMService()

    fmt.Println("Starting Router Service")
    RouterService = routerService.Router{IAM:IAMService}

    fmt.Println("\nRouter Service Initialized\n Starting Pingback Test")
    msg := RouterService.TestPing()
    fmt.Println(fmt.Sprintf("  Pingback test results:\n   Expecting: %s\n   Have: %s\n",expectedPingbackString, msg))

    if(msg != expectedPingbackString){
      fmt.Println("Pingback Failed. Check Router config and try again.")
      return
    }

    fmt.Println(" Starting Router IAM Provider Test")
    msg = RouterService.TestIAMProvider()
    if(msg == ""){
      fmt.Println("Router IAM Provider Test Failed. Check Router config and try again.")
      return
    }else{
      fmt.Println(fmt.Sprintf("   Router IAM Provider Loaded. DidKey: %s\n",msg))
    }

    fmt.Println(" Starting Router IAM Session Test")
    msg = RouterService.TestSession()
    fmt.Println(fmt.Sprintf("   Router IAM Session Test results:\n    Response: %s\n", msg))
    if(msg == ""){
      fmt.Println("Router IAM Session Test Failed. Check Router config and try again.")
      return
    } else {
      fmt.Println(fmt.Sprintf("   Router IAM Session Test Completed\n    Response: %s\n", msg))
    }

    fmt.Println(" Starting Router IAM Handshake Test")
    msg = RouterService.TestHandshake()
    fmt.Println(fmt.Sprintf("\n Router IAM Handshake Test results:\n   Response: %s\n", msg))
    if(msg == ""){
      fmt.Println("Router IAM Handshake Test Failed. Check Router config and try again.")
      return
    } else {
      fmt.Sprintln(" Router IAM Handshake Test Completed\n Response:%s\n", msg)
    }





    // portInIAM = make(chan iamService.IAM,
    // var PortInRouter chan routerService.Router
    // var PortInState chan stateProvider.StateProvider

    // RouterService.Route("iam","validate")
    // RouterService.Route("state","read")


    // result := make(chan int, 1)
    // go channel_test(result)
    //
    // value := <-result
    // fmt.Println(fmt.Sprintf("RESULT: %d",value))
    //
    // channel_test_add(result)
    //
    // value = <-result
    // fmt.Println(fmt.Sprintf("RESULT: %d",value))
    // go iam_bootstrap()
    // state_bootstrap()
    // if(StateProvider.Initialized){
    //   fmt.Println("State Service Initialized. Beginning Node Tests.")
    //   IAMInst :=
    //   fmt.Println(fmt.Sprintf("TEST VALUE: %s",IAMInst.Test))
    //   // iam_test()
    // // eventServ := events_bootstrap(iamServ, stateServ)
    // // assemblyServ := assembly_bootstrap(iamServ, stateServ, eventServ)
    // // consensusServ := consensus_bootstrap(iamServ, stateServ, eventServ, assemblyServ)
    // // stateServ.EstablishConsensus(consensusServ)
    // // iamServ.EstablishConsensus(consensusServ)
    // // eventServ.EstablishConsensus(consensusServ)
    // // assemblyServ.EstablishConsensus(consensusServ)
    // // iamProviderProvider()
    // } else {
    //   fmt.Println("State Service Failed To Initialize.")
    // }
    // close(result)
}
//
// func channel_test(r chan int){
//   r <- 5
// }
// func channel_test_add(r chan int){
//   t := <- r
//   t = t+7
//   r <- t
// }

func router_bootstrap(){

}

func iam_bootstrap() {
    IAMService = iamService.IAM{}
    IAMService = IAMService.IAMService()
    // PortIAM <- IAMService
}

// func state_bootstrap(){
//   stp := &stateProvider.StateProvider{}
//   // iamp := &iamServ
//   StateProvider = stp.Construct()
// }

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
