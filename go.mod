module github.com/tmunzer/mistapi-go

go 1.24.0

toolchain go1.24.7

require (
	github.com/apimatic/go-core-runtime v0.0.29
	github.com/google/uuid v1.6.0
)

retract [v0.1.0, v0.1.20] // versions based on older sdk generation

retract v0.3.9 // import issue
