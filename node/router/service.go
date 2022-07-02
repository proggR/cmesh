package router


/**
* Proposed Op Code Prefixes
* IAM: "0Ix:"
* State: "0Sx:"
* Events: "0Ex:"
* Assembly: "0Ax:"
* Consensus: "0Cx:"
* Registrar: "0Rx:"
* Router: "0Px:" (P for proxy, since you're routing to another router... a strange notion I'll need to consider)
*
* Then add the resource, which could vary by service
* IAM: action
* State: contract:action(:args)?
* Assembly: version:contract:action(:args)?
* Events: channel:action(:args)?
* Router: FQMN
* Registrar: action:entityHash:authority:authoritySig | namedService
* Consensus: service:action #(? haven't researched far enough on raft for IF)

* ### FGMN Examples
* Blockahain Greeter Function With Args: 0Sx:0x03389f0e08b9f:hello_world (without args, should only read from state if function marked as pure, with some kind of cache?)
* Event Stream For Registar Named Service Resolution: 0Ex:helloWorldExample.mcom
* Blockahain Greeter Function With Args: 0Sx:0x03389f0e08b9f:hello_world:["Dexter"] (with args, assumes computation and invokes smart contract through Assembler call)
* Assembly Exec: 0Ax:~4.1.3:0x03389f0e08b9f::hello_world:["Dexter"] (args passed through a "box" (to define) to the script, which returns into the "box" to the state service
* (^ Note: Assembly should only ever be called via state service, and Assembly should only exist to process state changes requiring processing (define consensused cache strat))
*/
