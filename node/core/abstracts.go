package core
import(
  "fmt"
)

type RouterIF interface{
  Route(string,string) string
  ParseRoute(string)
  Ping() string
  Session() string
  Handshake(bool) string
  TestIAMProvider() string
  TestState()
  TestRegistrar()
}

type ServiceIF struct{

}

type ProviderIF struct{

}

type PortIF struct{
  In ServiceIF
  Out ServiceIF
}

type ServiceServicePort struct{
  In ServiceIF
  Out ServiceIF
}

type ProviderProviderPort struct{
  In ProviderIF
  Out ProviderIF
}

type ServiceProviderPort struct{
  In ServiceIF
  Out ProviderIF
}

type ProviderServicePort struct{
  In ProviderIF
  Out ServiceIF
}
