package registrar
import(
  "fmt"
  "node/iam"
)

type Registrar struct{
  IAM iam.IAM
  domainMappings map[string]string
}

type RegistrarIF interface{
  Register(iam.JWT,string,string)
  Resolve(iam.JWT,string)
}


func (r *Registrar) Construct() Registrar{
  r.domainMappings = map[string]string{}
  return *r
}

func (r *Registrar) Register(jwt iam.JWT, domain string, fqmn string) string{
  fmt.Println(fmt.Sprintf("    Registering Domain: %s to service endpoint %s\n",domain,fqmn))
  if r.domainMappings[domain] != "" {
    fmt.Println(fmt.Sprintf("     Domain(%s) Already Registered to service endpoint (%s)\n     Try another domain name.\n",domain, r.domainMappings[domain]))
    return ""
  }
  r.domainMappings[domain] = fqmn
  return fmt.Sprintf("%s=>%s",domain,fqmn)
}


func (r *Registrar) Resolve(jwt iam.JWT, domain string) string{
  fmt.Println(fmt.Sprintf("    Resolving Domain: %s\n",domain))
  fqmn := "DOMAIN UNREGISTERED"
  if(r.domainMappings[domain] != ""){
    fqmn = r.domainMappings[domain]
  }
  return fmt.Sprintf(fqmn)
}
