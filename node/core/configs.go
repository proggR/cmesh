package core

var EventLogs map[string]string = map[string]string{
  "requests":"events/events.requests.log",
  "lastRequest":"events/events.request.last",
  "responses":"events/events.responses.log",
  "mined":"events/events.mined.log",
  "confirmed":"events/events.confirmed.log",
}

var ConsensusLogs map[string]string = map[string]string{
  "txClaim":"consensus/%d/%s/%s.claim", // Tx ID, DID, ResponseString
  "txConfirmation":"consensus/%d/%s/%s.confirmation", // Tx ID, DID, ResponseString
}

var CacheLogs map[string]string = map[string]string{
  "meta":"cache/%s.meta.json", // Hashed Request Object
  "data":"cache/%s.serialized", // Hashed Request Object
}

var ChainLogs map[string]string = map[string]string{
  "blocks":"chain/blocks/%d.json", // Block ID
  "state":"chain/state/%s/state.json", // Contract Address
  "assembly":"chain/assembly/%s/wasmBin.js", // Contract Address
}

var RepoLogs map[string]string = map[string]string{
  "respository":"repositories/%s/%s", // DID, Reponame
}

var GraphLogs map[string]string = map[string]string{
  "external":"graphs/%s/%s", // DID, Reponame
}

// /graphs
//   /internal
//     /requests
//     /responses
//     /mined
//     /confirmed
//   /external
//     /<DID>/<repoName>/
