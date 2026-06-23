
# Idp Profile Overwrite

Override rule that changes the IDP action for matching signatures

## Structure

`IdpProfileOverwrite`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Action` | [`*models.IdpProfileActionEnum`](../../doc/models/idp-profile-action-enum.md) | Optional | IDP profile action. enum: `alert`, `close`, `drop`. `alert` is the default, `drop` silently drops packets, and `close` asks the client/server to close the connection<br><br>**Default**: `"alert"` |
| `Matching` | [`*models.IdpProfileMatching`](../../doc/models/idp-profile-matching.md) | Optional | Criteria that select IDP signatures for an overwrite rule |
| `Name` | `*string` | Optional | Display name for this IDP profile overwrite rule |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    idpProfileOverwrite := models.IdpProfileOverwrite{
        Action:               models.ToPointer(models.IdpProfileActionEnum_ALERT),
        Matching:             models.ToPointer(models.IdpProfileMatching{
            AttackName:           []string{
                "attack_name9",
            },
            DstSubnet:            []string{
                "dst_subnet5",
            },
            Severity:             []models.IdpProfileMatchingSeverityValueEnum{
                models.IdpProfileMatchingSeverityValueEnum_CRITICAL,
                models.IdpProfileMatchingSeverityValueEnum_INFO,
            },
        }),
        Name:                 models.ToPointer("name8"),
    }

}
```

