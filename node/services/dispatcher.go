package services
import(
  "fmt"
  "strings"
  core "node/core"
)

type DispatcherIF interface {
    core.ProtectedIF
    Dispatch()
    Connect(core.RouterIF)
    Test()
    State()StateProviderIF
    SetRoute(core.Route)
    SetState(StateProviderIF)
    Registrar()core.RegistrarIF
    SetRegistrar(core.RegistrarIF)
}

type Dispatcher struct {
  core.ProtectedSeed
  Initialized bool
  Route core.Route
  StateProvider StateProviderIF
  RegistrarService core.RegistrarIF
}


// func (d *Dispatcher)Provider(provider ServiceProviderIF){
//   if provider.Service() == "0xS:" {
//     p := provider
//     d.State(p)
//   } else if provider.Service() == "0XR:" {
//     // d.RegistrarService = provider
//     fmt.Println("UPCOMING SERVICE, UNABLE TO DISPATCH TO REGISTRAR THIS WAY")
//   } else if provider.Service() == "0XE:" {
//     fmt.Println("UPCOMING SERVICE, UNABLE TO DISPATCH TO EVENTS")
//   } else {
//     fmt.Println(fmt.Sprintf("UNKNOWN SERVICE, UNABLE TO DISPATCH TO: %s\n",provider.Service()))
//   }
// }

func (d *Dispatcher) Test(){
  d.parse_test_routes()
  d.testState()
  d.testRegistrar()
  fmt.Println("   DISPATCHER TEST: OK\n")
  // return "DISPATCHER TEST: OK"
}

func (d *Dispatcher) State() StateProviderIF{
  return d.StateProvider
}

func (d *Dispatcher) Registrar() core.RegistrarIF{
  return d.RegistrarService
}

func (d *Dispatcher) SetRoute(route core.Route) {
  d.Route = route
}

func (d *Dispatcher) SetState(state StateProviderIF) {
  d.StateProvider = state
}

func (d *Dispatcher) SetRegistrar(registrar core.RegistrarIF) {
  d.RegistrarService = registrar
}

func (d *Dispatcher) Dispatch(){
  iam := d.IAM()
  router := d.Router()
  state := d.State()
  registrar := d.Registrar()
  consentString := iam.Provider.DIDSession()
  jwt := core.JWT{Public:consentString}

  if d.Route.Service == "0xS:"{
    s:= strings.Split(d.Route.ResourceString,":")
    contract := s[0]
    function := ""
    if len(s) > 1 {
        function = s[1]
    }
    fmt.Println(fmt.Sprintf("   DISPATCHING TO STATE\n    CONTRACT: %s\n    FUNCTION: %s\n",contract,function))
    state.Read(jwt,contract,function,[]byte{},"")
  } else if d.Route.Service == "0xI:"{
    fmt.Println("   DISPATCHING TO IAM\n")
  } else if d.Route.Service == "0xR:"{
    fmt.Println("   DISPATCHING TO REGISTRAR\n")
    fqmn := registrar.Resolve(jwt, d.Route.ResourceString)
    fmt.Println(fmt.Sprintf("   RESOLVED FQMN: %s\n",fqmn))
    d.Route = router.ParseRoute(jwt,fqmn)
    d.Dispatch()
  }
}

func (d *Dispatcher) Connect(router core.RouterIF){
  d.RouterInst = router
}
//
// func (d *Dispatcher) Router() core.RouterIF{
//   return d.RouterInst
// }
//
// func (d *Dispatcher) IAM() core.IAM{
//   r := d.Router()
//   return r.IAM()
// }

func (d *Dispatcher) testState(){
  fmt.Println("  Running State Test Sequence")

  iam := d.IAM()
  sP := d.State()
  fmt.Println(fmt.Sprintf(" Router:%s\n", iam))

  consentString := iam.Provider.DIDSession()
  jwt := core.JWT{Public:consentString}

  fmt.Println("   Client: Running State Read Check With JWT\n")
  sP.Read(jwt, "0x001", "hello_world", []byte{111,112,113,114}, "ping_world")

  fmt.Println("   Client: Running State Write Check With JWT\n")
  sP.Write(jwt, "0x001", "hello_world", []byte{11,12,13,14}, "pong_world")

  fmt.Println("   Client: Running State Write Check With JWT\n")
  sP.Write(jwt, "0x001", "hello_world", []byte{11,12,13,14}, "pong_world")

  fmt.Println("   Client: Running State Read Check With JWT\n")
  sP.Read(jwt, "0x001", "hello_world", []byte{111,112,113,114}, "ping_world")

  fmt.Println("   Client: Running State Read Check With JWT\n")
  sP.Read(jwt, "0x001", "hello_world", []byte{111,121,131,141}, "ping_world")
}

func (d *Dispatcher) testRegistrar(){
  fmt.Println("  Running Registrar Test Sequence")

  iam :=  d.IAM()
  consentString := iam.Provider.DIDSession()
  jwt := core.JWT{Public:consentString}
  rP := d.Registrar()


  fmt.Println("   Client: Running Registrar Named Contract Registration With JWT\n")
  msg := rP.Register(jwt, "helloWorld.mcom", "0xS:0x001")
  fmt.Println(fmt.Sprintf("   Named Contract Mapping Response: %s\n",msg))

  fmt.Println("   Client: Running Registrar Named Function Registration With JWT\n")
  msg = rP.Register(jwt, "helloWorldExample.mcom", "0xS:0x001:hello_world")
  fmt.Println(fmt.Sprintf("   Named Function Mapping Response: %s\n",msg))

  fmt.Println("   Client: Running Registrar Taken Name Registration With JWT\n")
  msg = rP.Register(jwt, "helloWorldExample.mcom", "0xS:0x001:goodnight_world")
  fmt.Println(fmt.Sprintf("   Named Function Mapping Response (should be blank): %s\n",msg))

  fmt.Println("   Client: Running Registrar Named Contract Resolution With JWT\n")
  msg = rP.Resolve(jwt, "helloWorld.mcom")
  fmt.Println(fmt.Sprintf("   Named Contract FQMN Response: %s\n",msg))

  fmt.Println("   Client: Running Registrar Named Function Resolution With JWT\n")
  msg = rP.Resolve(jwt, "helloWorldExample.mcom")
  fmt.Println(fmt.Sprintf("   Named Function FQMN Response: %s\n",msg))

  fmt.Println("   Client: Running Registrar Unregistered Name Resolution With JWT\n")
  msg = rP.Resolve(jwt, "google.com")
  fmt.Println(fmt.Sprintf("   Named Function FQMN Response: %s\n",msg))
}

func (d *Dispatcher) parse_test_routes(){
  iam := d.IAM()
  router := d.Router()
  consentString := iam.Provider.DIDSession()
  jwt := core.JWT{Public:consentString}

  fqmn1 := "0xS:0x001:hello_world"
  fqmn2 := "0xR:helloWorld.mcom"
  fqmn3 := "0xR:helloWorldExample.mcom"
  fqmn4 := "0xR:google.com"

  fmt.Println("   Client: Running Router Parse: STATE\n")
  route := router.ParseRoute(jwt, fqmn1)
  fmt.Println(fmt.Sprintf("   Router Response:\n    FQMN: %s\n    Service: %s\n    ResourceString: %s\n",route.FQMN,route.Service,route.ResourceString))

  d.Route = route
  d.Dispatch()

  fmt.Println("   Client: Running Router Parse: REGISTRAR CONTRACT\n")
  route = router.ParseRoute(jwt, fqmn2)
  fmt.Println(fmt.Sprintf("   Router Response:\n    FQMN: %s\n    Service: %s\n    ResourceString: %s\n",route.FQMN,route.Service,route.ResourceString))

  d.Route = route
  d.Dispatch()

  fmt.Println("   Client: Running Router Parse: REGISTRAR FUNCTION\n")
  route = router.ParseRoute(jwt, fqmn3)
  fmt.Println(fmt.Sprintf("   Router Response:\n    FQMN: %s\n    Service: %s\n    ResourceString: %s\n",route.FQMN,route.Service,route.ResourceString))

  d.Route = route
  d.Dispatch()

  fmt.Println("   Client: Running Router Parse: REGISTRAR UNREGISTERED DOMAIN\n")
  route = router.ParseRoute(jwt, fqmn4)
  fmt.Println(fmt.Sprintf("   Router Response:\n    FQMN: %s\n    Service: %s\n    ResourceString: %s\n",route.FQMN,route.Service,route.ResourceString))

  d.Route = route
  d.Dispatch()
}
