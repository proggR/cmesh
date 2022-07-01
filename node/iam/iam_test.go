package iam
import (
  "fmt"
  // "hash/fnv"
  "testing"
  // "regexp"
)
// import "assembly"
// import "consensus"
// import "events"

// import "state"

func TestDIDSessionUnauthed(t *testing.T){
  var callString string = DIDSession()
  fmt.Println(fmt.Sprintf("Session Request Call ID: %s",callString))
  var expectedCallString string = "0xDID:0:0:670539951"
  if callString != expectedCallString {
    t.Fatalf(fmt.Sprintf("DIDSession() Unauthed Test Fail: Expected %s, instead got %s", expectedCallString, callString))
  }
}

func TestDIDSessionInvalidAnswer(t *testing.T){
  var callString string = DIDSession()
  fmt.Println(fmt.Sprintf("Session Request Call ID: %s",callString))
  if callString == "" {
    t.Fatalf(fmt.Sprintf("DIDSession() Invalid Answer Test Fail: Empty Callstring"))
    return
  }

  fmt.Println("Client: Answering Call")
  fmt.Println("Client: Fake Answer")
  var fakeConfirmationString string = DIDSessionAnswer(0,callString,9001)
  fmt.Println(fmt.Sprintf("Session Fake Answer Confirmation ID: %s",fakeConfirmationString))


  var expectedAnswerString string = ""
  if fakeConfirmationString != expectedAnswerString {
    t.Fatalf(fmt.Sprintf("DIDSession() Invalid Answer Test Fail: Expected empty response, instead got %s", fakeConfirmationString))
  }
}

func TestDIDSessionValidAnswer(t *testing.T){
  var callString string = DIDSession()
  fmt.Println(fmt.Sprintf("Session Request Call ID: %s",callString))
  if callString == "" {
    t.Fatalf(fmt.Sprintf("DIDSession() Invalid Answer Test Fail: Empty Callstring"))
    return
  }

  fmt.Println("Client: Answering Call")
  fmt.Println("Client: Fake Answer")
  var confirmationString string = DIDSessionAnswer(0,callString,expectedAnswerSig(callString))
  fmt.Println(fmt.Sprintf("Session Fake Answer Confirmation ID: %s",confirmationString))


  if confirmationString == "" {
    t.Fatalf(fmt.Sprintf("DIDSession() Valid Answer Test Fail: Expected confirmation response, instead got empty response"))
  }
}

func TestDIDSessionInvalidConsent(t *testing.T){
  var callString string = DIDSession()
  fmt.Println(fmt.Sprintf("Session Request Call ID: %s",callString))
  if callString == "" {
    t.Fatalf(fmt.Sprintf("DIDSession() Invalid Consent Test Fail: Empty Callstring"))
    return
  }

  fmt.Println("Client: Answering Call")
  fmt.Println("Client: Fake Answer")
  var confirmationString string = DIDSessionAnswer(0,callString,expectedAnswerSig(callString))
  fmt.Println(fmt.Sprintf("Session Fake Answer Confirmation ID: %s",confirmationString))


  if confirmationString == "" {
    t.Fatalf(fmt.Sprintf("DIDSession() Invalid Consent Test Fail: Expected confirmation response, instead got empty response"))
  }

  fmt.Println("Client: Call Answered")
  fmt.Println("Client: Consenting to Answer Confirmation")

  fmt.Println("Client: Fake Consent")
  var fakeConsentString string = DIDSessionConsent(0,callString,confirmationString, 9001)
  fmt.Println(fmt.Sprintf("Session Fake Consent Confirmation ID: %s",fakeConsentString))

  var expectedAnswerString string = ""
  if fakeConsentString != expectedAnswerString {
    t.Fatalf(fmt.Sprintf("DIDSession() Invalid Consent Test Fail: Expected empty response, instead got %s", fakeConsentString))
  }
}

func TestDIDSessionValidConsent(t *testing.T){
  var callString string = DIDSession()
  fmt.Println(fmt.Sprintf("Session Request Call ID: %s",callString))
  if callString == "" {
    t.Fatalf(fmt.Sprintf("DIDSession() Valid Consent Test Fail: Empty Callstring"))
    return
  }

  fmt.Println("Client: Answering Call")
  fmt.Println("Client: Answer")
  var confirmationString string = DIDSessionAnswer(0,callString,expectedAnswerSig(callString))
  fmt.Println(fmt.Sprintf("Session Answer Confirmation ID: %s",confirmationString))


  if confirmationString == "" {
    t.Fatalf(fmt.Sprintf("DIDSession() Valid Consent Test Fail: Expected confirmation response, instead got empty response"))
  }

  fmt.Println("Client: Call Answered")
  fmt.Println("Client: Consenting to Answer Confirmation")

  fmt.Println("Client: Consent")
  var fakeConsentString string = DIDSessionConsent(0,callString,confirmationString, signConsent(confirmationString))
  fmt.Println(fmt.Sprintf("Session Consent Confirmation ID: %s",fakeConsentString))


  if fakeConsentString == "" {
    t.Fatalf(fmt.Sprintf("DIDSession() Valid Consent Test Fail: Expected empty response, instead got %s", fakeConsentString))
  }
}

//
//
// var ssiKey string = "i am the walrus"
// var invalidKey string = "i am one of the walruses"
//
// func expectedAnswerSig(callString string) uint32{
//     return hash(ssiKey+":"+callString)
// }
//
// func signConsent(confirmString string) uint32 {
//   return hash(ssiKey+":"+confirmString)
// }
//
//
// func hash(s string) uint32 {
//     h := fnv.New32a()
//     h.Write([]byte(s))
//     return h.Sum32()
// }
