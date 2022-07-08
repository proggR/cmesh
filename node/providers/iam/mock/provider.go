package providers

import (
  "fmt"
  "hash/fnv"
)

// Interface type check code. Stashing for later
// type T struct{}
// var _ I = T{}       // Verify that T implements I.
// var _ I = (T)(nil) // Verify that T implements I.
// Example:
// var _ iam.IRMAProviderIF = IRMAProvider{}  // Verify that IRMAProvider implements IRMAProviderIF
// If not implemented: compile time error raised

type IRMAProvider struct {
  Initialized bool
  // IAMService iam.IAM
  verifierSsiAddress string
  verifierSsiKey string
  verifierDidIdx int


  ssiAddress string
  ssiKey string
  didAttribute string

  dids []string
  sessionsCount int
  didPrefix string
  sessionCallPrefix string

  DidKey string
  sessionAlive bool
  IamSession string
}

func (p *IRMAProvider) Construct() IRMAProvider {
  if p.ssiKey == "" {
      p.verifierSsiAddress = "0xSSI:1"
      p.verifierSsiKey = "i ap the egg man"
      p.verifierDidIdx = 1
      p.ssiKey = "i am the walrus"
      p.didAttribute = "name=brando"
      p.sessionsCount = 0
      p.didPrefix ="0xDID:%d"
      p.sessionCallPrefix = "0xDID:%d:%d"
      p.DidKey = "kookookachoo"
      p.sessionAlive = false
      p.IamSession = ""
  }
  return *p
}

// START OF IF IMPLEMENTATION FUNCTIONS
// (including the private ones I had to comment out of the IF... figure out how that works :\ )

func (p *IRMAProvider) TestProvider() string {
  return p.didAttribute
}

func (p *IRMAProvider) DIDGen() string {
  fmt.Println("   Generating DID Identity")
  var idx int = len(p.dids);
  var address string = fmt.Sprintf(p.didPrefix,0,0)
  p.dids = append(p.dids,address)
  fmt.Println("   DID Identity Generated")
  return p.dIDSessionCall(idx)
}

func (p *IRMAProvider) DIDSession() string {
  fmt.Println("   DID Identity Network Session Fetched")
  if !p.sessionAlive {
    fmt.Println("   No Active DID Identity Session")
    // if args != nil && args.AutoAuth {
      // var genIds bool = false
      // if args != nil && args.GenIds {
        var genIds bool = true
      // }
      return p.DIDAuth(genIds)
    // } else { return "" }
  } else {
    fmt.Println("   Active DID Identity Session")
  }
  return p.IamSession
}

func (p *IRMAProvider) DIDAuth(genIds bool) string {
  fmt.Println("   DID Identity Network Auth Handshake Started")

  if p.ssiAddress == "" {
    fmt.Println("   DID Handshake Failed: No Root SSI ID")
    if genIds {
      p.rootSSIGen()
      return p.DIDGen()
    } else { return "" }
  } else if len(p.dids) == 0 {
    fmt.Println("   DID Handshake Failed: No DID Identities")
    if genIds {
      return p.DIDGen()
    } else { return "" }
  }
  return p.dIDSessionCall(0)
}

func (p *IRMAProvider) dIDSessionCall(did int) string {
  fmt.Println(fmt.Sprintf("   Generating DID Identity Session Request #%d",p.sessionsCount+1))
  var callString string = p.genSignCallString(did, p.sessionsCount)
  fmt.Println("   DID Identity Session Request Generated")
  fmt.Println(fmt.Sprintf("     Callstring: %s",callString))
  return callString
}

func (p *IRMAProvider) DIDSessionAnswer(did int, callString string, sig uint32) string {
    var expectedCall string = p.genSignCallString(did,p.sessionsCount)
    var expectedSig uint32 = p.expectedAnswerSig(expectedCall)

    if expectedSig != sig {
      fmt.Println(fmt.Sprintf("   DID Identity Session Invalid Answer Credentials For DID %d Provider DID Attribute Checked(%s) SSI Checked (%s):\n    Expected %d\n    Have %d\n", did, p.didAttribute,p.ssiAddress,expectedSig,sig))
      return ""
    }
    if expectedCall != callString {
      fmt.Println("   DID Identity Session Invalid Call String")
      return ""
    }
    fmt.Println(fmt.Sprintf("   Expected Call: %s",expectedCall))
    var answerAckSig uint32 = p.signAnswer(expectedCall,sig)
    var answerString string = p.genAnswerString(expectedCall,answerAckSig)

    fmt.Println("   DID Identity Session Request Answered")
    return p.dIDSessionConfirm(answerString, sig, answerAckSig)
}

func (p *IRMAProvider) dIDSessionConfirm(answerString string, sig uint32, confirmerSig uint32) string {
  var confirmedSig uint32 = p.signConfirm(answerString, sig, confirmerSig)
  var confirmString string = p.genConfirmString(answerString,confirmedSig)
  fmt.Println(fmt.Sprintf("   Generating DID Identity Session Confirmation #%d ID %d for call %s",p.sessionsCount+1, confirmedSig, answerString))
  return confirmString
}

func (p *IRMAProvider) DIDSessionConsent(did int, callString string, confirmString string, sig uint32) string{
  var expectedCallString string = p.genSignCallString(did,p.sessionsCount)

  if expectedCallString != callString {
    fmt.Println(fmt.Sprintf("   DID Identity Session Invalid Call String:\n    Expected %s\n    Have %s\n",expectedCallString,callString))
    return ""
  }

  var expectedAnswerSig uint32 = p.expectedAnswerSig(expectedCallString)
  var expectedAnswerString string = p.genAnswerString(expectedCallString,p.signAnswer(expectedCallString,expectedAnswerSig))
  var expectedOGSig uint32 = p.hash(p.ssiKey+":"+expectedCallString)
  var expectedAnswerAckSig = p.signAnswer(expectedCallString,expectedOGSig)
  var expectedConfirmSig uint32 = p.signConfirm(expectedAnswerString, expectedOGSig, expectedAnswerAckSig)
  var expectedConfirmString string = p.genConfirmString(expectedAnswerString,expectedConfirmSig)

  if expectedConfirmString != confirmString {
    fmt.Println(fmt.Sprintf("   DID Identity Session Invalid Confirmation String:\n Expected:    %s\n    Have: %s\n", expectedConfirmString, confirmString))
    return ""
  }

  var expectedConsentSig = p.signConsent(confirmString)

  if sig != expectedConsentSig {
    fmt.Println(fmt.Sprintf("   DID Identity Session Invalid Credentials:\n    Expected: %d\n    Have: %d\n", expectedConsentSig, sig))
    return ""
  }

  p.sessionAlive = true
  p.sessionsCount += 1
  fmt.Println(fmt.Sprintf("   DID Identity Network Auth Handshake #%d Complete",p.sessionsCount))

  var handshakeSig = p.signHandShake(callString, confirmString, sig)
  var session = fmt.Sprintf("%s:%s:%d", callString,confirmString,handshakeSig)
  fmt.Println(fmt.Sprintf("   Generating DID Identity Session Confirmation #%d",p.sessionsCount))

  p.IamSession = session

  return session
}

func (p *IRMAProvider) DIDSessionHangup() {
  if p.sessionAlive {
    p.sessionAlive = false
    p.IamSession = ""
    fmt.Println(fmt.Sprintf("   DID Identity Network Session #%d Terminated\n",p.sessionsCount))
  } else {
    fmt.Println("   No DID Identity Network Session To Terminated\n")
  }
}

// END OF IF IMPLEMENTATION FUNCTION

// INTERNAL SSI & DID FUNCTIONS

func (p *IRMAProvider) rootSSIGen() string {
    if p.ssiAddress == ""{
      p.ssiAddress = "0xSSI:0"
      // ssi = true
      fmt.Println(fmt.Sprintf("   Root SSI Identity Generated: Receiver %s",p.ssiAddress))
    }
    return p.ssiAddress
}

// END OF INTERNAL SSI & DID FUNCTIONS

// INTERNAL HANDSHAKE STRING & SIGNING FUNCTIONS

func (p *IRMAProvider) genSignCallString(did int,sessionsCount int) string{
  var callString = p.genCallString(did, sessionsCount)
  return fmt.Sprintf(callString+":%d", p.signCall(callString))
}

func (p *IRMAProvider) genCallString(did int,sessionsCount int) string{
  var callString = fmt.Sprintf(p.sessionCallPrefix,did,sessionsCount)
  return callString
}

func (p *IRMAProvider) signCall(callString string) uint32 {
    return p.hash(p.verifierSsiKey+":"+callString)
}

func (p *IRMAProvider) genAnswerString(expectedCall string, answerSig uint32) string {
  return fmt.Sprintf(expectedCall+":%d", answerSig)
}

func (p *IRMAProvider) signAnswer(expectedCall string, sig uint32) uint32 {
    return p.hash(fmt.Sprintf(p.verifierSsiKey+":"+expectedCall+":%d",sig))
}

func (p *IRMAProvider) expectedAnswerSig(callString string) uint32{
    return p.hash(p.ssiKey+":"+callString)
}

func (p *IRMAProvider) genConfirmString(answerString string, confirmSig uint32) string{
  return fmt.Sprintf("%s:%d", answerString, confirmSig)
}

func (p *IRMAProvider) signConfirm(answerString string,sig uint32,signed uint32) uint32 {
  return p.hash(fmt.Sprintf("%s:%s:%d:%d",p.verifierSsiKey,answerString,sig,signed))
}

func (p *IRMAProvider) signHandShake(callString string, confirmString string, sig uint32) uint32 {
  return p.hash(fmt.Sprintf(p.verifierSsiKey+":"+callString+":"+confirmString+":%d",sig))
}

func (p *IRMAProvider) signConsent(confirmString string) uint32 {
  return p.hash(p.ssiKey+":"+confirmString)
}

// END OF INTERNAL HANDSHAKE STRING & SIGNING FUNCTIONS

// UTILITY FUNCTIONS

func (p *IRMAProvider) hash(s string) uint32 {
    h := fnv.New32a()
    h.Write([]byte(s))
    return h.Sum32()
}

// END OF UTILITY FUNCTIONS

// FUTURE ATTRIBUTE RELATED FUNCTIONS

func (p *IRMAProvider) DIDAttrRead(key string){
    fmt.Println(fmt.Sprintf("   DID Identity Attribute %s Read value %s",key,""))
}

func (p *IRMAProvider) DIDAttrAdd(key string, value string){
    fmt.Println(fmt.Sprintf("   DID Identity Attribute %s Added value %s",key,value))
}

func (p *IRMAProvider) DIDAttrReplace(key string, value string){
    fmt.Println(fmt.Sprintf("   DID Identity Attribute %s Replaced with %s",key,value))
}

func (p *IRMAProvider) DIDAttrDel(key string){
    fmt.Println(fmt.Sprintf("   DID Identity Attribute %s Removed",key))
}

func (p *IRMAProvider) DIDSubscribe(){
    fmt.Println("   DID Identity Requested Auth To Network Service")
}

func (p *IRMAProvider) DIDVerifierApproveAttrs(){
    fmt.Println("   DID Identity Approved Auth For Attributes From Service")
}

func (p *IRMAProvider) DIDVerifierDenyAttrs(){
    fmt.Println("   DID Identity Denied Auth For Attributes From Service")
}

func (p *IRMAProvider) DIDVerifierRevokeAuth(){
    fmt.Println("   DID Identity Revoked Auth And All Attributes From Service")
}

func (p *IRMAProvider) DIDVerifierRevokeAttrs(){
    fmt.Println("   DID Identity Revoked Auth For Attributes From Service")
}

// END OF FUTURE ATTRIBUTE RELATED FUNCTIONS
