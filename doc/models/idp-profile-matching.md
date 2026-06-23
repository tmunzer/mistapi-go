
# Idp Profile Matching

Criteria that select IDP signatures for an overwrite rule

## Structure

`IdpProfileMatching`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AttackName` | `[]string` | Optional | IDP attack signature names matched by an overwrite rule |
| `DstSubnet` | `[]string` | Optional | Destination CIDR subnets matched by an IDP profile overwrite |
| `Severity` | [`[]models.IdpProfileMatchingSeverityValueEnum`](../../doc/models/idp-profile-matching-severity-value-enum.md) | Optional | Severity levels matched by an IDP profile overwrite |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    idpProfileMatching := models.IdpProfileMatching{
        AttackName:           []string{
            "attack_name3",
            "attack_name4",
        },
        DstSubnet:            []string{
            "dst_subnet5",
            "dst_subnet4",
        },
        Severity:             []models.IdpProfileMatchingSeverityValueEnum{
            models.IdpProfileMatchingSeverityValueEnum_CRITICAL,
        },
    }

}
```

