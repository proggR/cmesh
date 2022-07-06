package services
// import "fmt"
// import "cmesh/provider"
//
// func provider() {
//     fmt.Println("Consensus Provider Loaded")
// }
//
// func attrs_to_str(args array) string{
//   b := new(bytes.Buffer)
//   for key, value := range args {
//       fmt.Fprintf(b, "%s=\"%s\"\n", key, value)
//   }
//   return b.String()
// }
//
// // READ
//
// func consession_ssi_state(iamUser JWT) {
//     fmt.Println("Consensus SSI state for current user session loaded")
// }
//
// func consession_did_state(iamUser JWT) {
//     fmt.Println("Consensus DID state for current user session loaded")
// }
//
// func consession_session_state(iamUser JWT) {
//     fmt.Println("Consensus state for current user session loaded")
// }
//
// func consensus_auth_write_state(iamSession JWT, address string, function string) {
//     fmt.Println("Consensus authorization state checked write permissions for user %s against contract:function %s:%s",[iamSession.public,address,function])
// }
//
// func consensus_auth_read_state(iamSession JWT, address string, function string) {
//     fmt.Println("Consensus authorization state checked read permissions for user %s against contract:function %s:%s",[iamSession.public,address,function])
// }
//
//
// // WRITE
//
// func consession_session_state_write(iamSession JWT) {
//     fmt.Println("Consensus state for current user session %s committed to cluster",[iamSession.public])
// }
//
// func consession_ssi_state_write(iamSession JWT) {
//     fmt.Println("Consensus SSI state for current user session %s committed to cluster",[iamSession.public])
// }
//
// func consession_did_state_write(iamSession JWT) {
//     fmt.Println("Consensus DID state for current user session %s committed to cluster",[iamSession.public])
// }
//
// func consession_did_state_attr_write(iamSession JWT, key string, value string) {
//     fmt.Println("Consensus DID attribute %s state %s for current user session %s committed to cluster",[key,value,iamSession.public])
// }
//
// func consession_did_state_attrs_verify(attrs array, iamSession JWT, iamUser JWT) {
//     fmt.Println("Consensus DID attributes %s verified for %s by current user %s",[attrs_to_str(attrs),iamUser.public,iamSession.public])
// }
