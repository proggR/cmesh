package main
import (
  "fmt"
  "node/iam"
  "hash/fnv"
)
// import "assembly"
// import "consensus"
// import "events"

// import "state"

var ssiKey string = "i am the walrus"
var invalidKey string = "i am one of the walruses"

func main() {
    fmt.Println("Node Started")
    iam.Provider()
    iam_test()
//    events.provider()
//    state.provider()
//    consensus.provider()
//    assembly.provider()
    // iam_server(iam.provider)
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

  iam.DIDSessionHangup()
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

//
// func iam_server(provider iam.provider){
//   provider()
//   fmt.Println("IAM Server Started")
// }
//
// func iam_zkp(provider iam.provider){
//   fmt.Println("IAM ZKP generated")
//   fmt.Println("IAM ZKP streamed")
// }
//
// func events_server(provider events.provider){
//   provider()
//   fmt.Println("Event Streaming Server Started")
// }
//
// func consensus_validation(provider consensus.provider){
//   fmt.Println("Consensus validation requested")
// }
//
// func state_read(provider state.provider){
//   fmt.Println("State read request handler")
// }
//
// func state_read_zkp(provider state.provider){
//   fmt.Println("State read ZKP generated")
//   fmt.Println("State read ZKP streamed")
// }
//
// func state_write(provider state.provider){
//   fmt.Println("State write request handler")
// }
//
// func state_write_zkp(provider state.provider){
//   fmt.Println("State write ZKP generated")
//   fmt.Println("State write ZKP streamed")
// }




// FUTURE PROBLEMS

// func assembly_load(provider assembly.provider){
//   fmt.Println("Assembly state loaded into local interpretor")
// }
//
// func assembly_unload(provider assembly.provider){
//   fmt.Println("Assembly state unloaded from local interpretor")
// }
//
// func assembly_run(provider assembly.provider){
//   fmt.Println("Assembly state executed within local interpretor")
// }
//
// func assembly_zkp(provider assembly.provider){
//   fmt.Println("Assembly ZKP generated")
//   fmt.Println("Assembly ZKP streamed")
// }
//
// func cmtp_listener(){
//   fmt.Println("CMTP Server Listening on port 9080")
// }
//
// func cmtp_responder(){
//   fmt.Println("CMTP Server Responded to request on port 9080")
// }
