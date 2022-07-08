package providers
/**
* @NOTE
* This Provider does nothing at all/is not used within the system.
*
*   It served a purpose solely for helping walk through the early IoC work, but
*   now is just a vestigial shell to act as an MVP example of a service provider
*   conforming to the ProtectedIF via the embedded ServiceProviderSeed struct.
*/

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
