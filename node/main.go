package main
import (
  "fmt"
  core "node/core"
  iamProvider "node/providers/iam/mock"

  registrarProvider "node/providers/registrar"
  stateProvider "node/providers/state/mock"
  eventsProvider "node/providers/events/mock"
  minerProvider "node/miners"
)

var RouterService core.Router

var StateProvider stateProvider.StateProvider
var RegistrarProvider registrarProvider.RegistrarProvider
// var ServiceLayer serviceProvider.ServiceProvider

func main() {
    fmt.Println("Node Starting...")
    iam := iam_bootstrap()

    fmt.Println("Starting Router Service")
    RouterService = core.Router{}
    RouterService.Identify(iam)

    network_test()

    fmt.Println("\n\nRouter Services Started & Handshake Tested.\n")

    fmt.Println(" Initializing Protected Services")

    state := state_bootstrap()

    registrar_bootstrap()
    events_bootstrap()

    fmt.Println("   Running Dispatcher Tests\n")

    RouterService.DispatcherTest()

    fmt.Println("   Dispatcher Tests Completed\n")

    fmt.Println("   Route/Response Tests\n")

    res := RouterService.Route(core.Request{FQMN:"0xR:helloWorldExample.mcom"})
    str := res.String()
    fmt.Println(fmt.Sprintf("  Route/Response Test Results:\n   String: %s\n",str))

    fmt.Println("   Route/Response Tests Complete \n")

    state.TestRouterResolution()

    fmt.Println(" Services Initialized\n\n")

    fmt.Println("CMesh Node & Protected Services Initalized\n:)")

    miner_bootstrap()
}

func iam_bootstrap() core.IAM{
    iam := core.IAM{}
    iamp := &iamProvider.IRMAProvider{}
    iamp.Construct()
    return iam.IAMService(iamp)
}

func miner_bootstrap(){
  fmt.Println("Initializing CMesh Miner\n")
  miner := minerProvider.EventsMiner{}
  r := &RouterService
  miner.Connect(r)
  miner.Start()
  fmt.Println("CMesh Miner Initialized\n:)")
}

func state_bootstrap() stateProvider.StateProvider{
  fmt.Println("   Initializing State Provider\n")
  sp := &stateProvider.StateProvider{}
  r := &RouterService
  sP := sp.Construct(r)
  fmt.Println("   State Provider Loaded, Connected To Router & Dispatcher\n")
  fmt.Println("   State Bootstrapped\n")
  return sP
}

func events_bootstrap() eventsProvider.EventsProvider{
  fmt.Println("   Initializing Event Provider\n")
  ep := &eventsProvider.EventsProvider{}
  r := &RouterService
  eP := ep.Construct(r)
  fmt.Println("   Event Provider Loaded, Connected To Router & Dispatcher\n")
  fmt.Println("   Events Bootstrapped\n")
  return eP
}

func registrar_bootstrap() registrarProvider.RegistrarProvider{
  fmt.Println("   Initializing Registrar Service\n")
  reg := &registrarProvider.RegistrarProvider{}
  r := &RouterService
  rP := reg.Construct(r)
  fmt.Println("   Registrar Service Loaded, Connected To Router & Dispatcher\n")
  fmt.Println("   Registrar Bootstrapped\n")
  return rP
}

func network_test(){
  expectedPingbackString := "blah"
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
}
