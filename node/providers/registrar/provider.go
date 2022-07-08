package providers

import(
  "fmt"
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
      r.service = "0xR:"
      r.Init()
      r.Connect(router)
      ro := router
      ro.SetRegistrar(r)
      r.Initialized = true
  }
  return *r
}

func (sp *RegistrarProvider) Test(){
  fmt.Println("REGISTRAR TEST")
}
