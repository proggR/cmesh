# CMesh: Cumulus Mesh Transport Protocol

CMesh is for the moment a toy just to give me an excuse to play with a bunch of new, interesting tech all at once. If you want to blame someone for it, I blame NATS... its too pretty not to get the gears turning dreaming up a machine just to use it... this is that machine :\

![CMesh First Bad Model](assets/images/firstroughmodel.jpeg)

> Initial rough and very inaccurate imagined CMesh Node<->Client model

## The Dream Machine

The idea is still loosely formed, but the goal seems worth hacking toward: a privacy respecting, passwordless, identity requiring decentralized internet made up of modular and hotswappable service providers for each core component of the network (in build phase, service providers are: ID/Auth=IRMA;Event/Message Bus=NATS;Persistent State=HyperCore;Consensus=Raft;Smart Contracts=WASM).

## Modular & Decentralized Design, Familiar Monolithic Feeling UX

Seeing that NATS could be a component that helped underwrite a new kind of DLT and novel UX data flows (I think... probably not actually novel tbh), I started to think through/research additional components that would be necessary to stitch together the machine I was imagining. NATS serviced as the core means to immutably route events, but that does not a blockchain/decentralized smart contract platform make.

The components identified and providers/stacks chosen are the following:

### Event Sourcing Pipelines: NATS

![NATS](assets/images/NATS_logo.png)

NATS underwrites the entire machine's comms, from reads, to writes, to zk auth logging, to transaction bundling, and with some more thought on the necessary adapters likely even to interchain bridges, NATS is a critical component and inspiration to the entire machine. Its also very, very pretty.

### Identity Management and Auth: IRMA

![IRMA Flow](assets/images/irmaflow.png)

IRMA is a beautiful spec, and in this case not being dependent or built for blockchain makes it the perfect choice for... whatever this is :\ lol. All sessions are first authenticated through IRMA, and JWTs are passed as needed via relevant NATS subjects (need to research this a ton more to be sure I fully grok both systems... this portion is the most critical component. feel like gRPC will be necessary, with NATS subjects being used for zk logs, but need to research this architecture first/thoroughly).

This is the only part of the toy that "functions". In `node/main.go` there's an `iam_test` function that walks through the IRMA handshake process, and then spams it to confirm sessionless requests are rejected. Very much a toy given the private keys are shared between "client" and "server" (main and iam) and its using a hash function that's more convenient than secure. Also doesn't currently contain any of the attribute level functionality of IRMA systems.

Docs for existing toy version can be found in [iam/README](iam/README.md)

### Distributed Storage: HyperCore/Hyperbee/Hyperdrive

![Hypercore](assets/images/hypercore.png)

Requiring storage, but not wanting to bloat the node's logic with rolling some kind of custom storage mechanism, the HyperCore stack stands out as an ideal option for carving up necessary event sourced data and persisting it in a way that should respect privacy while improving in performance with scale instead of degrading (I believe).

### Smart Contract Runtime: WASMI

![WASM](assets/images/webassembly.jpeg)

EVM is very solid and well established tech, but opting for a more open standard leaning into WASM feels like the better call. Rust devs will be able to port existing smart contracts to the environment, and AssemblyScript will be the supported/native language to keep CMesh smart contracts both accessible to newcomers, and portable to/from other distributed systems if devs so choose.

### Consensus & Cluster Management: RAFT (potentially via Swarmkit)

![RAFT](assets/images/raft.gif)

RAFT conensus is an established, fault tolerant and fast consensus mechanism that will underwrite state requests requiring auth, as well as any state change event handling, which for the moment I assume will include any necessary IRMA session state (haven't dug deep enough to determine IRMA session needs or related persistence strat).

### ZKProofs: Undecided

Looking at the topology of the machine, it feels ideal to aim to architect as much of its innerworkings around ZKRollups/Cairo/Starknet-like ZK Proofs of Computation as possible, or at least to enable said featureset as part of the core functionality. This will be left for later given the added complexity and room for error it could introduce, but by the time the "toy" version of the above components are integrated a solution will be arrived at/model formed.

## CMTP: message protocol built for the distributed web

As HTTP did for the the web, CMTP allows a client to pass a request to a backend service operating on a distributed machine that is then able to be communicated to the other components of the machine with state being added to it in layers as each step of each opcode defined process entails, which then has the hash of the final message state streamed to create a complete audit trail and make the request's final state available for future processing/debugging/analytics/forensics.

One core aim of CMTP that makes it different is that its aim is to enable an auth'd by default environment flexible enough to protect privacy and enable multiple online identities to be managed from a core root identity. CMTP messages require a valid JWT, and without will be rejected with a status code instructing clients to prompt the identity creation/sync process/dialog. No valid IAM session, no network access for you.

Note: Currently this spec below is just capturing HTTP as pseudocode. CMTP development will likely not see much new work until the other components are more integrated, but initial HTTP mimicking Structs are loosely defined (though commented out) in `vendor/cmtp/cmtp.go` so any developments in the protocol will take place there.

<pre>

message_compiler(message){
  {
    headers:{
        status: message.status(),
        version: 'CMTP/0',
        transferred: message.bytesize(),
        referrerPolicy: message.referrerPolicy(),
        request: {
          accept: message.accept(),
          acceptEncoding: message.acceptEncoding(),
          acceptLanguage: message.acceptLanguage(),
          connection: message.connection(),
          host: message.host(),
          ifModSince: message.ifModSince(),
          ifNoneMatch: message.ifNoneMatch(),
          referrer: message.refferer(),
          secFetchDest: message.secFetchDest(),
          secFetchMode: message.secFetchMode(),
          secFetchUser: message.secFetchUser(),
          userAgent: message.userAgent(),
        },
        response: {
          acceptRanges: message.response.acceptRanges(),
          cacheControl: message.response.cacheControl(),
          contentEncoding: message.response.contentEncoding(),
          contentLength: message.response.contentLength()|message.response.bytesize(),
          contentSecurityPolicy: message.response.contentSecurityPolicy(),
          contentType: message.response.contentType(),
          date: message.response.date(),
          etag: message.response.etag(),
          expires: message.response.expires(),
          lastModified: message.response.lastModified(),
          strictTransportSecurity: message.response.strictTransportSecurity(),
          vary: message.response.vary(),
          (x(\w)) $1.each as $attr message.response.reflect($attr),
        }
    },
  }

}

</pre>


## Project Hierarchy

`/`

`->assets/ (README/shared assets)`

`->node/ (node/mining app)`

`->->node/assembly(/providers)? (WASMI component & providers)`

`->->node/consensus(/providers)? (consensus cluster/swarm component & providers)`

`->->node/events(/providers)? (event sourcing component & providers)`

`->->node/iam(/providers)? (SSI/DID component & providers)`

`->->node/state(/providers)? (persistent state component & providers)`


`->client/ (client app)`

`->->client/events(/providers)? (event sourcing component & providers)`

`->->client/iam(/providers)? (SSI/DID component & providers)`

`->->client/state(/providers)? (persistent state component & providers)
