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
* Then add the resource string, which could vary by service
* IAM: action | FQMN
* State: contract:action(:args)?
* Assembly: version:contract:action(:args)?
* Events: channel:action(:args)?
* Router: FQMN
* Registrar: action:entityHash:argString:authority:authoritySig | namedService
* Consensus: service:action #(? haven't researched far enough on raft for IF)

* ### FGMN Examples
* Blockahain Greeter Function With Args: 0Sx:0x03389f0e08b9f:hello_world (without args, should only read from state if function marked as pure, with some kind of cache?)
* Registrar registration mapping helloWorld.mcom to contract: 0Rx:register:hash:"helloworld.mcom,0Sx:0x03389f0e08b9f",0x079f9849dac562,874958794857983475893475)
* Registrar registration mapping helloWorldExample.mcom to contract function: 0Rx:register:hash:"helloWorldExample.mcom,0Sx:0x03389f0e08b9f:hello_world",0x079f9849dac562,874958794857983475893475)
* Event Stream For Registar Named Service Registration: 0Ex:0Rx:Registered
* Event Stream For Transferred Event From Named Contract: 0Ex:0Rx:helloworld.mcom:Transferred
* Blockahain Greeter Function With Args: 0Sx:0x03389f0e08b9f:hello_world:["Dexter"] (with args, assumes computation and invokes smart contract through Assembler call)
* Same Greeter Function Via Named Service: 0Rx:helloWorldExample.mcom:["Dexter"] (same rules re: args apply)
* Assembly Exec: 0Ax:~4.1.3:0x03389f0e08b9f:hello_world:["Dexter"] (args passed through a "box" (to define) to the script, which returns into the "box" to the state service
* (^ Note: Assembly should only ever be called via state service, and as such can't be registered with a name, and Assembly should only exist to process state changes requiring processing (define consensused cache strat))
*/
