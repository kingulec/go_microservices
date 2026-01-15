## Error Handling
| HTTP Status | Error Type            | Message Example                             | Description                                                      |
| ----------- | --------------------- | ------------------------------------------- | ---------------------------------------------------------------- |
| 400         | Invalid JSON          | `Bad request`                               | Request body is not a valid JSON or cannot be decoded            |
| 400         | Invalid Payload       | `Invalid payload: Missing required fields`  | Required JSON fields are missing or payload structure is invalid |
| 400         | Invalid Test Data     | `Error in tests: No tests provided`         | `tests` array is empty                                           |
| 400         | Invalid Test Data     | `Error in tests: Test name cannot be empty` | Test case contains an empty `name` field                         |
| 405         | Method Not Allowed    | `Method not allowed`                        | HTTP method other than `POST` was used                           |
| 500         | Internal Server Error | `Failed to encode response`                 | Server failed to encode JSON response                            |


