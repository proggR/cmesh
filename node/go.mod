module node

go 1.13

//replace basic-ci.skunk.services/vendor/cmesh/provider => ./vendor/cmesh/provider

//replace cmesh v0.0.0 => ./vendor/cmesh

//replace basic-ci.skunk.services/iam => ./iam

//require basic-ci.skunk.services/iam v0.0.0-00010101000000-000000000000 // indirect

require (
	github.com/hpcloud/tail v1.0.0
	golang.org/x/sys v0.0.0-20220708085239-5a0f0661e09d // indirect
	gopkg.in/fsnotify.v1 v1.4.7 // indirect
	gopkg.in/tomb.v1 v1.0.0-20141024135613-dd632973f1e7 // indirect
)
