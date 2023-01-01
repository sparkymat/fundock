# fundock API

`fundock` exposes an API which can be called for remote management and execution of functions. The API endpoints needs authentication, using API Tokens generated in the `API Tokens` section of the app. The generated token needs to be used in the `X-Api-Key` header.

## Endpoints

### 1. Execute function

`POST /api/exec/:function_name`

**Input**: `<any>`

**Output**: Invocation details

```json
{
  "ID": "8bc3ccaf-aefe-4afb-97a8-fefb7cfbf8d8",
  "StartedTime": "2023-01-02 07:41:43.458696 +0000 +0000",
  "EndedTime": "2023-01-02 07:41:44.481944 +0000 +0000",
  "Output": "response-from-function"
}
```

Curl example:

```bash
curl -X POST -H "X-Api-Key: api-key-here" --data '{"name": "foo", "organisation": "bar"}' http://localhost:8080/api/exec/format-json
```

### 2. Start function asynchronously

`POST /api/start/:function_name`

**Input**: `<any>`

**Output**: Invocation details

```json
{
  "ID": "8bc3ccaf-aefe-4afb-97a8-fefb7cfbf8d8",
  "StartedTime": "2023-01-02 07:41:43.458696 +0000 +0000",
  "EndedTime": "2023-01-02 07:41:44.481944 +0000 +0000",
  "Output": "" // will be empty since it was not actually executed
}
```

Curl example:

```bash
curl -X POST -H "X-Api-Key: api-key-here" --data '{"name": "foo", "organisation": "bar"}' http://localhost:8080/api/start/format-json
```

### 3. Get invocation details

`GET /api/invocations/:id`

**Input**:

- `id` - The id of the invocation

**Output**: Invocation details

```json
{
  "ID": "8bc3ccaf-aefe-4afb-97a8-fefb7cfbf8d8",
  "StartedTime": "2023-01-02 07:41:43.458696 +0000 +0000",
  "EndedTime": "2023-01-02 07:41:44.481944 +0000 +0000",
  "Output": "this-can-be-empty-too"
}
```

Curl example:

```bash
curl -H "X-Api-Key: api-key-here" http://localhost:8080/api/invocations/8bc3ccaf-aefe-4afb-97a8-fefb7cfbf8d8
```
