package state
import "fmt"

func Provider() {
    fmt.Println("Distributed State Provider Loaded")
}

type StateService struct{

}

type StateProviderPathSchema struct {
  blocks string
  events string
  iam string
  scripts string
  state string
}

type StateProviderBlocksState struct {

}

type StateProviderEventsState struct {

}

type StateProviderIAMState struct {
  DIDAddress string
}

type StateProviderAddressScripts struct {
  Address string

}

type StateProviderAddressState struct {
  Address string
}

type Transaction struct {

}

type Block struct {
  Hash uint32
  ParentHash uint32
  Nonce uint
  Prev *Block
  Timestamp uint32
  Transactions []Transaction
  MinedByDID string
  BlockReward uint32
  Difficulty uint32
  TotalDifficulty uint32
  Size uint32
  Cycles uint32
  CyclesLimit uint32
  CyclesCost uint32
  CyclesBurnt uint32
  ExtraData string
}
