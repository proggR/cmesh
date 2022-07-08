package miners

import(
  "fmt"
  core "node/core"
  services "node/services"
)

type EventsMiner struct{
  services.ServiceProviderSeed
  Initialized bool
  service string
  Transactions []string
}

func (e *EventsMiner) Start(){
  tx := []string{"0xS:0x001:blah_blah","0xR:helloWorldExample.mcom","0xE:events.state.contracts.0x001.blah_blah:read"}
  e.Transactions = tx
  e.Mine()
}

func (e *EventsMiner) Mine(){
  for i := range e.Transactions{
    e.forward(e.Transactions[i])
  }
}

func (e *EventsMiner) forward(fqmn string) core.Response{
  router := e.Router()
  res := router.Route(core.Request{FQMN:fqmn})
  str := res.String()
  router.Route(core.Request{FQMN:fmt.Sprintf("0xEW:events.mined:%s",str)})
  return res
}
