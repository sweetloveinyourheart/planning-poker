//go:build tools || generate
// +build tools generate

//go:generate sh -c "go list -f '{{.ImportPath}}@{{.Module.Version}}' $(sed -n 's/.*_ \"\\(.*\\)\".*/\\1/p' <$GOFILE) | GOBIN=$(git rev-parse --show-toplevel)/.gobincache xargs -n 1 go install"

package tools

import (
	_ "github.com/bufbuild/buf/cmd/buf"
	_ "google.golang.org/grpc/cmd/protoc-gen-go-grpc"
	_ "google.golang.org/protobuf/cmd/protoc-gen-go"
)