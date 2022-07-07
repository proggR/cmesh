package services

import(
  "fmt"
  core "node/core"
)

var RouterService core.RouterIF
var StateProvider StateProviderIF
var RegistrarService core.Registrar
var DispatcherService Dispatcher

type ServiceProviderIF interface {
  core.ProtectedIF
  Connect(core.RouterIF) ServiceProviderIF
  Attach(DispatcherIF)
  Test()
  Service() string
  Dispatcher() DispatcherIF
}

type ServiceProviderSeed struct {
  core.ProtectedSeed
  DispatcherInst DispatcherIF
  Provider ServiceProviderIF
  service string
}

func (sp *ServiceProviderSeed) Dispatcher() DispatcherIF {
  return sp.DispatcherInst
}

func (sp *ServiceProviderSeed) Service() string {
  return sp.service
}

func (sp *ServiceProviderSeed) Test() {
  fmt.Println("SEED TEST")
}

func (sp *ServiceProviderSeed) Attach(dispatcher DispatcherIF) {
  sp.DispatcherInst = dispatcher
}

func (sp *ServiceProviderSeed) Connect(router core.RouterIF) ServiceProviderIF{
  if !DispatcherService.Initialized  {
      fmt.Println("    Initializing Dispatcher Service\n")
      DispatcherService = Dispatcher{Initialized:true}
      r := router
      DispatcherService.Connect(r)
      fmt.Println("    Dispatcher Service Initialized And Connected To Router\n")
  }

  ds := &DispatcherService
  sp.RouterInst = router
  sp.Attach(ds)
  fmt.Println(fmt.Sprintf("      Attached service provider to dispatcher\n"))

  return sp
}
