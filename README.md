# Go build information utilities

These are some small build information utilities I wrote for tracking Go binary
version information. Rather than trying to assign a linear version number to a
binary, the tag names and version control commit hashes of all dependencies are
tracked. This information is then burnt into the binary at build time.

You use the shell script `gen`, included, to generate the version information
blob. It outputs arguments suitable for `go build`, so you can use it like `go
build ./... -ldflags "$($GOPATH/src/github.com/hlandau/buildinfo/gen ./...)"`.

## For upstream packagers

OS upstream packagers which want to specify their own information may do so:
simply set the `github.com/hlandau/buildinfo.RawBuildInfo` string variable at
build time, by passing `-ldflags -X
github.com/hlandau/buildinfo.RawBuildInfo=BUILDINFO...` to `go build`. This
expects base64-encoded data in a particular format.

If you want to specify an arbitrary freeform string instead, set `BuildInfo`
instead of `RawBuildInfo`. The string you specify will be used verbatim without
modification.

Example:

```
go install -ldflags '-X github.com/hlandau/buildinfo.BuildInfo="Packaged by Distro X"' \
  github.com/my/project 
```

### Previous location

This package was previously located at
`github.com/hlandau/degoutils/buildinfo`. Packagers setting
`github.com/hlandau/degoutils/buildinfo.BuildInfo` will need to change to
`github.com/hlandau/buildinfo.BuildInfo`.

