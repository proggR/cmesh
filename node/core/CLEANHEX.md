# Operation SCREAMING CLEANHEX
## SCREAMING
- Project hierarchy should scream domain level needs in your face at first glance

## CLEAN
- All dependencies flow inward (FWs/Drivers->Interface Adapters->App Business Rules->Enterprise Business Rules)
- domain/entity (enterprise business rules) level constructs are more abstractly defined, interface adapters like controllers/presenters/gateways more concretely defined
- [IF]Controller-> ([UC]Input Port <- ([UC]Interactor)) -> [IF]Presenter -> [UC]Output Port

## HEX
- Ports defines component->component I/O and mappings
- Adapters defines component->external I/O and mappings

## SCREAMING CLEANHEX
- project hierarchy will aim to remain clean, making clear on first glance exactly what services are required, and which elements are central to domain needs, opting to split up files if the added filenames within each package/layer add to the clarity at a glance
- core defines entities, primarily those of the router, registrar, and the IF for the IAM (routing + identity = core domain) given at its core a node is an authenticated/authenticating router
- services define the use cases the domain is responsible for authenticating against and routing to/between
- providers define the interface adapters services leverage to transform business logic into implementation details against the service provider (for now in-memory mock providers) and back again
- ports will define necessary interfaces between services, and particularly between services and the router defined in the core (<- critical to nail this down)
- adapters will define necessary boilerplate and/or complex mutations of data provided to/received from providers

### Hierarchy/Package Restructuring
/core
-> abstracts.go, iam.go, router.go, registrar.go
/services
-> ports.go, dispatcher.go, assembly.go, consensus.go, events.go, iam.go, registrar.go, router.go, state.go
/providers
-> ports.go, adapters.go, assembly/mock/provider.go, consensus/mock/provider.go, events/mock/provider.go, iam/mock/provider.go, registrar/mock/provider.go, state/mock/provider.go
/miners
-> events.go (more to follow, but events are central to the machine so this miner should be mocked first)

#### Entities
##### IAM/JWT
= every action requires access to/validation of the current state of the JWT
- ValidatePermissions() should incorporate the roughed out Permissions/ACL structs using FQMNs to bind DID->Permissions->FQMN

##### Registrar
- validates domain hasn't already been taken and otherwise maps it
- resolves domain if mapped, should return empty if unmapped (until request/response interfaces/structs defined to switch everything to)
- will need to be accessible by the router in order to resolve named services nested within the FQMN to return the resulting processed FQMN
- in a more complete version, authority and signatures will be verified (if claiming a domain for a contract/function, the DID should already have authority for that service endpoint)

##### Router
= core service of the node (the node _is_ the router), and embedded into/accessible by every service out of necessity (this is where port abstractions could come in handy)
- should return a formatted route struct that is then passed to a service layer dispatcher that invokes the necessary service/action
- services themselves may/should be able to make FQMN requests of the router mid runtime. in future versions these added requests/results will be compiled into a zk hash to help guarantee computation, but for now simply the ability to invoke FQMN resolution from each service should suffice for the toy
