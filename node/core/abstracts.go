package core
import(
)

type IAMIF interface {

}

type JSONIF interface{

}

type JSON struct{

}

type RequestIF interface {
  Identify(JWT)
  JWT() JWT
}

type ResponseIF interface {
  String() string
  JSON() JSON
  Body() ResponseBodyIF
}

type ResponseServiceIF interface {

}

type ResponseBodyIF interface {

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
  Dispatcher() DispatcherIF
  HasDispatcher() bool
  Attach(DispatcherIF)
  Dispatch(Route)
  Ping() string
  Session() string
  Handshake(bool) string
  TestIAMProvider() string
  TestState()
  TestRegistrar()
}


type DispatcherIF interface {
    ProtectedIF
    IsInitialized() bool
    Init()
    Dispatch()
    Connect(RouterIF)
    Test()
    State()StateProviderIF
    SetRoute(Route)
    SetState(StateProviderIF)
    Registrar()RegistrarIF
    SetRegistrar(RegistrarIF)
}


type ServiceProviderIF interface {
  ProtectedIF
  Connect(RouterIF) ServiceProviderIF
  Attach(DispatcherIF)
  Test()
  Service() string
  Dispatcher() DispatcherIF
}

type StateProviderIF interface{
  ServiceProviderIF
  Read(JWT, string, string, []byte, string)
  Write(JWT, string, string, []byte, string)
  TestRouterResolution()
}
