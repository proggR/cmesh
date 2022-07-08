package miners

import(
  "fmt"
  "hash/fnv"
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
  router.Route(core.Request{FQMN:fmt.Sprintf("0xEW:events.mined:%d.%d",e.hash(fqmn),e.hash(str))})
  return res
}

// UTILITY FUNCTIONS

func (e *EventsMiner) hash(s string) uint32 {
    h := fnv.New32a()
    h.Write([]byte(s))
    return h.Sum32()
}

// END OF UTILITY FUNCTIONS
