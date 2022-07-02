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
  // "hash/fnv"

  //"node/iam/providers/mock"
  //"vendor/cmesh/provider"
)

// var Provider IRMAProviderIF
// var Service IAM

type IAM struct {
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


func (i *IAM) IAMService(p IRMAProviderIF) IAM {
  i.Test = "blah"
  i.Provider = p // = IAM{Provider:p}
  // Service = i
  // Provider = p
  fmt.Println("SSI/DID IAM Service & Provider Loaded")
  return *i
}

func (i *IAM) ValidatePermissions(jwt JWT, component string, serviceProvider string, service string, action string) bool{
  // || jwt.Public != provider.DIDSession()
  // iamSession := "nope"
  // iamSession := i.Provider.DIDSession()
  // @TODO: Get this properly working with IAM response... segfault, late night + early morning = this bug wins... for tonight :\
  iamSession := "0xDID:0:0:3442982940:0xDID:0:0:3442982940:2217691735:17689483:4255629929"
  // iamSession := provider.IamSession
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

// var provider IRMAProviderIF
// var service IAM
//
// var verifierSsiAddress string = "0xSSI:1"
// var verifierSsiKey string = "i am the egg man"
// var verifierDidIdx int = 1
//
//
// var ssiAddress string
// var ssiKey string = "i am the walrus"
// var didAttribute string ="name=brando"
//
// var dids []string
// var sessionsCount int = 0
// var didPrefix string = "0xDID:%d"
// var sessionCallPrefix string = "0xDID:%d:%d"
//
// var didKey string = "kookookachoo"
// var sessionAlive bool = false
// var IamSession string = "";
//
// type SessionArgs struct {
//   AutoAuth bool
//   GenIds bool
// }
//
// func hash(s string) uint32 {
//     h := fnv.New32a()
//     h.Write([]byte(s))
//     return h.Sum32()
// }
//
//
// func IAMService(p IRMAProviderIF) IAM{
//   Provider = p
//   Service = IAM{Provider:p}
//   fmt.Println("SSI/DID IAM Service & Provider Loaded")
//   return Service
// }
//
//
// // func DIDSession(args SessionArgs) string{
// func DIDSession() string{
//     fmt.Println("DID Identity Network Session Fetched")
//     if !sessionAlive {
//       fmt.Println("No Active DID Identity Session")
//       // if args != nil && args.AutoAuth {
//         // var genIds bool = false
//         // if args != nil && args.GenIds {
//           var genIds bool = true
//         // }
//         return DIDAuth(genIds)
//       // } else { return "" }
//     } else {
//       fmt.Println("Active DID Identity Session")
//     }
//     return IamSession
// }
//
// func DIDAuth(genIds bool) string {
//     fmt.Println("DID Identity Network Auth Handshake Started")
//
//     if ssiAddress == "" {
//       fmt.Println("DID Handshake Failed: No Root SSI ID")
//       if genIds {
//         rootSSIGen()
//         return DIDGen()
//       } else { return "" }
//     } else if len(dids) == 0 {
//       fmt.Println("DID Handshake Failed: No DID Identities")
//       if genIds {
//         return DIDGen()
//       } else { return "" }
//     }
//     return dIDSessionCall(0)
// }
//
// func DIDSessionHangup(){
//     if sessionAlive {
//       sessionAlive = false
//       IamSession = ""
//       fmt.Println(fmt.Sprintf("DID Identity Network Session #%d Terminated",sessionsCount))
//     } else {
//       fmt.Println("No DID Identity Network Session To Terminated")
//     }
// }
//
//
// func dIDSessionCall(did int) string {
//     fmt.Println(fmt.Sprintf("Generating DID Identity Session Request #%d",sessionsCount+1))
//     var callString string = genSignCallString(did, sessionsCount)
//     fmt.Println("DID Identity Session Request Generated")
//     return callString
// }
//
// func genSignCallString(did int,sessionsCount int) string{
//   var callString = genCallString(did, sessionsCount)
//   return fmt.Sprintf(callString+":%d", signCall(callString))
// }
//
// func genCallString(did int,sessionsCount int) string{
//   var callString = fmt.Sprintf(sessionCallPrefix,did,sessionsCount)
//   return callString
//   // callString = fmt.Sprintf(callString+":%d", signCall(callString))
//   // return callString
// }
//
// func signCall(callString string) uint32 {
//     return hash(verifierSsiKey+":"+callString)
// }
//
// func DIDSessionAnswer(did int, callString string, sig uint32) string {
//     var expectedCall string = genSignCallString(did,sessionsCount)
//     var expectedSig uint32 = expectedAnswerSig(expectedCall)
//
//     if expectedSig != sig {
//       fmt.Println("DID Identity Session Invalid Credentials")
//       return ""
//     }
//     if expectedCall != callString {
//       fmt.Println("DID Identity Session Invalid Call String")
//       return ""
//     }
//     fmt.Println(fmt.Sprintf("Expected Call: %s",expectedCall))
//     var answerAckSig uint32 = signAnswer(expectedCall,sig)
//     var answerString string = genAnswerString(expectedCall,answerAckSig)
//
//     fmt.Println("DID Identity Session Request Answered")
//     return dIDSessionConfirm(answerString, sig, answerAckSig)
// }
//
// func genAnswerString(expectedCall string, answerSig uint32) string {
//   return fmt.Sprintf(expectedCall+":%d", answerSig)
// }
//
// func signAnswer(expectedCall string, sig uint32) uint32 {
//     return hash(fmt.Sprintf(verifierSsiKey+":"+expectedCall+":%d",sig))
// }
//
// func expectedAnswerSig(callString string) uint32{
//     return hash(ssiKey+":"+callString)
// }
//
// func dIDSessionConfirm(answerString string, sig uint32, confirmerSig uint32) string {
//   var confirmedSig uint32 = signConfirm(answerString, sig, confirmerSig)
//   var confirmString string = genConfirmString(answerString,confirmedSig)
//   fmt.Println(fmt.Sprintf("Generating DID Identity Session Confirmation #%d ID %d for call %s",sessionsCount+1, confirmedSig, answerString))
//   return confirmString
// }
//
// func genConfirmString(answerString string, confirmSig uint32) string{
//   return fmt.Sprintf("%s:%d", answerString, confirmSig)
// }
//
// func signConfirm(answerString string,sig uint32,signed uint32) uint32 {
//   return hash(fmt.Sprintf("%s:%s:%d:%d",verifierSsiKey,answerString,sig,signed))
// }
//
// func DIDSessionConsent(did int, callString string, confirmString string, sig uint32) string{
//   var expectedCallString string = genSignCallString(did,sessionsCount)
//
//   if expectedCallString != callString {
//     fmt.Println("DID Identity Session Invalid Call String")
//     return ""
//   }
//
//
//   //var expectedCallSig uint32 = signCall(expectedCallString)
//   var expectedAnswerSig uint32 = expectedAnswerSig(expectedCallString)
//   var expectedAnswerString string = genAnswerString(expectedCallString,signAnswer(expectedCallString,expectedAnswerSig))
//   var expectedOGSig uint32 = hash(ssiKey+":"+expectedCallString)
//   var expectedAnswerAckSig = signAnswer(expectedCallString,expectedOGSig)
//
//
//   //var answerAckSig uint32 = signAnswer(expectedCall,sig)
//
//   var confirmedSig uint32 = signConfirm(expectedAnswerString, expectedOGSig, expectedAnswerAckSig)
//   var expectedConfirmString string = genConfirmString(expectedAnswerString,confirmedSig)
//
//   //var expectedConfirmString string = fmt.Sprintf(expectedAnswerSig+":%d", expectedConfirmSig)
//
//   if expectedConfirmString != confirmString {
//     fmt.Println(fmt.Sprintf("DID Identity Session Invalid Confirmation String: %s ; Expected: %s", confirmString, expectedConfirmString))
//     return ""
//   }
//
//   var expectedConsentSig = signConsent(confirmString)
//
//
//   if sig != expectedConsentSig {
//     fmt.Println("DID Identity Session Invalid Credentials")
//     return ""
//   }
//
//     sessionAlive = true
//   sessionsCount += 1
//   fmt.Println(fmt.Sprintf("DID Identity Network Auth Handshake #%d Complete",sessionsCount))
//
//
//
//   var handshakeSig = signHandShake(callString, confirmString, sig)
//   var session = fmt.Sprintf("%s:%s:%d", callString,confirmString,handshakeSig)
//   fmt.Println(fmt.Sprintf("Generating DID Identity Session Confirmation #%d",sessionsCount+1))
//
//   IamSession = session
//
//   return session
//
// }
//
// func signHandShake(callString string, confirmString string, sig uint32) uint32 {
//   return hash(fmt.Sprintf(verifierSsiKey+":"+callString+":"+confirmString+":%d",sig))
// }
//
//
// func signConsent(confirmString string) uint32 {
//   return hash(ssiKey+":"+confirmString)
// }
//
// func rootSSIGen() string {
//     if ssiAddress == ""{
//       ssiAddress = "0xSSI:0"
//       // ssi = true
//       fmt.Println("Root SSI Identity Generated")
//     }
//     return ssiAddress
// }
//
// func DIDGen() string {
//     fmt.Println("Generating DID Identity")
//     var idx int = len(dids);
//     var address string = fmt.Sprintf(didPrefix,)
//     dids = append(dids,address)
//     fmt.Println("DID Identity Generated")
//     return dIDSessionCall(idx)
// }
