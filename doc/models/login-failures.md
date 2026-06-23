
# Login Failures

Failed login attempt summary with source IPs and user agents

## Structure

`LoginFailures`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Email` | `*string` | Optional | User email address for the failed login attempts |
| `LastFailureAt` | `*int` | Optional | Time of the most recent failed login attempt |
| `NumAttempts` | `*int` | Optional | Number of failed login attempts |
| `SrcIps` | `[]string` | Optional | List of up to 32 unique source IP addresses, ordered with the most recent first |
| `UserAgents` | `[]string` | Optional | List of up to 32 unique User-Agent strings, ordered with the most recent first |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    loginFailures := models.LoginFailures{
        Email:                models.ToPointer("admin@test.com"),
        LastFailureAt:        models.ToPointer(1509161968),
        NumAttempts:          models.ToPointer(1),
        SrcIps:               []string{
            "192.168.1.39",
            "192.168.1.38",
            "192.168.1.37",
        },
        UserAgents:           []string{
            "Test UA 39",
            "Test UA 38",
            "Test UA 37",
        },
    }

}
```

