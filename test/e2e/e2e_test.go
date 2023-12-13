package e2e_test

import (
	"context"
	"os"
	"testing"

	"github.com/otakakot/ninshow/pkg/api"
)

func TestE2E(t *testing.T) {
	t.Parallel()

	endpoint := os.Getenv("ENDPOINT")
	if endpoint == "" {
		endpoint = "http://localhost:5555"
	}

	cli, err := api.NewClient(endpoint, nil)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("health", func(t *testing.T) {
		t.Parallel()

		res, err := cli.Health(context.Background())
		if err != nil {
			t.Error(err)
		}

		switch res.(type) {
		case *api.HealthOK:
			return
		default:
			t.Errorf("unexpected type: %T", res)
		}
	})
}
