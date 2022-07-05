package main
import (
  "fmt"
  routerService "node/router"
  iamService "node/iam"
)

var IAMService iamService.IAM
var RouterService routerService.Router

func main() {
    expectedPingbackString := "blah"
    fmt.Println("Node Started")
    iam_bootstrap()

    fmt.Println("Starting Router Service")
    RouterService = routerService.Router{IAM:IAMService}

    fmt.Println("\nRouter Service Initialized\n Starting Pingback Test")
    msg := RouterService.TestPing()
    fmt.Println(fmt.Sprintf("  Pingback test results:\n   Expecting: %s\n   Have: %s\n",expectedPingbackString, msg))

    if(msg != expectedPingbackString){
      fmt.Println("Pingback Failed. Check Router config and try again.")
      return
    }

    fmt.Println(" Starting Router IAM Provider Test")
    msg = RouterService.TestIAMProvider()
    if(msg == ""){
      fmt.Println("Router IAM Provider Test Failed. Check Router config and try again.")
      return
    }else{
      fmt.Println(fmt.Sprintf("   Router IAM Provider Loaded. DidKey: %s\n",msg))
    }

    fmt.Println(" Starting Router IAM Session Test")
    msg = RouterService.TestSession()
    fmt.Println(fmt.Sprintf("   Router IAM Session Test results:\n    Response: %s\n", msg))
    if(msg == ""){
      fmt.Println("Router IAM Session Test Failed. Check Router config and try again.")
      return
    } else {
      fmt.Println(fmt.Sprintf("   Router IAM Session Test Completed\n    Response: %s\n", msg))
    }

    fmt.Println(" Starting Router IAM Handshake Test")
    msg = RouterService.TestHandshake()
    fmt.Println(fmt.Sprintf("\n   Router IAM Handshake Test results:\n    Response: %s\n", msg))
    if(msg == ""){
      fmt.Println(" Router IAM Handshake Test Failed. Check Router config and try again.")
      return
    } else {
      fmt.Println(fmt.Sprintf(" Router IAM Handshake Test Completed\n Response:%s\n", msg))
    }

    RouterService.InitializeServices()
}

func router_bootstrap(){

}

func iam_bootstrap() {
    IAMService = iamService.IAM{}
    IAMService = IAMService.IAMService()
}
