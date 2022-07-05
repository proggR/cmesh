/*
* NOTE: This is an initial toy rough draft.
* Process to port this to the mock provider implementing this IF started.
* Once finished this should contain only:
* IAM, IRMAProviderIF; provider, service; IAMService(provider)
* Snapshot of toy version captured in providers/toy to test hot swapping providers later
*/

package iam
import (
  "fmt"
  iamProvider "node/iam/providers/mock"
  "hash/fnv"
)


var ssiKey string = "i am the walrus"
var invalidKey string = "i am one of the walruses"


type IAM struct {
  Initialized bool
  Provider IRMAProviderIF
  Test string
  Jwt JWT
  AccountJwts []JWT
  ServiceJwts []JWT
  AppJwts []JWT
  ContractJwts []JWT
  StreamJwts []JWT
  PersonalJwts []JWT
}


type JWT struct {
  Public string
  private string
}

type IAMIF interface {

}

type IRMAProviderIF interface {
    DIDGen() string
    DIDSession() string
    DIDAuth(bool) string
    // dIDSessionCall() string
    DIDSessionAnswer(int,string,uint32) string
    // dIDSessionConfirm() string
    DIDSessionConsent(int, string, string, uint32) string
    DIDSessionHangup()
    TestProvider() string
}

type Condition struct {

}

type UserServiceACL struct {
  service string
  authedForDID string
  authedByDID string
  ACL Permissions
}

type Permissions struct {
  read bool
  write bool
  exec bool
  deploy bool
  proxy bool
  replace bool
  conditions []Condition
}


func (i *IAM) IAMService() IAM {
  iamp := &iamProvider.IRMAProvider{}
  iamp.Construct()
  i.setProvider(iamp)
  i.Test = "blah"
  fmt.Println("SSI/DID IAM Service & Provider Loaded")
  return *i
}

func (i *IAM) setProvider(iamp IRMAProviderIF){
  i.Provider = iamp
}

func (i *IAM) TestProvider() string {
    return i.Provider.TestProvider()
}

func (i *IAM) ValidatePermissions(jwt JWT, component string, serviceProvider string, service string, action string) bool{
  iamSession := i.Provider.DIDSession() // to test valid credentials against IAM service (only valid if valid, else invalid)
  // iamSession := "nope" // to test invalid credentials (even with valid ones)
  // iamSession := "0xDID:0:0:3442982940:0xDID:0:0:3442982940:2217691735:17689483:4255629929" // to test valid credentials (even with invalid ones)
  if jwt.Public == "" || jwt.Public != iamSession {
    fmt.Println("No valid IAM session to validate")
    return false
  } else {
    fmt.Println(fmt.Sprintf("JWT: %s ; IAM: %s", jwt.Public, ""))
  }
  // if action == "write"{
  //   return false
  // }
  return true
}


func (i *IAM) TestHandshake() string{
  fmt.Println("   Client: Beginning IRMA Session Handshake\n")

  var callString string = i.Provider.DIDSession()
  fmt.Println(fmt.Sprintf("   Session Request Call ID: %s\n",callString))
  if callString == "" {
    return callString
  }

  fmt.Println("   Client: Answering Call")
  fmt.Println("   Client: Fake Answer")
  var fakeConfirmationString string = i.Provider.DIDSessionAnswer(0,callString,9001)
  fmt.Println(fmt.Sprintf("   Session Fake Answer Confirmation ID (should be blank): %s\n",fakeConfirmationString))

  fmt.Println("   Client: Valid Answer")
  var confirmationString string = i.Provider.DIDSessionAnswer(0,callString,expectedAnswerSig(callString))
  fmt.Println(fmt.Sprintf("\n   Session Answer Confirmation ID:%s\n",confirmationString))
  if confirmationString == "" {
    return confirmationString
  }

  fmt.Println("   Client: Call Answered\n")
  fmt.Println("   Client: Consenting to Answer Confirmation")

  fmt.Println("   Client: Fake Consent")
  var fakeConsentString string = i.Provider.DIDSessionConsent(0,callString,confirmationString, 9001)
  fmt.Println(fmt.Sprintf("   Session Fake Consent Confirmation ID (should be blank): %s\n",fakeConsentString))

  fmt.Println("   Client: Valid Consent")
  var consentString string = i.Provider.DIDSessionConsent(0,callString,confirmationString, signConsent(confirmationString))
  fmt.Println(fmt.Sprintf("\n   Client: Consented Session ID: %s\n",consentString))

  if consentString == "" {
    return consentString
  }

  fmt.Println("   Client: Call Consented\n")

  // i.Provider.DIDSessionHangup()
  return consentString
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
