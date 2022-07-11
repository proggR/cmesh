## State/Event Structure Noodlings


### Potential Event Channels
- events.requests
- events.responses.txId
- events.mined.txId
- events.confirmed.txId

### Potential State Hierarchy
/events
  events.requests.log, events.responses.log, events.mined.log, events.confirmed.log

/consensus
  /txId/userId/<responseString>.confirmation (compare filenames if possible for faster performance)

/chain
  /blocks
  /state
    /<contract address>/state.json
  /assembly
    /<contract address>/wasmBin.js

/repositories
  /DID/repoName
