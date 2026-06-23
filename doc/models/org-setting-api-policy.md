
# Org Setting Api Policy

Organization API response policy for hiding secrets and passwords

## Structure

`OrgSettingApiPolicy`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `NoReveal` | `*bool` | Optional | By default, API hides password/secrets when the user doesn't have write access<br><br>* `true`: API will hide passwords/secrets for all users<br>* `false`: API will hide passwords/secrets for read-only users<br><br>**Default**: `false` |
| `SrcIps` | `[]string` | Optional | Optional list of IP addresses or CIDR subnets from which org API access is allowed. At most 10 entries. The source IP of the request making this update must be within one of the specified subnets.<br><br>**Constraints**: *Maximum Items*: `10` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    orgSettingApiPolicy := models.OrgSettingApiPolicy{
        NoReveal:             models.ToPointer(false),
        SrcIps:               []string{
            "63.3.56.0/24",
            "63.3.55.4",
        },
    }

}
```

