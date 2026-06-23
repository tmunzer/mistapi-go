
# Sw Routing Policy Term Action

Actions applied to routes matched by a switch routing policy term

## Structure

`SwRoutingPolicyTermAction`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Accept` | `*bool` | Optional | Whether to accept routes that match this term |
| `Community` | `[]string` | Optional | When used as export policy, optional |
| `LocalPreference` | [`*models.RoutingPolicyLocalPreference`](../../doc/models/containers/routing-policy-local-preference.md) | Optional | Optional, for an import policy, local_preference can be changed, value in range 1-4294967294. Can be a Variable (e.g. `{{bgp_as}}`) |
| `PrependAsPath` | `[]string` | Optional | When used as export policy, optional. By default, the local AS will be prepended, to change it. Can be a Variable (e.g. `{{as_path}}`) |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    swRoutingPolicyTermAction := models.SwRoutingPolicyTermAction{
        Accept:               models.ToPointer(false),
        Community:            []string{
            "community0",
            "community1",
            "community2",
        },
        LocalPreference:      models.ToPointer(models.RoutingPolicyLocalPreferenceContainer.FromString("String5")),
        PrependAsPath:        []string{
            "prepend_as_path1",
            "prepend_as_path2",
            "prepend_as_path3",
        },
    }

}
```

