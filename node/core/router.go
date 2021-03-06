package core
import(
  "fmt"
  "regexp"
  // stateProvider "node/state/providers/mock"
  // registrarService "node/registrar"
  // iam "node/iam/providers/mock"
  // "node/events/providers/mock"
  // "node/state/providers/mock"
  // "node/consensus/providers/mock"
  // "node/assembly/providers/wasmi"
)

// /**
// * Proposed Op Code Prefixes
// * IAM: "0Ix:"
// * State: "0Sx:"
// * Events: "0Ex:"
// * Assembly: "0Ax:"
// * Consensus: "0Cx:"
// * Registrar: "0Rx:"
// * Router: "0Px:" (P for proxy, since you're routing to another router... a strange notion I'll need to consider)
// *
// * Then add the resource string, which could vary by service
// * IAM: action | FQMN
// * State: contract:action(:args)?
// * Assembly: version:contract:action(:args)?
// * Events: channel:action(:args)?
// * Router: FQMN
// * Registrar: action:entityHash:FQMN:authority:authoritySig | namedService
// * Consensus: service:action #(? haven't researched far enough on raft for IF)
//
// * ### FQMN Examples
// * Blockahain Greeter Function With Args: 0Sx:0x03389f0e08b9f:hello_world (without args, should only read from state if function marked as pure, with some kind of cache?)
// * Registrar registration mapping helloWorld.mcom to contract: 0Rx:register:hash:0Sx:0x03389f0e08b9f,0x079f9849dac562,874958794857983475893475)
// * Registrar registration mapping helloWorldExample.mcom to contract function: 0Rx:register:hash:0Sx:0x03389f0e08b9f:hello_world,0x079f9849dac562,874958794857983475893475)
// * And again leveraging the existing named service: 0Rx:register:hash:0Rx:helloWorld.mcom:hello_world,0x079f9849dac562,874958794857983475893475)
// * Event Stream For Registar Named Service Registration: 0Ex:0Rx:Registered
// * Event Stream For Transferred Event From Named Contract: 0Ex:0Rx:helloworld.mcom:Transferred
// * Blockahain Greeter Function With Args: 0Sx:0x03389f0e08b9f:hello_world:["Dexter"] (with args, assumes computation and invokes smart contract through Assembler call)
// * Same Greeter Function Via Named Service: 0Rx:helloWorldExample.mcom:["Dexter"] (same rules re: args apply)
// * Assembly Exec: 0Ax:~4.1.3:0x03389f0e08b9f:hello_world:["Dexter"] (args passed through a "box" (to define) to the script, which returns into the "box" to the state service
// * (^ Note: Assembly should only ever be called via state service, and as such can't be registered with a name, and Assembly should only exist to process state changes requiring processing (define consensused cache strat))
// */
//

type Request struct {
  EventID int
  FQMN string
  Jwt JWT
}

type Response struct {
  FQMN string
  ResolvedFQMN string
  ResponseString string
  ResponseJSON JSONIF
  ResponseBody ResponseBodyIF
  ResponseCode int
}

type Router struct {
  iam IAM
  dispatcher DispatcherIF
  RouterDID string
  OperatorDID string
  RegistrarTx uint32
  RegistrarSig string
  ZKHash uint32
}

type Route struct {
  FQMN string
  Service string
  Request RequestIF
  Action string
  ResourceString string
  ResponseCode int
}

func (req *Request) ID() int{
  return req.EventID
}

func (req *Request) Fqmn() string {
  return req.FQMN
}

func (req *Request) Identify(jwt JWT){
  req.Jwt = jwt
}

func (req *Request) JWT() JWT {
  return req.Jwt
}


func (res *Response) String() string {
  return res.ResponseString
}

func (res *Response) JSON() JSONIF {
  return res.ResponseJSON
}

func (res *Response) Body() ResponseBodyIF {
  return res.ResponseBody
}

// func (r *Router) InitializeServices(){
//   fmt.Println("Initializing Protected Services\n")
//   // r.state_bootstrap()
//   // r.registrar_bootstrap()
// }

// func (r *Router) Route(fqdn string) {
func (r *Router) Route(req Request) Response {
  msg := fmt.Sprintf("   Routing to %s",req.FQMN)
  fmt.Println(msg)

  route := r.ParseRoute(r.iam.Jwt,req)
  return r.Dispatch(route)
}

func (r *Router) IAM() IAM{
  return r.iam
}

func (r *Router) Router() RouterIF{
  return r
}

func (r *Router) Identify(iam IAM){
  r.iam = iam
}

func (r *Router) Dispatcher() DispatcherIF{
  return r.dispatcher
}

func (r *Router) HasDispatcher() bool{
  return r.dispatcher != nil && r.dispatcher.IsInitialized()
}

func (r *Router) Attach(dispatcher DispatcherIF){
  r.dispatcher = dispatcher
  r.dispatcher.Init()
}

func (r *Router) Dispatch(route Route) Response{
  d := r.dispatcher
  d.SetRoute(route)
  return d.Dispatch()
}

func (r *Router) SetState(state StateIF){
  d := r.dispatcher
  d.SetState(state)
}

func (r *Router) SetRegistrar(registrar RegistrarIF){
  d := r.dispatcher
  d.SetRegistrar(registrar)
}

func (r *Router) SetEvents(events EventsIF){
  d := r.dispatcher
  d.SetEvents(events)
}

/**
* Example image of groups in assets/images/parsing_regex.png
* Group 1: Target Service Opcode
* Group 2: Target Service Unresolved Group up to 3 levels deep of address resolution (generally not used on its own unless you know why you need to)
* Group 3: Target Service Address String
* Group 4: First FQMN Identified
* Group 5: Protocol Opcode of FQMN Identified
* Group 6: Service Adress String for Protocol
* Group 7: same as Group 4, but for second FQMN
* Group 8: same as Group 5, but for second FQMN
* Group 9: same as Group 6, but for second FQMN
* Group 10: same as Group 4, but for third FQMN
* Group 11: same as Group 5, but for third FQMN
* Group 12: same as Group 6, but for third FQMN
* If Group 10 ! empty: process first
* Else If Group 7 ! empty: process first
* Else If Group 4 ! empty: process first
* Else use Group 3 as address string
*/
func (r *Router) ParseRoute(jwt JWT, req Request) Route{
  fqmn := req.FQMN
  /**
  * /^(0xS:|0xR:|0xI:|0xE:)((.*)((0xS:|0xR:|0xI:|0xE:)(.*))((0xS:|0xR:|0xI:|0xE:)(.*))((0xS:|0xR:|0xI:|0xE:)(.*)))/
  */

  fmt.Println(fmt.Sprintf("   Matching %s",fqmn))
  // reg, _ := regexp.Compile("/^(0xS:|0xR:|0xI:|0xE:)((.*)((0xS:|0xR:|0xI:|0xE:)(.*))((0xS:|0xR:|0xI:|0xE:)(.*))((0xS:|0xR:|0xI:|0xE:)(.*)))/")
  regex := *regexp.MustCompile(`^(0xS:|0xSW:|0xR:|0xRW:|0xI:|0xE:|0xEW:)((.*)((0xS:|0xSW:|0xR:|0xRW:|0xI:|0xE:|0xEW:)(.*))?((0xS:|0xSW:|0xR:|0xRW:|0xI:|0xE:|0xEW:)(.*))?((0xS:|0xSW:|0xR:|0xRW:|0xI:|0xE:|0xEW:)(.*))?)`)
  res := regex.FindAllStringSubmatch(fqmn, -1)
  request := &req
  if len(res) == 0 {
    fmt.Println("   NO MATCHES? If unexpected, investigate")
  }
  var rt Route = Route{FQMN: fqmn, Request: request, ResponseCode: 400}
  for i := range res {
      //like Java: match.group(1), match.gropu(2), etc
      fmt.Printf("    OpCode: %s,\n    Unresolved Address: %s,\n    Address String: %s,\n    Unresolved SubAddress 1: %s\n", res[i][1], res[i][2], res[i][3], res[i][4])
      rt = Route{FQMN:fqmn,Service: res[i][1],Request:request,ResourceString:res[i][3], ResponseCode: 200}
  }

  return rt
}

func (r *Router) Ping() string {
  iam := r.iam
  msg := iam.Test
  fmt.Println(msg)
  return msg
}

func (r *Router) Session() string {
  iam := r.iam
  msg := iam.Provider.DIDSession()
  fmt.Println(fmt.Sprintf("   Router Session Test:\n    Response: %s",msg))
  return msg
}

func (r *Router) Handshake(disconnectAfter bool) string {
    iam := r.iam
    return iam.TestHandshake()
}

func (r *Router) DispatcherTest() {
    r.dispatcher.Test()
}

func (r *Router) TestIAMProvider() string {
    iam := r.iam
    return iam.TestProvider()
}

func (r *Router) TestRegistrar() {
    // iam := r.iam
    // return iam.TestProvider()
}

func (r *Router) TestState() {
    // iam := r.iam
    // return iam.TestProvider()
}
