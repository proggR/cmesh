package services

import(
  "fmt"
  core "node/core"
  // stateProvider "node/providers/state/mock"
)



var RouterService core.RouterIF
var StateProvider StateProviderIF
var RegistrarService core.Registrar
var DispatcherService Dispatcher

type ServiceProviderIF interface {
  core.ProtectedIF
  Connect(core.RouterIF) ServiceProviderIF
  Attach(DispatcherIF)
  Test() string
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

func (sp *ServiceProviderSeed) Test() string {
  return "Provided Check: SEED"
}

func (sp *ServiceProviderSeed) Attach(dispatcher DispatcherIF) {
  sp.DispatcherInst = dispatcher
}


// func (sp *ServiceProviderSeed) Router() core.RouterIF{
//   return sp.RouterInst
// }
//
// func (sp *ServiceProviderSeed) IAM() core.IAM{
//   router := sp.RouterInst
//   return router.IAM()
// }

func (sp *ServiceProviderSeed) Connect(router core.RouterIF) ServiceProviderIF{
  // StateProvider = state
  // RouterService = router
  // sp.Router = router
  // state_bootstrap()
  // registrar := registrar_bootstrap()
  //,RegistrarService:registrar

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
  fmt.Println(fmt.Sprintf("    Attached service provider %s to dispatcher\n",sp.service))
  // parse_test_routes()
  return sp
}

// func state_bootstrap(){
//   fmt.Println(" Initializing State Provider\n")
//   StateProvider = stateProvider.StateProvider{IAM:IAMService,Router:RouterService}
//   StateProvider = StateProvider.Construct()
//   fmt.Println(" State Provider Loaded\n")
// }

// func registrar_bootstrap() core.Registrar{
//   fmt.Println(" Initializing Registrar Service\n")
//   RegistrarService = core.Registrar{RouterInst:RouterService}
//   RegistrarService = RegistrarService.Construct()
//   fmt.Println(" Registrar Service Loaded\n")
//   return RegistrarService
// }
//
// func router_connect_services(){
//   // RouterService.State = StateProvider
//   // RouterService.Registrar = RegistrarService
//   TestState()
//   TestRegistrar()
// }
