package services

import(
  "fmt"
  core "node/core"
  // stateProvider "node/providers/state/mock"
)


var IAMService core.IAM
var RouterService core.RouterIF
var StateProvider StateProviderIF
var RegistrarService core.Registrar
var DispatcherService Dispatcher

type ServiceProviderIF interface {
  Router() core.RouterIF
  IAM() core.IAM
}

type ServiceLayerIF interface {
  ServiceProviderIF
  // Router core.RouterIF
  Connect(core.RouterIF, StateProviderIF) ServiceLayerIF
  Test() string
}

type ServiceProviderSeed struct {
  RouterInst core.Router
}

func (sp *ServiceProviderSeed) Test() string {
  return "Check"
}


func (sp *ServiceProviderSeed) Router() core.RouterIF{
  return &sp.RouterInst
}

func (sp *ServiceProviderSeed) IAM() core.IAM{
  router := sp.RouterInst
  return router.IAM()
}

func (sp *ServiceProviderSeed) Connect(router core.RouterIF, state StateProviderIF) ServiceLayerIF{
  StateProvider = state
  RouterService = router
  // sp.Router = router
  // state_bootstrap()
  registrar_bootstrap()
  router_connect_services()
  DispatcherService := Dispatcher{}
  r := RouterService
  DispatcherService.Connect(r)
  DispatcherService.Test()
  StateProvider.TestRouterResolution(DispatcherService)
  // parse_test_routes()
  return sp
}

// func state_bootstrap(){
//   fmt.Println(" Initializing State Provider\n")
//   StateProvider = stateProvider.StateProvider{IAM:IAMService,Router:RouterService}
//   StateProvider = StateProvider.Construct()
//   fmt.Println(" State Provider Loaded\n")
// }

func registrar_bootstrap(){
  fmt.Println(" Initializing Registrar Service\n")
  RegistrarService = core.Registrar{IAM:IAMService,Router:RouterService}
  RegistrarService = RegistrarService.Construct()
  fmt.Println(" Registrar Service Loaded\n")
}

func router_connect_services(){
  // RouterService.State = StateProvider
  // RouterService.Registrar = RegistrarService
  TestState()
  TestRegistrar()
}

func TestState(){
  fmt.Println("  Running State Test Sequence")

  iam :=  RouterService.IAM()
  consentString := iam.Provider.DIDSession()
  jwt := core.JWT{Public:consentString}

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

  iam :=  RouterService.IAM()
  consentString := iam.Provider.DIDSession()
  jwt := core.JWT{Public:consentString}

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
