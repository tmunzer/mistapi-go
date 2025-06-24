
# Login Failures

Failed login attempts

*This model accepts additional fields of type interface{}.*

## Structure

`LoginFailures`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Email` | `*string` | Optional | Email address of the user |
| `LastFailureAt` | `*int` | Optional | Last failure time |
| `NumAttempts` | `*int` | Optional | Number of failed login attempts |
| `ScrIps` | `[]string` | Optional | List of up to 32 unique source IP addresses, ordered with the most recent first |
| `UserAgents` | `[]string` | Optional | List of up to 32 unique User-Agent strings, ordered with the most recent first |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "email": "admin@test.com",
  "last_failure_at": 1509161968,
  "num_attempts": 1,
  "scr_ips": [
    "192.168.1.39",
    "192.168.1.38",
    "192.168.1.37"
  ],
  "user_agents": [
    "Test UA 39",
    "Test UA 38",
    "Test UA 37"
  ],
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

