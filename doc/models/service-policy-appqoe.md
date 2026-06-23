
# Service Policy Appqoe

SRX application QoE settings for a service policy

## Structure

`ServicePolicyAppqoe`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enabled` | `*bool` | Optional | Whether application QoE is enabled for the service policy<br><br>**Default**: `false` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    servicePolicyAppqoe := models.ServicePolicyAppqoe{
        Enabled:              models.ToPointer(false),
    }

}
```

