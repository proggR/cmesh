package providers

import (
  "fmt"
  services "node/services"
)

// func stahp(){
//   state.Provider()
// }

type StateProvider struct {
  Initialized bool
  Blocks []services.Block
  IAM services.IAM
  Router services.Router
}

func (s *StateProvider) Construct() StateProvider {
  if !s.Initialized {
      s.Initialized = true
      // b := state.Block{Hash:0,ExtraData: "Genesis Block"}
      // s.Blocks = append(s.Blocks, b)
      s.WriteBlock("Genesis Block")
      fmt.Println("   GENESIS BLOCK GENERATED\n")
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

  if len(s.Blocks) == 0 {
    fmt.Println("   WARNING!: NO BLOCKS")
    prevIdx = -1
    hash = 0
    idx = 0;
  } else {
    idx = len(s.Blocks)
    prevIdx = idx-1

    hash = s.Blocks[prevIdx].Hash
    // block = s.Blocks[prevIdx]
    fmt.Println(fmt.Sprintf("      Last Block Index: %d ;\n      Last Block Hash: %d ;\n      MSG:\n       %s ;\n",prevIdx,hash,s.Blocks[prevIdx].ExtraData))
  }

  fmt.Println("   Generating New Block On 'Chain'")

  // b := state.Block{Hash:hash+1, Prev: &block,ExtraData:msg}
  //b := s.generateBlock(hash,msg)
  hash = hash+1
  // blocks := s.Blocks
  b := services.Block{Hash:hash,ExtraData:msg}
  fmt.Println(fmt.Sprintf("   Block %d Generated",b.Hash))

  fmt.Println("   Writing New Block To 'Chain'")
  fmt.Println(fmt.Sprintf("   Block Count Before Append: %d",len(s.Blocks)))
  // blocks = append(blocks, b)
  // s.Blocks = blocks
  s.Blocks = append(s.Blocks,b) //append(blocks, b)
  // s.Blocks[prevIdx] = b
  fmt.Println(fmt.Sprintf("   Block Count After Append: %d",len(s.Blocks)))

  fmt.Println(fmt.Sprintf("   Wrote New Block: Prev Hash of %d, New Hash %d\n",hash-1,hash))
}

// @TODO: switch WriteBlock to work with this to clean things up
func (s *StateProvider) generateBlock(prevIdx uint32, msg string) services.Block {
  b:= services.Block{Hash:uint32(prevIdx+1),ExtraData:msg} //"Generated Block With generateBlock(int)"}
  s.Blocks = append(s.Blocks, b)
  return b
}

func (s *StateProvider) Read(iamSession services.JWT, address string, function string, args []byte, callbackFunction string){
    fmt.Println(fmt.Sprintf("   Session public:%s",iamSession.Public))

    if !s.IAM.ValidatePermissions(iamSession, "state", "mock", fmt.Sprintf("%s:%s", address, function), "read") {
      msg := fmt.Sprintf("   Read permissions for %s:%s denied for JWT %s",address,function,iamSession.Public)
      fmt.Println(msg)
      return
    }
    msg := fmt.Sprintf("State of %s:%s read by %s passing args: %s",address,function,iamSession.Public,string(args))
    fmt.Println(fmt.Sprintf("   %s",msg))
    s.WriteBlock(msg)
}

func (s *StateProvider) Write(iamSession services.JWT, address string, function string, args []byte, callbackFunction string){
    if !s.IAM.ValidatePermissions(iamSession, "state", "mock", fmt.Sprintf("%s:%s", address, function), "write") {
      msg := fmt.Sprintf("   Write permissions for %s:%s denied for JWT %s",address,function,iamSession.Public)
      fmt.Println(msg)
      return
    }
    msg := fmt.Sprintf("State of %s:%s wrote to by %s passing args: %s",address,function,iamSession.Public,string(args))
    fmt.Println(fmt.Sprintf("   %s",msg))
    s.WriteBlock(msg)
}

func deploy(iamSession services.JWT, address string, abi services.ABI, script services.WASMScript, callbackFunction string){
  fmt.Println(fmt.Sprintf("State of %s had contract %s deployed by %s passing callback: %s",address,script.Script,iamSession.Public,callbackFunction))
}

func replace(iamSession services.JWT, address string, addressIngress string, addressEgress string, updateStrategy string, abi services.ABI, script services.WASMScript, callbackFunction string){
  fmt.Println(fmt.Sprintf("State of %s had contract %s replaced by %s using update strategy %s  passing callback: %s",address,script.Script,iamSession.Public,updateStrategy,callbackFunction))
}

func proxy(iamSession services.JWT, proxyStrategy string, addressIngress string, addressEgress string, callbackFunction string){
  fmt.Println(fmt.Sprintf("State of %s proxied to %s by %s using proxy strategy %s passing callback: %s",addressIngress, addressEgress,iamSession.Public,proxyStrategy,callbackFunction))
}
