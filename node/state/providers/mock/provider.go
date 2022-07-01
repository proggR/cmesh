package state_mock

import (
  "fmt"
  // "hash/fnv"
  "node/state"
  "node/iam"
  "node/assembly"
  //"vendor/cmesh/IRMAProvider"
)

func stahp(){
  state.Provider()
}


var Provider StateProvider = new(StateProvider).Construct()


type StateProvider struct {
  Initialized bool
  Blocks []state.Block
}


func (s StateProvider) Construct() StateProvider {
  if !s.Initialized {
      s.Initialized = true
      b := state.Block{Hash:000,ExtraData: "Genesis Block"}
      s.Blocks = append(s.Blocks,b)
  }
  return s
}


// func args_to_str(args []byte) string{
//   b := new(args.Buffer)
//   for key, value := range args {
//       fmt.Fprintf(b, "%d=\"%b\"\n", key, value)
//   }
//   return b.String()
// }

func (s StateProvider) WriteBlock(msg string){
  var idx int
  var hash uint32
  var block state.Block
  if len(s.Blocks) == 0 {
    idx = -1
    hash = 0
  } else {
    idx = len(s.Blocks)-1
    hash = s.Blocks[idx].Hash
    block = s.Blocks[idx]
  }

  fmt.Println("Generating New Block On 'Chain'")
  b := state.Block{Hash:hash+1,Prev:&block,ExtraData:msg}
  fmt.Println("Writing New Block To 'Chain'")
  s.Blocks = append(s.Blocks, b)
  fmt.Println(fmt.Sprintf("Wrote New Block: Prev Hash of %d, New Hash %d",idx,len(s.Blocks)-1))
}

func (s StateProvider) Read(iamSession iam.JWT, address string, function string, args []byte, callbackFunction string){
    msg := fmt.Sprintf("State of %s:%s read by %s passing args: %s",address,function,iamSession.Public,string(args))
    fmt.Println(msg)
    s.WriteBlock(msg)
}

func (s StateProvider) Write(iamSession iam.JWT, address string, function string, args []byte, callbackFunction string){
    msg := fmt.Sprintf("State of %s:%s wrote to by %s passing args: %s",address,function,iamSession.Public,string(args))
    fmt.Println(msg)
    s.WriteBlock(msg)
}

func deploy(iamSession iam.JWT, address string, abi assembly.ABI, script assembly.WASMScript, callbackFunction string){
  fmt.Println(fmt.Sprintf("State of %s had contract %s deployed by %s passing callback: %s",address,script.Script,iamSession.Public,callbackFunction))
}

func replace(iamSession iam.JWT, address string, addressIngress string, addressEgress string, updateStrategy string, abi assembly.ABI, script assembly.WASMScript, callbackFunction string){
  fmt.Println(fmt.Sprintf("State of %s had contract %s replaced by %s using update strategy %s  passing callback: %s",address,script.Script,iamSession.Public,updateStrategy,callbackFunction))
}

func proxy(iamSession iam.JWT, proxyStrategy string, addressIngress string, addressEgress string, callbackFunction string){
  fmt.Println(fmt.Sprintf("State of %s proxied to %s by %s using proxy strategy %s passing callback: %s",addressIngress, addressEgress,iamSession.Public,proxyStrategy,callbackFunction))
}
