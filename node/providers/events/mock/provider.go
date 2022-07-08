package providers

import(
  "fmt"
  core "node/core"
  services "node/services"
)

type EventsProvider struct{
  // core.EventsIF
  services.ServiceProviderSeed
  Initialized bool
  service string
}

func (e *EventsProvider) IAM() core.IAM{
  return e.RouterInst.IAM()
}

func (e *EventsProvider) Construct(router core.RouterIF) EventsProvider {
  if !e.Initialized {
      e.service = "0xE:"
      // e.Init()
      e.Connect(router)
      ro := router
      ro.SetEvents(e)
      e.Initialized = true
  }
  return *e
}

func (e *EventsProvider) Test(){
  fmt.Println("EVENTS TEST")
}

func (e *EventsProvider) Read(iamSession core.JWT, channel string, payload string){
    fmt.Println(fmt.Sprintf("   Session public:%s",iamSession.Public))
    iam := e.IAM()
    if !iam.ValidatePermissions(iamSession, "events", "mock", fmt.Sprintf("%s:%s", channel, payload), "read") {
      msg := fmt.Sprintf("   Read permissions for %s:%s denied for JWT %s",channel,payload,iamSession.Public)
      fmt.Println(msg)
      return
    }
    msg := fmt.Sprintf("Events of %s:%s read by %s",channel,payload,iamSession.Public)
    fmt.Println(fmt.Sprintf("   %s",msg))
}

func (e *EventsProvider) Write(iamSession core.JWT, channel string, payload string){
    iam := e.IAM()
    if !iam.ValidatePermissions(iamSession, "events", "mock", fmt.Sprintf("%s:%s", channel, payload), "write") {
      msg := fmt.Sprintf("   Write permissions for %s:%s denied for JWT %s",channel,payload,iamSession.Public)
      fmt.Println(msg)
      return
    }
    msg := fmt.Sprintf("Events of %s:%s wrote to by %s",channel,payload,iamSession.Public)
    fmt.Println(fmt.Sprintf("   %s",msg))
}
