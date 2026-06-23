
# Service Policy Syslog

Syslog logging settings for a service policy

## Structure

`ServicePolicySyslog`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enabled` | `*bool` | Optional | Whether syslog logging is enabled for the service policy<br><br>**Default**: `false` |
| `ServerNames` | `[]string` | Optional | Names of syslog servers that receive logs for this service policy |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    servicePolicySyslog := models.ServicePolicySyslog{
        Enabled:              models.ToPointer(false),
        ServerNames:          []string{
            "dc_syslog_server",
        },
    }

}
```

