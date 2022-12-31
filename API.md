# fundock API

`fundock` exposes an API which can be called for remote management and execution of functions. The API endpoints needs authentication, using API Tokens generated in the `API Tokens` section of the app. The generated token needs to be used in the `X-Api-Key` header.

## Endpoints

### 1. Execute function

|         |                                 |
| ------- | ------------------------------- |
| Request | `POST /api/exec/:function_name` |
| Input   | `any`                           |
| Output  | `any`                           |

Example:

```bash
curl -H "X-Api-Key: api-key-here" --data '{"name": "foo", "organisation": "bar"}' http://localhost:8080/api/exec/format-json
```
