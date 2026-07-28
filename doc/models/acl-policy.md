
# Acl Policy

ACL Policy:

- for GBP-based policy, all src_tags and dst_tags have to be gbp-based
- for ACL-based policy, `network` is required in either the source or destination so that we know where to attach the policy to

## Structure

`AclPolicy`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Actions` | [`[]models.AclPolicyAction`](../../doc/models/acl-policy-action.md) | Optional | ACL Policy Actions:<br><br>- for GBP-based policy, all src_tags and dst_tags have to be gbp-based<br>- for ACL-based policy, `network` is required in either the source or destination so that we know where to attach the policy to |
| `Disabled` | `*bool` | Optional | Whether this ACL policy is disabled<br><br>**Default**: `false` |
| `Name` | `*string` | Optional | Display name of the ACL policy |
| `SrcTags` | `[]string` | Optional | ACL Policy Source Tags:<br><br>- for GBP-based policy, all src_tags and dst_tags have to be gbp-based<br>- for ACL-based policy, `network` is required in either the source or destination so that we know where to attach the policy to |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    aclPolicy := models.AclPolicy{
        Actions:              []models.AclPolicyAction{
            models.AclPolicyAction{
                Action:               models.ToPointer(models.AllowDenyEnum_ALLOW),
                DstTag:               "dst_tag0",
            },
            models.AclPolicyAction{
                Action:               models.ToPointer(models.AllowDenyEnum_ALLOW),
                DstTag:               "dst_tag0",
            },
        },
        Disabled:             models.ToPointer(false),
        Name:                 models.ToPointer("guest access"),
        SrcTags:              []string{
            "src_tags5",
            "src_tags4",
            "src_tags3",
        },
    }

}
```

