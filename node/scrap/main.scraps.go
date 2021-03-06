package main // lulz
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


// Potentially useful for reference for the router
// var IAMProvider iamProvider.IRMAProvider // deprecate to service, only exist to satisfy unported test
// var StateProvider stateProvider.StateProvider // deprecate to service, only exist to satisfy unported test

//
// var PortIAM chan iamService.IAM
// var PortRouter chan routerService.Router
// var PortState chan stateProvider.StateProvider

// var ssiKey string = "i am the walrus"
// var invalidKey string = "i am one of the walruses"


// portInIAM = make(chan iamService.IAM,
// var PortInRouter chan routerService.Router
// var PortInState chan stateProvider.StateProvider

// RouterService.Route("iam","validate")
// RouterService.Route("state","read")

// PortIAM <- IAMService

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
//
// func channel_test(r chan int){
//   r <- 5
// }
// func channel_test_add(r chan int){
//   t := <- r
//   t = t+7
//   r <- t
// }

// func iam_test(){
//
//   fmt.Println("Client: Beginning IRMA Session Handshake")
//
//   fmt.Println("Client: Test One: Valid Complete Walkthrough")
//
//   var callString string = IAMProvider.DIDSession()
//   fmt.Println(fmt.Sprintf("Session Request Call ID: %s",callString))
//   if callString == "" {
//     return
//   }
//
//   fmt.Println("Client: Answering Call")
//   fmt.Println("Client: Fake Answer")
//   var fakeConfirmationString string = IAMProvider.DIDSessionAnswer(0,callString,9001)
//   fmt.Println(fmt.Sprintf("Session Fake Answer Confirmation ID: %s",fakeConfirmationString))
//
//   fmt.Println("Client: Valid Answer")
//   var confirmationString string = IAMProvider.DIDSessionAnswer(0,callString,expectedAnswerSig(callString))
//   fmt.Println(fmt.Sprintf("Session Answer Confirmation ID: %s",confirmationString))
//   if confirmationString == "" {
//     return
//   }
//
//   fmt.Println("Client: Call Answered")
//   fmt.Println("Client: Consenting to Answer Confirmation")
//
//   fmt.Println("Client: Fake Consent")
//   var fakeConsentString string = IAMProvider.DIDSessionConsent(0,callString,confirmationString, 9001)
//   fmt.Println(fmt.Sprintf("Session Fake Consent Confirmation ID: %s",fakeConsentString))
//
//   fmt.Println("Client: Valid Consent")
//   var consentString string = IAMProvider.DIDSessionConsent(0,callString,confirmationString, signConsent(confirmationString))
//   fmt.Println(fmt.Sprintf("Client: Consented Session ID: %s",consentString))
//
//   if consentString == "" {
//     return
//   }
//
//   fmt.Println("Client: Call Consented")
//
//   state_test(consentString)
//
//   IAMProvider.DIDSessionHangup()
// }



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
