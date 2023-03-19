module github.com/the-singularity-labs/holocron/lib/wasm

go 1.20

replace github.com/the-singularity-labs/holocron/pkg/forge => ../../pkg/forge

require github.com/the-singularity-labs/holocron/pkg/forge v0.0.0-00010101000000-000000000000

require (
	github.com/nogoegst/balloon v1.0.0 // indirect
	github.com/skip2/go-qrcode v0.0.0-20200617195104-da1b6568686e // indirect
)
