
# Routing Policy Term Action

When used as import policy

*This model accepts additional fields of type interface{}.*

## Structure

`RoutingPolicyTermAction`

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
| `LocalPreference` | `*string` | Optional | Optional, for an import policy, local_preference can be changed |
| `PrependAsPath` | `[]string` | Optional | When used as export policy, optional. By default, the local AS will be prepended, to change it |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

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
  ],
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

