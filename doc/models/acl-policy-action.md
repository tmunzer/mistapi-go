
# Acl Policy Action

Action applied to traffic that matches a destination ACL tag

## Structure

`AclPolicyAction`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Action` | [`*models.AllowDenyEnum`](../../doc/models/allow-deny-enum.md) | Optional | Policy action value that either allows or denies matching traffic. enum: `allow`, `deny` |
| `DstTag` | `string` | Required | Destination ACL tag matched by this policy action |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    aclPolicyAction := models.AclPolicyAction{
        Action:               models.ToPointer(models.AllowDenyEnum_ALLOW),
        DstTag:               "corp",
    }

}
```

