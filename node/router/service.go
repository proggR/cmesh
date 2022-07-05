package router
import(
  "fmt"
  "node/iam"
)

/**
* Proposed Op Code Prefixes
* IAM: "0Ix:"
* State: "0Sx:"
* Events: "0Ex:"
* Assembly: "0Ax:"
* Consensus: "0Cx:"
* Registrar: "0Rx:"
* Router: "0Px:" (P for proxy, since you're routing to another router... a strange notion I'll need to consider)
*
* Then add the resource string, which could vary by service
* IAM: action | FQMN
* State: contract:action(:args)?
* Assembly: version:contract:action(:args)?
* Events: channel:action(:args)?
* Router: FQMN
* Registrar: action:entityHash:FQMN:authority:authoritySig | namedService
* Consensus: service:action #(? haven't researched far enough on raft for IF)

* ### FQMN Examples
* Blockahain Greeter Function With Args: 0Sx:0x03389f0e08b9f:hello_world (without args, should only read from state if function marked as pure, with some kind of cache?)
* Registrar registration mapping helloWorld.mcom to contract: 0Rx:register:hash:0Sx:0x03389f0e08b9f,0x079f9849dac562,874958794857983475893475)
* Registrar registration mapping helloWorldExample.mcom to contract function: 0Rx:register:hash:0Sx:0x03389f0e08b9f:hello_world,0x079f9849dac562,874958794857983475893475)
* And again leveraging the existing named service: 0Rx:register:hash:0Rx:helloWorld.mcom:hello_world,0x079f9849dac562,874958794857983475893475)
* Event Stream For Registar Named Service Registration: 0Ex:0Rx:Registered
* Event Stream For Transferred Event From Named Contract: 0Ex:0Rx:helloworld.mcom:Transferred
* Blockahain Greeter Function With Args: 0Sx:0x03389f0e08b9f:hello_world:["Dexter"] (with args, assumes computation and invokes smart contract through Assembler call)
* Same Greeter Function Via Named Service: 0Rx:helloWorldExample.mcom:["Dexter"] (same rules re: args apply)
* Assembly Exec: 0Ax:~4.1.3:0x03389f0e08b9f:hello_world:["Dexter"] (args passed through a "box" (to define) to the script, which returns into the "box" to the state service
* (^ Note: Assembly should only ever be called via state service, and as such can't be registered with a name, and Assembly should only exist to process state changes requiring processing (define consensused cache strat))
*/

type Router struct {
  IAM iam.IAM
  RouterDID string
  OperatorDID string
  RegistrarTx uint32
  RegistrarSig string
  ZKHash uint32
}

// func (r *Router) Route(fqdn string) {
func (r *Router) Route(service string, action string) string {
  msg := fmt.Sprintf("Routing to %s @ %s",action,service)
  fmt.Println(msg)
  return msg
}

func (r *Router) ParseRoute(fqdn string) {

}

func (r *Router) TestPing() string {
  msg := r.IAM.Test
  fmt.Println(msg)
  return msg
}

func (r *Router) TestSession() string {
  msg := r.IAM.Provider.DIDSession()
  fmt.Println(fmt.Sprintf("   Router Session Test:\n    Response: %s",msg))
  return msg
}

func (r *Router) TestHandshake() string {
    return r.IAM.TestHandshake()//r.IAM.TestProvider()
}

func (r *Router) TestIAMProvider() string {
    return r.IAM.TestProvider()
}
