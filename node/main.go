package main
import (
  "fmt"
  "strings"
  // core "node/core"
  services "node/services"
  // providers "node/providers"
  stateProvider "node/providers/state/mock"
  // iamService "node/iam"
  // routerService "node/router"
  // registrarService "node/registrar"
  // stateProvider "node/state/providers/mock"
)

var IAMService services.IAM
var RouterService services.Router
var StateProvider stateProvider.StateProvider
var RegistrarService services.Registrar

// IAM;
//Router(IAM);
//Currently: Service(IAM, Router)
//Switch to: Service(Router) once parsing/routing is working, making IAM calls via opcode sent to router instead of leveraging object directly?

type Dispatcher struct {
  Route services.Route
}

func main() {
    expectedPingbackString := "blah"
    fmt.Println("Node Started")
    iam_bootstrap()

    fmt.Println("Starting Router Service")
    RouterService = services.Router{IAM:IAMService}

    fmt.Println("\nRouter Service Initialized\n Starting Pingback Test")
    msg := RouterService.TestPing()
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
    msg = RouterService.TestSession()
    fmt.Println(fmt.Sprintf("   Router IAM Session Test results:\n    Response: %s\n", msg))
    if(msg == ""){
      fmt.Println("Router IAM Session Test Failed. Check Router config and try again.")
      return
    } else {
      fmt.Println(fmt.Sprintf("   Router IAM Session Test Completed\n    Response: %s\n", msg))
    }

    fmt.Println(" Starting Router IAM Handshake Test")
    msg = RouterService.TestHandshake()
    fmt.Println(fmt.Sprintf("\n   Router IAM Handshake Test results:\n    Response: %s\n", msg))
    if(msg == ""){
      fmt.Println(" Router IAM Handshake Test Failed. Check Router config and try again.")
      return
    } else {
      fmt.Println(fmt.Sprintf(" Router IAM Handshake Test Completed\n Response:%s\n", msg))
    }

    initialize_protected_services()
    parse_test_routes()
}

func (d *Dispatcher) Dispatch(){
  consentString := IAMService.Provider.DIDSession()
  jwt := services.JWT{Public:consentString}

  if d.Route.Service == "0xS:"{
    s:= strings.Split(d.Route.ResourceString,":")
    contract := s[0]
    function := ""
    if len(s) > 1 {
        function = s[1]
    }
    fmt.Println(fmt.Sprintf("   DISPATCHING TO STATE\n    CONTRACT: %s\n    FUNCTION: %s\n",contract,function))
    StateProvider.Read(jwt,contract,function,[]byte{},"")
  } else if d.Route.Service == "0xI:"{
    fmt.Println("   DISPATCHING TO IAM\n")
  } else if d.Route.Service == "0xR:"{
    fmt.Println("   DISPATCHING TO REGISTRAR\n")
    fqmn := RegistrarService.Resolve(jwt, d.Route.ResourceString)
    fmt.Println(fmt.Sprintf("   RESOLVED FQMN: %s\n",fqmn))
    d.Route = RouterService.ParseRoute(jwt,fqmn)
    d.Dispatch()
  }
}

// func Dispatch(){
//   dispatcher := Dispatcher{Route:r}
//   dispatcher.Dispatch()
// }

func iam_bootstrap() {
    IAMService = services.IAM{}
    IAMService = IAMService.IAMService()
}

func initialize_protected_services(){
  state_bootstrap()
  registrar_bootstrap()
  router_connect_services()
}

func state_bootstrap(){
  fmt.Println(" Initializing State Provider\n")
  StateProvider = stateProvider.StateProvider{IAM:IAMService,Router:RouterService}
  StateProvider = StateProvider.Construct()
  fmt.Println(" State Provider Loaded\n")
}

func registrar_bootstrap(){
  fmt.Println(" Initializing Registrar Service\n")
  RegistrarService = services.Registrar{IAM:IAMService,Router:RouterService}
  RegistrarService = RegistrarService.Construct()
  fmt.Println(" Registrar Service Loaded\n")
}


func router_connect_services(){
  // RouterService.State = StateProvider
  // RouterService.Registrar = RegistrarService
  TestState()
  TestRegistrar()
}

func parse_test_routes(){
  consentString := IAMService.Provider.DIDSession()
  jwt := services.JWT{Public:consentString}

  fqmn1 := "0xS:0x001:hello_world"
  fqmn2 := "0xR:helloWorld.mcom"
  fqmn3 := "0xR:helloWorldExample.mcom"
  fqmn4 := "0xR:google.com"

  fmt.Println("   Client: Running Router Parse: STATE\n")
  route := RouterService.ParseRoute(jwt, fqmn1)
  fmt.Println(fmt.Sprintf("   Router Response:\n    FQMN: %s\n    Service: %s\n    ResourceString: %s\n",route.FQMN,route.Service,route.ResourceString))

  dispatcher := Dispatcher{Route:route}
  dispatcher.Dispatch()

  fmt.Println("   Client: Running Router Parse: REGISTRAR CONTRACT\n")
  route = RouterService.ParseRoute(jwt, fqmn2)
  fmt.Println(fmt.Sprintf("   Router Response:\n    FQMN: %s\n    Service: %s\n    ResourceString: %s\n",route.FQMN,route.Service,route.ResourceString))

  dispatcher = Dispatcher{Route:route}
  dispatcher.Dispatch()

  fmt.Println("   Client: Running Router Parse: REGISTRAR FUNCTION\n")
  route = RouterService.ParseRoute(jwt, fqmn3)
  fmt.Println(fmt.Sprintf("   Router Response:\n    FQMN: %s\n    Service: %s\n    ResourceString: %s\n",route.FQMN,route.Service,route.ResourceString))

  dispatcher = Dispatcher{Route:route}
  dispatcher.Dispatch()

  fmt.Println("   Client: Running Router Parse: REGISTRAR UNREGISTERED DOMAIN\n")
  route = RouterService.ParseRoute(jwt, fqmn4)
  fmt.Println(fmt.Sprintf("   Router Response:\n    FQMN: %s\n    Service: %s\n    ResourceString: %s\n",route.FQMN,route.Service,route.ResourceString))

  dispatcher = Dispatcher{Route:route}
  dispatcher.Dispatch()
}

func TestState(){
  fmt.Println("  Running State Test Sequence")

  consentString := IAMService.Provider.DIDSession()
  jwt := services.JWT{Public:consentString}

  fmt.Println("   Client: Running State Read Check With JWT\n")
  StateProvider.Read(jwt, "0x001", "hello_world", []byte{111,112,113,114}, "ping_world")

  fmt.Println("   Client: Running State Write Check With JWT\n")
  StateProvider.Write(jwt, "0x001", "hello_world", []byte{11,12,13,14}, "pong_world")

  fmt.Println("   Client: Running State Write Check With JWT\n")
  StateProvider.Write(jwt, "0x001", "hello_world", []byte{11,12,13,14}, "pong_world")

  fmt.Println("   Client: Running State Read Check With JWT\n")
  StateProvider.Read(jwt, "0x001", "hello_world", []byte{111,112,113,114}, "ping_world")

  fmt.Println("   Client: Running State Read Check With JWT\n")
  StateProvider.Read(jwt, "0x001", "hello_world", []byte{111,121,131,141}, "ping_world")
}

func TestRegistrar(){
  fmt.Println("  Running Registrar Test Sequence")

  consentString := IAMService.Provider.DIDSession()
  jwt := services.JWT{Public:consentString}

  fmt.Println("   Client: Running Registrar Named Contract Registration With JWT\n")
  msg := RegistrarService.Register(jwt, "helloWorld.mcom", "0xS:0x001")
  fmt.Println(fmt.Sprintf("   Named Contract Mapping Response: %s\n",msg))

  fmt.Println("   Client: Running Registrar Named Function Registration With JWT\n")
  msg = RegistrarService.Register(jwt, "helloWorldExample.mcom", "0xS:0x001:hello_world")
  fmt.Println(fmt.Sprintf("   Named Function Mapping Response: %s\n",msg))

  fmt.Println("   Client: Running Registrar Taken Name Registration With JWT\n")
  msg = RegistrarService.Register(jwt, "helloWorldExample.mcom", "0xS:0x001:goodnight_world")
  fmt.Println(fmt.Sprintf("   Named Function Mapping Response (should be blank): %s\n",msg))

  fmt.Println("   Client: Running Registrar Named Contract Resolution With JWT\n")
  msg = RegistrarService.Resolve(jwt, "helloWorld.mcom")
  fmt.Println(fmt.Sprintf("   Named Contract FQMN Response: %s\n",msg))

  fmt.Println("   Client: Running Registrar Named Function Resolution With JWT\n")
  msg = RegistrarService.Resolve(jwt, "helloWorldExample.mcom")
  fmt.Println(fmt.Sprintf("   Named Function FQMN Response: %s\n",msg))

  fmt.Println("   Client: Running Registrar Unregistered Name Resolution With JWT\n")
  msg = RegistrarService.Resolve(jwt, "google.com")
  fmt.Println(fmt.Sprintf("   Named Function FQMN Response: %s\n",msg))
}
