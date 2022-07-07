package providers

import(
  "fmt"
  core "node/core"
  services "node/services"
)



type ServiceProvider struct{
  services.ServiceProviderSeed
}

func (sp *ServiceProvider) Test(){
  // router_connect_services()
  sp.TestState()
  sp.TestRegistrar()
  d := sp.Dispatcher()
  d.Test()
  // sp.Dispatcher.StateProvider.TestRouterResolution(sp.Dispatcher)
}


func (sp *ServiceProvider) TestState(){
  fmt.Println("  Running State Test Sequence")

  d := sp.Dispatcher()
  iam := d.IAM()
  sP := d.State()
  fmt.Println(fmt.Sprintf(" Router:%s\n", iam))
  // router :=  DispatcherService.Router()
  // fmt.Println(fmt.Sprintf(" Router:%s\n", router))
  // iam :=  router.IAM()
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

func (sp *ServiceProvider) TestRegistrar(){
  fmt.Println("  Running Registrar Test Sequence")

  d := sp.Dispatcher()
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
