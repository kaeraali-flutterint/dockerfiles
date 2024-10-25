package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/docker"
	"github.com/stretchr/testify/assert"
)

func TestDockerHelloWorldExample(t *testing.T) {
	tag := "kaeraali-flutterint/helloworld"
	buildOptions := &docker.BuildOptions{
		Tags: []string{tag},
	}

	docker.Build(t, "../", buildOptions)

	opts := &docker.RunOptions{Command: []string{"cat", "/test.txt"}}
	output := docker.Run(t, tag, opts)
	assert.Equal(t, "Hello World!", output)
}
