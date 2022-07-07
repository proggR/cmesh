package services

import (
  "fmt"
)

func provider() {
    fmt.Println("AssemblyScript Interpretor Provider Loaded")
}

type ABI struct {

}

type WASMScript struct {
  Script string
}
