package state
import "fmt"
import "cmesh/provider"

func provider() {
    fmt.Println("Distributed State Provider Loaded")
}

func args_to_str(args array) string{
  b := new(bytes.Buffer)
  for key, value := range args {
      fmt.Fprintf(b, "%s=\"%s\"\n", key, value)
  }
  return b.String()
}

func read(iamSession JWT, address string, function string, args array, callbackFunction string){
    fmt.Println("State of %s:%s read by %s passing args: %s",[address,function,iamSession.public,args_to_str(args)])
}

func write(iamSession JWT, address string, function string, args array, callbackFunction string){
    fmt.Println("State of %s:%s wrote to by %s passing args: %s",[address,function,iamSession.public,args_to_str(args)])
}

func deploy(iamSession JWT, address string, ami AMI, script WASMScript, callbackFunction string){
  fmt.Println("State of %s had contract %s deployed by %s passing callback: %s",[address,script.script,iamSession.public,callbackFunction])
}

func replace(iamSession JWT, updateStrategy string, addressIngress string, addressEgress string, ami AMI, script WASMScript, callbackFunction string){
  fmt.Println("State of %s had contract %s replaced by %s using update strategy %s  passing callback: %s",[address,script.script,iamSession.public,updateStrategy,callbackFunction])
}

func proxy(iamSession JWT, proxyStrategy string, addressIngress string, addressEgress string, callbackFunction string){
  fmt.Println("State of %s proxied to %s by %s using proxy strategy %s passing callback: %s",[addressIngress, addressEgress,iamSession.public,proxyStrategy,callbackFunction])
}
