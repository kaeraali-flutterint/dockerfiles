package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/docker"
	"github.com/stretchr/testify/assert"
)

func TestDockerHelloWorldExample(t *testing.T) {
	tag := "flutterint/helloworld"
	buildOptions := &docker.BuildOptions{
		Tags: []string{tag},
	}

	docker.Build(t, "../ghcr.io/flutterint/helloworld", buildOptions)

	opts := &docker.RunOptions{Command: []string{"cat", "/test.txt"}}
	output := docker.Run(t, tag, opts)
	assert.Equal(t, "Hello, World!", output)
}
