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

var MinerService core.MinerIF

func main() {
    fmt.Println("Node Starting...")
    iam := iam_bootstrap()

    fmt.Println("Starting Router Service")
    router := core.Router{}
    router.Identify(iam)

    network_test(router)

    fmt.Println("\n\nRouter Services Started & Handshake Tested.\n")

    fmt.Println(" Initializing Protected Services\n")

    state, router := state_bootstrap(router)
    _, router = registrar_bootstrap(router)
    _, router = events_bootstrap(router)

    fmt.Println("   Running Dispatcher Tests\n")

    router.DispatcherTest()

    fmt.Println("   Dispatcher Tests Completed\n")

    fmt.Println("   Route/Response Tests\n")

    res := router.Route(core.Request{FQMN:"0xR:helloWorldExample.mcom"})
    str := res.String()
    fmt.Println(fmt.Sprintf("  Route/Response Test Results:\n   String: %s\n",str))

    fmt.Println("   Route/Response Tests Complete \n")

    state.TestRouterResolution()

    fmt.Println(" Services Initialized\n\n")

    fmt.Println("CMesh Node & Protected Services Initalized\n:)")

    miner_bootstrap(router)
}

func iam_bootstrap() core.IAM{
    iam := core.IAM{}
    iamp := &iamProvider.IRMAProvider{}
    iamp.Construct()
    return iam.IAMService(iamp)
}

func miner_bootstrap(router core.Router){
  fmt.Println("Initializing CMesh Miner\n")
  MinerService := minerProvider.EventsMiner{}
  r := &router
  MinerService.Connect(r)
  MinerService.Start()
  fmt.Println("CMesh Miner Initialized\n:)")
}

func state_bootstrap(router core.Router) (stateProvider.StateProvider, core.Router){
  fmt.Println("   Initializing State Provider\n")
  sp := &stateProvider.StateProvider{}
  r := &router
  sP := sp.Construct(r)
  fmt.Println("   State Provider Loaded, Connected To Router & Dispatcher\n")
  fmt.Println("   State Bootstrapped\n")
  return sP, router
}

func events_bootstrap(router core.Router) (eventsProvider.EventsProvider, core.Router){
  fmt.Println("   Initializing Event Provider\n")
  ep := &eventsProvider.EventsProvider{}
  r := &router
  eP := ep.Construct(r)
  fmt.Println("   Event Provider Loaded, Connected To Router & Dispatcher\n")
  fmt.Println("   Events Bootstrapped\n")
  return eP, router
}

func registrar_bootstrap(router core.Router) (registrarProvider.RegistrarProvider, core.Router){
  fmt.Println("   Initializing Registrar Service\n")
  reg := &registrarProvider.RegistrarProvider{}
  r := &router
  rP := reg.Construct(r)
  fmt.Println("   Registrar Service Loaded, Connected To Router & Dispatcher\n")
  fmt.Println("   Registrar Bootstrapped\n")
  return rP, router
}

func network_test(router core.Router){
  expectedPingbackString := "blah"
  fmt.Println("\nRouter Service Initialized\n Starting Pingback Test")
  msg := router.Ping()
  fmt.Println(fmt.Sprintf("  Pingback test results:\n   Expecting: %s\n   Have: %s\n",expectedPingbackString, msg))

  if(msg != expectedPingbackString){
    fmt.Println("Pingback Failed. Check Router config and try again.")
    return
  }

  fmt.Println(" Starting Router IAM Provider Test")
  msg = router.TestIAMProvider()
  if(msg == ""){
    fmt.Println("Router IAM Provider Test Failed. Check Router config and try again.")
    return
  }else{
    fmt.Println(fmt.Sprintf("   Router IAM Provider Loaded. DidKey: %s\n",msg))
  }

  fmt.Println(" Starting Router IAM Session Test")
  msg = router.Session()
  fmt.Println(fmt.Sprintf("   Router IAM Session Test results:\n    Response: %s\n", msg))
  if(msg == ""){
    fmt.Println("Router IAM Session Test Failed. Check Router config and try again.")
    return
  } else {
    fmt.Println(fmt.Sprintf("   Router IAM Session Test Completed\n    Response: %s\n", msg))
  }

  fmt.Println(" Starting Router IAM Handshake Test")
  msg = router.Handshake(false)
  fmt.Println(fmt.Sprintf("\n   Router IAM Handshake Test results:\n    Response: %s\n", msg))
  if(msg == ""){
    fmt.Println(" Router IAM Handshake Test Failed. Check Router config and try again.")
    return
  } else {
    fmt.Println(fmt.Sprintf(" Router IAM Handshake Test Completed\n Response:%s\n", msg))
  }
}
