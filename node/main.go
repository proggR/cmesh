package main
import (
  "fmt"
  core "node/core"
  // services "node/services"
  // services "node/services"
  // providers "node/providers"
  iamProvider "node/providers/iam/mock"
  serviceProvider "node/providers/services"
  stateProvider "node/providers/state/mock"
  // stateProvider "node/providers/state/mock"
  // iamService "node/iam"
  // routerService "node/router"
  // registrarService "node/registrar"
  // stateProvider "node/state/providers/mock"
)

// var IAMService core.IAM
var RouterService core.Router
// var ServiceLayer serviceProvider.ServiceProvider
var StateProvider stateProvider.StateProvider
// var RegistrarService core.Registrar

// IAM;
//Router(IAM);
//Currently: Service(IAM, Router)
//Switch to: Service(Router) once parsing/routing is working, making IAM calls via opcode sent to router instead of leveraging object directly?

// type Dispatcher struct {
//   Route core.Route
// }

func main() {
    expectedPingbackString := "blah"
    fmt.Println("Node Starting...")
    iam := iam_bootstrap()

    fmt.Println("Starting Router Service")
    RouterService = core.Router{}
    RouterService.Identify(iam)

    fmt.Println("\nRouter Service Initialized\n Starting Pingback Test")
    msg := RouterService.Ping()
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
    msg = RouterService.Session()
    fmt.Println(fmt.Sprintf("   Router IAM Session Test results:\n    Response: %s\n", msg))
    if(msg == ""){
      fmt.Println("Router IAM Session Test Failed. Check Router config and try again.")
      return
    } else {
      fmt.Println(fmt.Sprintf("   Router IAM Session Test Completed\n    Response: %s\n", msg))
    }

    fmt.Println(" Starting Router IAM Handshake Test")
    msg = RouterService.Handshake(false)
    fmt.Println(fmt.Sprintf("\n   Router IAM Handshake Test results:\n    Response: %s\n", msg))
    if(msg == ""){
      fmt.Println(" Router IAM Handshake Test Failed. Check Router config and try again.")
      return
    } else {
      fmt.Println(fmt.Sprintf(" Router IAM Handshake Test Completed\n Response:%s\n", msg))
    }


    fmt.Println("Node Started. Initializing Services")
    // initialize_protected_services()

    stateP := state_bootstrap()
    state := &stateP
    fmt.Println("State Bootstrapped")
    fmt.Println("Connecting Service Layer")
    serp := serviceProvider.ServiceProvider{}
    router := &RouterService
    ServiceLayer := serp.Connect(router, state)
    fmt.Println("Service Layer Connected")
    ServiceLayer.Test()
    // services.BootstrapServices(router)
    // parse_test_routes()
}

// func (d *Dispatcher) Dispatch(){
//   consentString := IAMService.Provider.DIDSession()
//   jwt := core.JWT{Public:consentString}
//
//   if d.Route.Service == "0xS:"{
//     s:= strings.Split(d.Route.ResourceString,":")
//     contract := s[0]
//     function := ""
//     if len(s) > 1 {
//         function = s[1]
//     }
//     fmt.Println(fmt.Sprintf("   DISPATCHING TO STATE\n    CONTRACT: %s\n    FUNCTION: %s\n",contract,function))
//     StateProvider.Read(jwt,contract,function,[]byte{},"")
//   } else if d.Route.Service == "0xI:"{
//     fmt.Println("   DISPATCHING TO IAM\n")
//   } else if d.Route.Service == "0xR:"{
//     fmt.Println("   DISPATCHING TO REGISTRAR\n")
//     fqmn := RegistrarService.Resolve(jwt, d.Route.ResourceString)
//     fmt.Println(fmt.Sprintf("   RESOLVED FQMN: %s\n",fqmn))
//     d.Route = RouterService.ParseRoute(jwt,fqmn)
//     d.Dispatch()
//   }
// }

// func Dispatch(){
//   dispatcher := Dispatcher{Route:r}
//   dispatcher.Dispatch()
// }

func iam_bootstrap() core.IAM{
    iam := core.IAM{}
    iamp := &iamProvider.IRMAProvider{}
    iamp.Construct()
    return iam.IAMService(iamp)
}

func state_bootstrap() stateProvider.StateProvider{
  fmt.Println(" Initializing State Provider\n")
  sp := stateProvider.StateProvider{} //RouterInst:router
  sp = StateProvider.Construct(RouterService)
  fmt.Println(" State Provider Loaded\n")
  return sp
}


// func initialize_protected_services(){
//   state_bootstrap()
//   registrar_bootstrap()
//   router_connect_services()
// }
//
// func state_bootstrap(){
//   fmt.Println(" Initializing State Provider\n")
//   StateProvider = stateProvider.StateProvider{IAM:IAMService,Router:RouterService}
//   StateProvider = StateProvider.Construct()
//   fmt.Println(" State Provider Loaded\n")
// }
//
// func registrar_bootstrap(){
//   fmt.Println(" Initializing Registrar Service\n")
//   RegistrarService = core.Registrar{IAM:IAMService,Router:RouterService}
//   RegistrarService = RegistrarService.Construct()
//   fmt.Println(" Registrar Service Loaded\n")
// }


// func router_connect_services(){
//   // RouterService.State = StateProvider
//   // RouterService.Registrar = RegistrarService
//   TestState()
//   TestRegistrar()
// }

// func parse_test_routes(){
//   consentString := IAMService.Provider.DIDSession()
//   jwt := core.JWT{Public:consentString}
//
//   fqmn1 := "0xS:0x001:hello_world"
//   fqmn2 := "0xR:helloWorld.mcom"
//   fqmn3 := "0xR:helloWorldExample.mcom"
//   fqmn4 := "0xR:google.com"
//
//   fmt.Println("   Client: Running Router Parse: STATE\n")
//   route := RouterService.ParseRoute(jwt, fqmn1)
//   fmt.Println(fmt.Sprintf("   Router Response:\n    FQMN: %s\n    Service: %s\n    ResourceString: %s\n",route.FQMN,route.Service,route.ResourceString))
//
//   dispatcher := Dispatcher{Route:route}
//   dispatcher.Dispatch()
//
//   fmt.Println("   Client: Running Router Parse: REGISTRAR CONTRACT\n")
//   route = RouterService.ParseRoute(jwt, fqmn2)
//   fmt.Println(fmt.Sprintf("   Router Response:\n    FQMN: %s\n    Service: %s\n    ResourceString: %s\n",route.FQMN,route.Service,route.ResourceString))
//
//   dispatcher = Dispatcher{Route:route}
//   dispatcher.Dispatch()
//
//   fmt.Println("   Client: Running Router Parse: REGISTRAR FUNCTION\n")
//   route = RouterService.ParseRoute(jwt, fqmn3)
//   fmt.Println(fmt.Sprintf("   Router Response:\n    FQMN: %s\n    Service: %s\n    ResourceString: %s\n",route.FQMN,route.Service,route.ResourceString))
//
//   dispatcher = Dispatcher{Route:route}
//   dispatcher.Dispatch()
//
//   fmt.Println("   Client: Running Router Parse: REGISTRAR UNREGISTERED DOMAIN\n")
//   route = RouterService.ParseRoute(jwt, fqmn4)
//   fmt.Println(fmt.Sprintf("   Router Response:\n    FQMN: %s\n    Service: %s\n    ResourceString: %s\n",route.FQMN,route.Service,route.ResourceString))
//
//   dispatcher = Dispatcher{Route:route}
//   dispatcher.Dispatch()
// }

// func TestState(){
//   fmt.Println("  Running State Test Sequence")
//
//   consentString := IAMService.Provider.DIDSession()
//   jwt := core.JWT{Public:consentString}
//
//   fmt.Println("   Client: Running State Read Check With JWT\n")
//   StateProvider.Read(jwt, "0x001", "hello_world", []byte{111,112,113,114}, "ping_world")
//
//   fmt.Println("   Client: Running State Write Check With JWT\n")
//   StateProvider.Write(jwt, "0x001", "hello_world", []byte{11,12,13,14}, "pong_world")
//
//   fmt.Println("   Client: Running State Write Check With JWT\n")
//   StateProvider.Write(jwt, "0x001", "hello_world", []byte{11,12,13,14}, "pong_world")
//
//   fmt.Println("   Client: Running State Read Check With JWT\n")
//   StateProvider.Read(jwt, "0x001", "hello_world", []byte{111,112,113,114}, "ping_world")
//
//   fmt.Println("   Client: Running State Read Check With JWT\n")
//   StateProvider.Read(jwt, "0x001", "hello_world", []byte{111,121,131,141}, "ping_world")
// }

// func TestRegistrar(){
//   fmt.Println("  Running Registrar Test Sequence")
//
//   consentString := IAMService.Provider.DIDSession()
//   jwt := core.JWT{Public:consentString}
//
//   fmt.Println("   Client: Running Registrar Named Contract Registration With JWT\n")
//   msg := RegistrarService.Register(jwt, "helloWorld.mcom", "0xS:0x001")
//   fmt.Println(fmt.Sprintf("   Named Contract Mapping Response: %s\n",msg))
//
//   fmt.Println("   Client: Running Registrar Named Function Registration With JWT\n")
//   msg = RegistrarService.Register(jwt, "helloWorldExample.mcom", "0xS:0x001:hello_world")
//   fmt.Println(fmt.Sprintf("   Named Function Mapping Response: %s\n",msg))
//
//   fmt.Println("   Client: Running Registrar Taken Name Registration With JWT\n")
//   msg = RegistrarService.Register(jwt, "helloWorldExample.mcom", "0xS:0x001:goodnight_world")
//   fmt.Println(fmt.Sprintf("   Named Function Mapping Response (should be blank): %s\n",msg))
//
//   fmt.Println("   Client: Running Registrar Named Contract Resolution With JWT\n")
//   msg = RegistrarService.Resolve(jwt, "helloWorld.mcom")
//   fmt.Println(fmt.Sprintf("   Named Contract FQMN Response: %s\n",msg))
//
//   fmt.Println("   Client: Running Registrar Named Function Resolution With JWT\n")
//   msg = RegistrarService.Resolve(jwt, "helloWorldExample.mcom")
//   fmt.Println(fmt.Sprintf("   Named Function FQMN Response: %s\n",msg))
//
//   fmt.Println("   Client: Running Registrar Unregistered Name Resolution With JWT\n")
//   msg = RegistrarService.Resolve(jwt, "google.com")
//   fmt.Println(fmt.Sprintf("   Named Function FQMN Response: %s\n",msg))
// }
