package services
import(
  "fmt"
  "strings"
  core "node/core"
)

type DispatcherIF interface {
    Dispatch()
    Connect(core.RouterIF)
    Router()core.RouterIF
    IAM()core.IAM
    Test()
}

type Dispatcher struct {
  Route core.Route
  RouterInst core.RouterIF
}


func (d *Dispatcher) Dispatch(){
  iam := d.IAM()
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

func (d *Dispatcher) Connect(router core.RouterIF){
  d.RouterInst = router
}

func (d *Dispatcher) Router() core.RouterIF{
  return d.RouterInst
}

func (d *Dispatcher) IAM() core.IAM{
  return d.RouterInst.IAM()
}

func (d *Dispatcher) Test(){
  d.parse_test_routes()
}


func (d *Dispatcher) parse_test_routes(){
  iam := d.IAM()
  consentString := iam.Provider.DIDSession()
  jwt := core.JWT{Public:consentString}

  fqmn1 := "0xS:0x001:hello_world"
  fqmn2 := "0xR:helloWorld.mcom"
  fqmn3 := "0xR:helloWorldExample.mcom"
  fqmn4 := "0xR:google.com"

  fmt.Println("   Client: Running Router Parse: STATE\n")
  route := RouterService.ParseRoute(jwt, fqmn1)
  fmt.Println(fmt.Sprintf("   Router Response:\n    FQMN: %s\n    Service: %s\n    ResourceString: %s\n",route.FQMN,route.Service,route.ResourceString))

  d.Route = route
  d.Dispatch()

  fmt.Println("   Client: Running Router Parse: REGISTRAR CONTRACT\n")
  route = RouterService.ParseRoute(jwt, fqmn2)
  fmt.Println(fmt.Sprintf("   Router Response:\n    FQMN: %s\n    Service: %s\n    ResourceString: %s\n",route.FQMN,route.Service,route.ResourceString))

  d.Route = route
  d.Dispatch()

  fmt.Println("   Client: Running Router Parse: REGISTRAR FUNCTION\n")
  route = RouterService.ParseRoute(jwt, fqmn3)
  fmt.Println(fmt.Sprintf("   Router Response:\n    FQMN: %s\n    Service: %s\n    ResourceString: %s\n",route.FQMN,route.Service,route.ResourceString))

  d.Route = route
  d.Dispatch()

  fmt.Println("   Client: Running Router Parse: REGISTRAR UNREGISTERED DOMAIN\n")
  route = RouterService.ParseRoute(jwt, fqmn4)
  fmt.Println(fmt.Sprintf("   Router Response:\n    FQMN: %s\n    Service: %s\n    ResourceString: %s\n",route.FQMN,route.Service,route.ResourceString))

  d.Route = route
  d.Dispatch()
}
