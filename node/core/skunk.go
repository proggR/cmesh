package core

type CMeshMeta interface{

}

type CMeshMetaSeed struct {
  ID string
}

type CMeshMetaRepository struct {
  ID []string
  Object []WrappedCMeshObject
}

type CMeshObjectReposistoryRead interface {
  Fetch(string) WrappedCMeshObject
  Cached(string) CMeshObject
}

type CMeshObjectReposistoryWrite interface {
  Store(string, CMeshObject) bool
  Cache(string, CMeshObject) bool
  Delegate(string,string) bool
  Execute(string) bool
  Expire(string) bool
}

type WrappedCMeshObject interface {
  CMeshWrapper
  CMeshObject
}

type CMeshWrapper interface {
  Meta() CMeshMeta
}

type CMeshObject interface {
  //@TODO: implement Object ID system and enforce ID() method for nearly everything
    // Build() CMeshObject
}

type GSeed struct {
  ID string
}

type GraphIF interface {
  ID() string
  Node(string) NodeIF
  Edge(string) EdgeIF
}

type NodeIF interface {
  ID() string
  Peer(string,NodeIF) EdgeIF
  Peers(string) []NodeIF
}

type EdgeIF interface {
  ID() string
  FetchAssociations(string,NodeIF)[]NodeIF
  FetchSubgraphs(string)[]EdgeIF
}

type GraphSeed struct {
  GSeed
  nodes map[string] NodeIF
  edges map[string] EdgeIF
}

type NodeSeed struct {
  GSeed
}

type EdgeSeed struct {
  GSeed
  Selector string
  NodeIn NodeIF
  NodeOut NodeIF
}
