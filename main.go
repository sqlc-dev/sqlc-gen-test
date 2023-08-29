package main

import (
	"context"
	"encoding/json"
	"os"

	"buf.build/gen/go/sqlc/sqlc/protocolbuffers/go/protos/plugin"
	"github.com/sqlc-dev/sqlc-go/codegen"
)

func generate(_ context.Context, _ *plugin.CodeGenRequest) (*plugin.CodeGenResponse, error) {
	env := os.Environ()
	blob, err := json.Marshal(env)
	if err != nil {
		return nil, err
	}
	return &plugin.CodeGenResponse{
		Files: []*plugin.File{
			{
				Name:     "env.json",
				Contents: append(blob, '\n'),
			},
		},
	}, nil
}

func main() {
	codegen.Run(generate)
}
