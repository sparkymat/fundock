package docker_test

import (
	"context"
	"testing"

	"github.com/sparkymat/fundock/docker"
	"github.com/stretchr/testify/require"
)

func TestDockerRun(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name           string
		image          string
		input          string
		errorExpected  bool
		expectedOutput string
	}{
		{
			name:          "hello-world returns static output",
			image:         "docker.io/hello-world:latest",
			input:         "",
			errorExpected: false,
			expectedOutput: `
Hello from Docker!
This message shows that your installation appears to be working correctly.

To generate this message, Docker took the following steps:
 1. The Docker client contacted the Docker daemon.
 2. The Docker daemon pulled the "hello-world" image from the Docker Hub.
    (amd64)
 3. The Docker daemon created a new container from that image which runs the
    executable that produces the output you are currently reading.
 4. The Docker daemon streamed that output to the Docker client, which sent it
    to your terminal.

To try something more ambitious, you can run an Ubuntu container with:
 $ docker run -it ubuntu bash

Share images, automate workflows, and more with a free Docker ID:
 https://hub.docker.com/

For more examples and ideas, visit:
 https://docs.docker.com/get-started/

`,
		},
	}

	for _, testCase := range testCases { //nolint:paralleltest
		testCase := testCase
		t.Run(testCase.name, func(it *testing.T) {
			t.Parallel()

			svc, err := docker.New()
			require.NoError(t, err)

			output, err := svc.Run(context.Background(), testCase.image, testCase.input)
			if testCase.errorExpected {
				require.Error(it, err)
			} else {
				require.NoError(it, err)
				require.Equal(it, testCase.expectedOutput, output)
			}
		})
	}
}
