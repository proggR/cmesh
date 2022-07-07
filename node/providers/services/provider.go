package providers

import(
  "fmt"
  // core "node/core"
  services "node/services"
)

type ServiceProvider struct{
  services.ServiceProviderSeed
}

func (sp *ServiceProvider) Test(){
  fmt.Println("SERVICE TEST")
}
