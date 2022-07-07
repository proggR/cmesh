package providers

import(
  // "fmt"
  core "node/core"
  services "node/services"
)



type RegistrarProvider struct{
  core.Registrar
  services.ServiceProviderSeed
  Initialized bool
  service string
}

func (r *RegistrarProvider) IAM() core.IAM{
  return r.RouterInst.IAM()
}

func (r *RegistrarProvider) Construct(router core.RouterIF) RegistrarProvider {
  if !r.Initialized {
      // r.RouterInst = router
      r.service = "0xR:"
      r.Connect(router)
      d := r.Dispatcher()
      d.SetRegistrar(r)
      r.Initialized = true
  }
  return *r
}

func (sp *RegistrarProvider) Test(){
  // router_connect_services()
  d := sp.Dispatcher()
  d.Test()
  // sp.Dispatcher.StateProvider.TestRouterResolution(sp.Dispatcher)
}
