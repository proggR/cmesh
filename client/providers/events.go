package miners

import(
  "fmt"
  "hash/fnv"
  "os"
  "time"
  // tail "github.com/hpcloud/tail"
  // core "node/core"
  // services "node/services"
)

type EventsMiner struct{
  Initialized bool
  service string
  Transactions []string
}

// var Tail bool

func (e *EventsMiner) Start(){

  // e.Transactions = tx
  e.Mine()
}

func (e *EventsMiner) Mine(){
  file  := "/home/sysadmin/systems/cmesh/node/miners/events/mock/events.requests.log"
  f, err := os.Create(file)
  check(err)
  defer f.Close()
  tx := []string{"0xS:0x001:blah_blah","0xR:helloWorldExample.mcom","0xE:events.state.contracts.0x001.blah_blah:read"}
  start := time.Now()
  txCount := 0
  for i := 0; i < 10000 ; i++ {
    for t := range tx {
        txCount++
        fmt.Println(tx[t])
        e.forward(f,txCount,tx[t])
    }
  }

  elapsed := time.Since(start)
  fmt.Println(fmt.Sprintf("\n\nSent %d transactions in %s \n",txCount,elapsed))
  // for i := range e.Transactions{
  //   e.forward(e.Transactions[i])
  // }
}

func (e *EventsMiner) forward(f *os.File, id int, fqmn string){

  // router := e.Router()
  // res := router.Route(core.Request{FQMN:fqmn})
  // str := res.String()
  // router.Route(core.Request{FQMN:fmt.Sprintf("0xEW:events.mined:%d.%d",e.hash(fqmn),e.hash(str))})
  f.WriteString(fmt.Sprintf("%d;%s\n",id,fqmn))
  // os.WriteFile(file,fqmn)
}

// UTILITY FUNCTIONS



func check(e error) {
  if e != nil {
      panic(e)
  }
}


func (e *EventsMiner) hash(s string) uint32 {
    h := fnv.New32a()
    h.Write([]byte(s))
    return h.Sum32()
}

// END OF UTILITY FUNCTIONS
