package miners

import(
  "fmt"
  "hash/fnv"
  tail "github.com/hpcloud/tail"
  core "node/core"
  services "node/services"
)

type EventsMiner struct{
  services.ServiceProviderSeed
  Initialized bool
  service string
  Transactions []string
}

// var Tail bool

func (e *EventsMiner) Start(){
  // tx := []string{"0xS:0x001:blah_blah","0xR:helloWorldExample.mcom","0xE:events.state.contracts.0x001.blah_blah:read"}
  // e.Transactions = tx
  e.Mine()
}

func (e *EventsMiner) Mine(){
  Tail, err := tail.TailFile("/home/sysadmin/systems/cmesh/node/miners/events/mock/events.requests.log", tail.Config{Follow: true})
  if err != nil {
    fmt.Println(fmt.Sprintf("      ERROR: %s",err))
  }
  for line := range Tail.Lines {
      fmt.Println(line.Text)
      e.forward(line.Text)
  }
  // for i := range e.Transactions{
  //   e.forward(e.Transactions[i])
  // }
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
