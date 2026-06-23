
# Service Policy Ewf Rule

Enhanced web filtering rule applied by a service policy

## Structure

`ServicePolicyEwfRule`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AlertOnly` | `*bool` | Optional | Whether matching enhanced web filtering traffic is logged without being blocked |
| `BlockMessage` | `*string` | Optional | Message returned when enhanced web filtering blocks a request |
| `Enabled` | `*bool` | Optional | Whether this enhanced web filtering rule is enabled<br><br>**Default**: `false` |
| `Profile` | [`*models.ServicePolicyEwfRuleProfileEnum`](../../doc/models/service-policy-ewf-rule-profile-enum.md) | Optional | enum: `critical`, `standard`, `strict`<br><br>**Default**: `"strict"` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    servicePolicyEwfRule := models.ServicePolicyEwfRule{
        AlertOnly:            models.ToPointer(false),
        BlockMessage:         models.ToPointer("Access to this URL Category has been blocked"),
        Enabled:              models.ToPointer(false),
        Profile:              models.ToPointer(models.ServicePolicyEwfRuleProfileEnum_STRICT),
    }

}
```

