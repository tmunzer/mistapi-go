
# Gw Routing Policy Term Action

When used as import policy

## Structure

`GwRoutingPolicyTermAction`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Accept` | `*bool` | Optional | - |
| `AddCommunity` | `[]string` | Optional | - |
| `AddTargetVrfs` | `[]string` | Optional | For SSR, hub decides how VRF routes are leaked on spoke |
| `Community` | `[]string` | Optional | When used as export policy, optional |
| `ExcludeAsPath` | `[]string` | Optional | When used as export policy, optional. To exclude certain AS |
| `ExcludeCommunity` | `[]string` | Optional | - |
| `ExportCommunities` | `[]string` | Optional | When used as export policy, optional |
| `LocalPreference` | [`*models.RoutingPolicyLocalPreference`](../../doc/models/containers/routing-policy-local-preference.md) | Optional | Optional, for an import policy, local_preference can be changed, value in range 1-4294967294. Can be a Variable (e.g. `{{bgp_as}}`) |
| `PrependAsPath` | `[]string` | Optional | When used as export policy, optional. By default, the local AS will be prepended, to change it. Can be a Variable (e.g. `{{as_path}}`) |

## Example (as JSON)

```json
{
  "accept": false,
  "add_community": [
    "add_community3"
  ],
  "add_target_vrfs": [
    "add_target_vrfs9",
    "add_target_vrfs8",
    "add_target_vrfs7"
  ],
  "community": [
    "community2",
    "community3",
    "community4"
  ],
  "exclude_as_path": [
    "exclude_as_path8"
  ]
}
```

