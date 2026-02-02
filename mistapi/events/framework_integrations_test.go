// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package events

import (
	"context"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
	"time"
)

// runInSubmodule runs a Go command inside the nested framework_integrations module.
// It returns the command output or fails/skips the test on error.
func runInSubmodule(
	t *testing.T,
	timeout time.Duration,
	args ...string) string {
	t.Helper()
	testDir := filepath.Join("..", "framework_integrations")
	modFile := filepath.Join(testDir, "go.mod")
	if _, err := os.Stat(modFile); err != nil {
		t.Skipf("Skipping: %s not found (%v)", modFile, err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	cmd := exec.CommandContext(ctx, "go", args...)
	cmd.Dir = testDir
	out, err := cmd.CombinedOutput()

	if ctx.Err() == context.DeadlineExceeded {
		t.Fatalf("Timeout running: go %v\nPartial output:\n%s", args, out)
	}
	if err != nil {
		t.Fatalf("Failed: go %v\nError: %v\nOutput:\n%s", args, err, out)
	}
	return string(out)
}

// TestFrameworkIntegrationsTests ensures the framework_integrations submodule is tidy, builds, and passes its tests.
// It sequentially runs: go mod tidy, go get -t, go build, and go test.
func TestFrameworkIntegrationsTests(t *testing.T) {
	const stepTimeout = 60 * time.Second

	t.Log("Running go mod tidy...")
	outTidy := runInSubmodule(t, stepTimeout, "mod", "tidy")
	t.Logf("go mod tidy output:\n%s", outTidy)

	t.Log("Fetching test dependencies with go get -t...")
	outGetTest := runInSubmodule(t, stepTimeout, "get", "-t", "./...")
	t.Logf("go get -t output:\n%s", outGetTest)

	t.Log("Building all packages...")
	outBuild := runInSubmodule(t, stepTimeout, "build", "./...")
	t.Logf("go build output:\n%s", outBuild)

	t.Log("Running tests...")
	outTest := runInSubmodule(t, stepTimeout, "test", "./...")
	t.Logf("go test output:\n%s", outTest)
}
