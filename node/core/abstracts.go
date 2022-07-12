package core
import(
)

type IAMIF interface {

}

type JSONIF interface{

}

type JSON struct{

}

type MinerIF interface{
  Start()
  Mine()
}

type RequestIF interface {
  Identify(JWT)
  JWT() JWT
  ID() int
  Fqmn() string
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
  Route(Request) Response
  ParseRoute(JWT,Request) Route
  HasDispatcher() bool
  Attach(DispatcherIF)
  Dispatch(Route) Response
  SetState(StateIF)
  SetRegistrar(RegistrarIF)
  SetEvents(EventsIF)
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
    SetRoute(Route)
    Dispatch() Response
    Connect(RouterIF)
    Record(Route,Response)
    Test()
    State() StateIF
    SetState(StateIF)
    Registrar() RegistrarIF
    SetRegistrar(RegistrarIF)
    Events() EventsIF
    SetEvents(EventsIF)
}

type ServiceProviderIF interface {
  ProtectedIF
  Connect(RouterIF) ServiceProviderIF
  Attach(DispatcherIF)
  Service() string
  IsInitialized() bool
  Test()
}

type StateIF interface{
  ServiceProviderIF
  Read(JWT, string, string, []byte, string)
  Write(JWT, string, string, []byte, string)
  TestRouterResolution()
}

type EventsIF interface{
  ServiceProviderIF
  Read(JWT, string, string) string
  Write(JWT, string, string)
}
