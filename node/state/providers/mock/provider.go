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


var Provider StateProvider// = new(StateProvider).Construct()


type StateProvider struct {
  Initialized bool
  Blocks []state.Block
  IAM iam.IAM
}


func (s *StateProvider) Construct(iamService iam.IAM) StateProvider {
  if !s.Initialized {
      s.IAM = iamService
      s.Initialized = true
      // b := state.Block{Hash:0,ExtraData: "Genesis Block"}
      // s.Blocks = append(s.Blocks, b)
      s.WriteBlock("Genesis Block")
      fmt.Println("GENESIS BLOCK GENERATED")
  }
  return *s
}


// func args_to_str(args []byte) string{
//   b := new(args.Buffer)
//   for key, value := range args {
//       fmt.Fprintf(b, "%d=\"%b\"\n", key, value)
//   }
//   return b.String()
// }

func (s *StateProvider) WriteBlock(msg string){
  var prevIdx int
  var idx int
  var hash uint32
  // var block state.Block
  if len(s.Blocks) == 0 {
    fmt.Println("WARNING!: NO BLOCKS")
    prevIdx = -1
    hash = 0
    idx = 0;
  } else {
    fmt.Println("NOTICE!: BLOCKS")
    idx = len(s.Blocks)
    prevIdx = idx-1
    fmt.Println(fmt.Sprintf("NOTICE!: LAST BLOCK INDEX: %d ; MSG: %s",prevIdx,s.Blocks[prevIdx].ExtraData))
    hash = s.Blocks[prevIdx].Hash
    // block = s.Blocks[prevIdx]
    fmt.Println(fmt.Sprintf("NOTICE!: BLOCK: %d",hash))
  }

  fmt.Println("Generating New Block On 'Chain'")

  // b := state.Block{Hash:hash+1, Prev: &block,ExtraData:msg}

  //b := s.generateBlock(hash,msg)
  hash = hash+1
  // blocks := s.Blocks
  b := state.Block{Hash:hash,ExtraData:msg}
  fmt.Println(fmt.Sprintf("Block %d Generated",b.Hash))

  fmt.Println("Writing New Block To 'Chain'")
  fmt.Println(fmt.Sprintf("Block Count Before Append: %d",len(s.Blocks)))
  // blocks = append(blocks, b)
  // s.Blocks = blocks
  s.Blocks = append(s.Blocks,b) //append(blocks, b)
  // s.Blocks[prevIdx] = b
  fmt.Println(fmt.Sprintf("Block Count After Append: %d",len(s.Blocks)))

  fmt.Println(fmt.Sprintf("Wrote New Block: Prev Hash of %d, New Hash %d",prevIdx,len(s.Blocks)-1))
}

func (s *StateProvider) generateBlock(prevIdx uint32, msg string) state.Block {
  b:= state.Block{Hash:uint32(prevIdx+1),ExtraData:msg}//"Generated Block With generateBlock(int)"}
  s.Blocks = append(s.Blocks, b)
  return b
}

func (s *StateProvider) Read(iamSession iam.JWT, address string, function string, args []byte, callbackFunction string){
    fmt.Println(fmt.Sprintf("Session public:%s",iamSession.Public))
    if !s.IAM.ValidatePermissions(iamSession, "state", "mock", fmt.Sprintf("%s:%s", address, function), "read") {
      msg := fmt.Sprintf("Read permissions for %s:%s denied for JWT %s",address,function,iamSession.Public)
      fmt.Println(msg)
      return
    }
    msg := fmt.Sprintf("State of %s:%s read by %s passing args: %s",address,function,iamSession.Public,string(args))
    fmt.Println(msg)
    s.WriteBlock(msg)
}

func (s *StateProvider) Write(iamSession iam.JWT, address string, function string, args []byte, callbackFunction string){
    if !s.IAM.ValidatePermissions(iamSession, "state", "mock", fmt.Sprintf("%s:%s", address, function), "write") {
      msg := fmt.Sprintf("Write permissions for %s:%s denied for JWT %s",address,function,iamSession.Public)
      fmt.Println(msg)
      return
    }
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
