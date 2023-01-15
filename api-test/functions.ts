import {
  assertEquals,
  assertExists,
} from "https://deno.land/std@0.103.0/testing/asserts.ts";

Deno.test("register and execute function", async () => {
  // Create function
  let url = new URL("/api/functions", "http://localhost:8080/");

  let requestBody = JSON.stringify({
    name: "hello",
    image: "docker.io/hello-world:latest",
    skip_logging: false,
  });

  let response = await fetch(url, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      "X-Api-Key": "foobar",
    },
    body: requestBody,
  });

  assertEquals(response.status, 200);
  assertExists(response.body);

  let responseBody = await response.json();
  assertEquals(responseBody.name, "hello");

  // Execute it
  url = new URL("/api/fn/hello/exec", "http://localhost:8080/");

  response = await fetch(url, {
    method: "POST",
    headers: {
      "X-Api-Key": "foobar",
    },
  });

  assertEquals(response.status, 200);
  assertExists(response.body);

  responseBody = await response.json();
  assertEquals(
    responseBody.output,
    `
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

`
  );

  // Fetch invocations
  url = new URL(
    "/api/invocations?page_size=10&page_number=1&fn=hello",
    "http://localhost:8080/"
  );

  response = await fetch(url, {
    method: "GET",
    headers: {
      "X-Api-Key": "foobar",
    },
  });

  assertEquals(response.status, 200);
  assertExists(response.body);

  responseBody = await response.json();
  assertExists(responseBody.items);
  assertEquals(responseBody.items[0].function_name, "hello");
});
