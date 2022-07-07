package services


type Condition struct {

}

type UserServiceACL struct {
  service string
  authedForDID string
  authedByDID string
  ACL Permissions
}

type Permissions struct {
  read bool
  write bool
  exec bool
  deploy bool
  proxy bool
  replace bool
  conditions []Condition
}
