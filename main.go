package main

import (
	"context"
	"encoding/json"
	"os"

	"buf.build/gen/go/sqlc/sqlc/protocolbuffers/go/protos/plugin"
	"github.com/sqlc-dev/sqlc-go/codegen"
)

type Output struct {
	Env []string `json:"env"`
}

func generate(_ context.Context, _ *plugin.CodeGenRequest) (*plugin.CodeGenResponse, error) {
	blob, err := json.MarshalIndent(Output{
		Env: os.Environ(),
	}, "", "  ")
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
