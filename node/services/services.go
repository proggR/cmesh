package services

import(
  "fmt"
  core "node/core"
)
//
// var RouterService core.RouterIF
// var StateProvider core.StateProviderIF
// var RegistrarService core.Registrar
// var DispatcherService Dispatcher

type ServiceProviderSeed struct {
  core.ProtectedSeed
  DispatcherInst core.DispatcherIF
  Provider core.ServiceProviderIF
  service string
}

func (sp *ServiceProviderSeed) Service() string {
  return sp.service
}

func (sp *ServiceProviderSeed) Test() {
  fmt.Println("SERVICES TEST")
}

func (sp *ServiceProviderSeed) Attach(dispatcher core.DispatcherIF) {
  sp.DispatcherInst = dispatcher
}

func (sp *ServiceProviderSeed) IsInitialized() bool{
  return sp.IsInitialized()
}


func (sp *ServiceProviderSeed) Connect(router core.RouterIF) core.ServiceProviderIF{
  r := router
  if !r.HasDispatcher()  {
      fmt.Println("    Initializing Dispatcher Service\n")
      d := &Dispatcher{Initialized:true}
      d.Connect(r)
      r.Attach(d)
      fmt.Println("    Dispatcher Service Initialized And Connected To Router\n")
  }


  sp.RouterInst = router
  fmt.Println(fmt.Sprintf("      Attached service provider to dispatcher\n"))

  return sp
}
