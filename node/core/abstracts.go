package core
import(
  // "fmt"
)

type IAMIF interface {

}

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


// type RouterIF interface{
//   Route(string,string) string
//   ParseRoute(string)
//   Ping() string
//   Session() string
//   Handshake(bool) string
//   TestIAMProvider() string
//   TestState()
//   TestRegistrar()
// }
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
