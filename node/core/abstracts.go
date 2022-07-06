package core
import(
  // "fmt"
)

type IAMIF interface {

}

// type ServiceLayerIF interface {
//
// }

type IRMAProviderIF interface {
    DIDGen() string
    DIDSession() string
    DIDAuth(bool) string
    // dIDSessionCall() string
    DIDSessionAnswer(int,string,uint32) string
    // dIDSessionConfirm() string
    DIDSessionConsent(int, string, string, uint32) string
    DIDSessionHangup()
    TestProvider() string
}

// type RouterSeed struct{
//   IAM IAM
// }

type RouterIF interface{
  IAM() IAM
  Identify(IAM) 
  Route(string,string) string
  ParseRoute(JWT,string) Route
  Ping() string
  Session() string
  Handshake(bool) string
  TestIAMProvider() string
  TestState()
  TestRegistrar()
}
//
// type ServiceIF struct{
//
// }
//
// type ProviderIF struct{
//
// }
//
// type PortIF struct{
//   In ServiceIF
//   Out ServiceIF
// }
//
// type ServiceServicePort struct{
//   In ServiceIF
//   Out ServiceIF
// }
//
// type ProviderProviderPort struct{
//   In ProviderIF
//   Out ProviderIF
// }
//
// type ServiceProviderPort struct{
//   In ServiceIF
//   Out ProviderIF
// }
//
// type ProviderServicePort struct{
//   In ProviderIF
//   Out ServiceIF
// }
