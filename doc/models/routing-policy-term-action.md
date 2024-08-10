
# Routing Policy Term Action

when used as import policy

## Structure

`RoutingPolicyTermAction`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Accept` | `*bool` | Optional | - |
| `AddCommunity` | `[]string` | Optional | - |
| `AddTargetVrfs` | `[]string` | Optional | for SSR, hub decides how VRF routes are leaked on spoke |
| `Community` | `[]string` | Optional | when used as export policy, optional |
| `ExcludeAsPath` | `[]string` | Optional | when used as export policy, optional. To exclude certain AS |
| `ExcludeCommunity` | `[]string` | Optional | - |
| `ExportCommunitites` | `[]string` | Optional | when used as export policy, optional |
| `LocalPreference` | `*string` | Optional | optional, for an import policy, local_preference can be changed |
| `PrependAsPath` | `[]string` | Optional | when used as export policy, optional. By default, the local AS will be prepended, to change it |

## Example (as JSON)

```json
{
  "accept": false,
  "add_community": [
    "add_community5",
    "add_community6",
    "add_community7"
  ],
  "add_target_vrfs": [
    "add_target_vrfs9"
  ],
  "community": [
    "community0",
    "community9",
    "community8"
  ],
  "exclude_as_path": [
    "exclude_as_path0",
    "exclude_as_path9",
    "exclude_as_path8"
  ]
}
```

