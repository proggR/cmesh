package main
import (
  "fmt"
  core "node/core"
  iamProvider "node/providers/iam/mock"

  registrarProvider "node/providers/registrar"
  stateProvider "node/providers/state/mock"
  eventsProvider "node/providers/events/mock"
  eventsMiner "node/miners/events/mock"
)

var MinerService core.MinerIF

func main() {
    fmt.Println("Node Starting...")
    iam := iam_bootstrap()

    fmt.Println("Starting Router Service")
    router := core.Router{}
    router.Identify(iam)

    network_test(router)

    fmt.Println("\n\nRouter Services Started & Handshake Tested.\n\n      :)\n\n")

    fmt.Println("Initializing Protected Services\n")

    state, router := state_bootstrap(router)
    _, router = registrar_bootstrap(router)
    _, router = events_bootstrap(router)

    fmt.Println(" Protected Services Loaded\n\n      :)\n\n")
    fmt.Println(" Running Dispatcher Tests\n")

    router.DispatcherTest()

    fmt.Println("   Dispatcher Tests Completed\n")

    fmt.Println("   Route/Response Tests\n")

    res := router.Route(core.Request{EventID: 0, FQMN:"0xR:helloWorldExample.mcom"})
    str := res.String()
    fmt.Println(fmt.Sprintf("  Route/Response Test Results:\n   String: %s\n",str))

    fmt.Println("   Route/Response Tests Complete \n")

    state.TestRouterResolution()

    fmt.Println(" Protected Services Tested\n\n      :)\n\n")

    fmt.Println("CMesh Node & Protected Services Initalized\n\n      :)\n\n\n")

    miner_bootstrap(router)

    fmt.Println("\n\nCMesh Miner Initialized\n\nNode Now Operating on Cumulus with DID #<insert when nodes have id>\n   (PS: no its not, this is still just a toy model against mock providers)\n   Local Admin Dashboard @ localhost:<insert future port>\n\n      :)\n\nPPS: Currently looking for fun fullstack/backend/systems/blockchain work so if you're recruiting hmu @ proggR@pm.me before I accidentally remake the internet from scratch out of sheer boredom... after already attempting to make a model for new people/commons owned/operated/underwritten money with fractional.foundation :\\... SOS, pls send halp!\n\n   <3\n\n")
}

func iam_bootstrap() core.IAM{
    iam := core.IAM{}
    iamp := &iamProvider.IRMAProvider{}
    iamp.Construct()
    return iam.IAMService(iamp)
}

func miner_bootstrap(router core.Router) core.MinerIF{
  fmt.Println("Initializing CMesh Miner\n")
  MinerService := eventsMiner.EventsMiner{}
  r := &router
  MinerService.Connect(r)
  MinerService.Start()
  m := &MinerService
  return m
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

func state_bootstrap(router core.Router) (stateProvider.StateProvider, core.Router){
  fmt.Println("   Initializing State Provider\n")
  sp := &stateProvider.StateProvider{}
  r := &router
  sP := sp.Construct(r)
  fmt.Println("    State Provider Loaded, Connected To Router & Dispatcher\n")
  fmt.Println("   State Bootstrapped\n")
  return sP, router
}

func events_bootstrap(router core.Router) (eventsProvider.EventsProvider, core.Router){
  fmt.Println("   Initializing Event Provider\n")
  ep := &eventsProvider.EventsProvider{}
  r := &router
  eP := ep.Construct(r)
  fmt.Println("    Event Provider Loaded, Connected To Router & Dispatcher\n")
  fmt.Println("   Events Bootstrapped\n")
  return eP, router
}

func registrar_bootstrap(router core.Router) (registrarProvider.RegistrarProvider, core.Router){
  fmt.Println("   Initializing Registrar Service\n")
  reg := &registrarProvider.RegistrarProvider{}
  r := &router
  rP := reg.Construct(r)
  fmt.Println("    Registrar Service Loaded, Connected To Router & Dispatcher\n")
  fmt.Println("   Registrar Bootstrapped\n")
  return rP, router
}
