package main

import (
	"fmt"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/require"
)

const cliPkg = "github.com/danger/golang/cmd/danger-go"

func execute(args ...string) (string, error) {
	cmdArgs := append([]string{"run", cliPkg}, args...)
	cmd := exec.Command("go", cmdArgs...)
	res, err := cmd.CombinedOutput()
	return string(res), err
}

func TestShowsUsage(t *testing.T) {
	cc := []struct {
		name string
		args []string
	}{
		{name: "no args", args: nil},
		{name: "-h arg", args: []string{"-h"}},
		{name: "--help arg", args: []string{"--help"}},
	}

	for _, c := range cc {
		t.Run(c.name, func(t *testing.T) {
			res, err := execute(c.args...)
			require.Nil(t, err)
			require.Equal(t, usage, res)
		})
	}
}

func TestShowsVersion(t *testing.T) {
	res, err := execute("version")
	require.Nil(t, err)
	require.Equal(t, fmt.Sprintf("danger-go %s\n", version), res)
}
