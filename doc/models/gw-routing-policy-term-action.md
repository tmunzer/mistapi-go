
# Gw Routing Policy Term Action

Actions applied to routes matched by a gateway routing policy term

## Structure

`GwRoutingPolicyTermAction`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Accept` | `*bool` | Optional | Whether to accept routes that match this term. Precedence is `accept` > `next_term` > `next_policy`; routes are rejected if all three are false |
| `AddCommunity` | `[]string` | Optional | BGP communities added to routes matched by a gateway routing policy term |
| `AddTargetVrfs` | `[]string` | Optional | For SSR, target VRFs used by a hub when leaking VRF routes to spokes |
| `Community` | `[]string` | Optional | When used as export policy, optional |
| `ExcludeAsPath` | `[]string` | Optional | Optional when used as an export policy. AS path values to exclude |
| `ExcludeCommunity` | `[]string` | Optional | BGP communities excluded by a gateway routing policy term |
| `ExportCommunities` | `[]string` | Optional | Optional when used as an export policy. BGP communities that may be exported |
| `LocalPreference` | [`*models.RoutingPolicyLocalPreference`](../../doc/models/containers/routing-policy-local-preference.md) | Optional | Optional, for an import policy, local_preference can be changed, value in range 1-4294967294. Can be a Variable (e.g. `{{bgp_as}}`) |
| `NextPolicy` | `*bool` | Optional | When true, continue evaluating the next routing policy in the chain after this term matches; default is false<br><br>**Default**: `false` |
| `NextTerm` | `*bool` | Optional | When true, continue evaluating the next term in the same routing policy after this term matches; default is false<br><br>**Default**: `false` |
| `PrependAsPath` | `[]string` | Optional | When used as export policy, optional. By default, the local AS will be prepended, to change it. Can be a Variable (e.g. `{{as_path}}`) |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    gwRoutingPolicyTermAction := models.GwRoutingPolicyTermAction{
        Accept:               models.ToPointer(false),
        AddCommunity:         []string{
            "add_community9",
            "add_community0",
        },
        AddTargetVrfs:        []string{
            "add_target_vrfs9",
            "add_target_vrfs8",
            "add_target_vrfs7",
        },
        Community:            []string{
            "community8",
        },
        ExcludeAsPath:        []string{
            "exclude_as_path4",
            "exclude_as_path5",
        },
        NextPolicy:           models.ToPointer(false),
        NextTerm:             models.ToPointer(false),
    }

}
```

