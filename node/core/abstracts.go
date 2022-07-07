package core
import(
)

type IAMIF interface {

}

type ProtectedIF interface {
  Router() RouterIF
  IAM() IAM
}

type ProtectedSeed struct {
  RouterInst RouterIF
}

func (ps *ProtectedSeed) IAM() IAM{
  return ps.RouterInst.IAM()
}

func (ps *ProtectedSeed) Router() RouterIF{
  return ps.RouterInst
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

type RouterIF interface{
  ProtectedIF
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
