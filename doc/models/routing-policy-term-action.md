
# Routing Policy Term Action

when used as import policy

*This model accepts additional fields of type interface{}.*

## Structure

`RoutingPolicyTermAction`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Accept` | `*bool` | Optional | - |
| `AddCommunity` | `[]string` | Optional | - |
| `AddTargetVrfs` | `[]string` | Optional | for SSR, hub decides how VRF routes are leaked on spoke |
| `Aggregate` | `[]string` | Optional | route aggregation |
| `Community` | `[]string` | Optional | when used as export policy, optional |
| `ExcludeAsPath` | `[]string` | Optional | when used as export policy, optional. To exclude certain AS |
| `ExcludeCommunity` | `[]string` | Optional | - |
| `ExportCommunitites` | `[]string` | Optional | when used as export policy, optional |
| `LocalPreference` | `*string` | Optional | optional, for an import policy, local_preference can be changed |
| `PrependAsPath` | `[]string` | Optional | when used as export policy, optional. By default, the local AS will be prepended, to change it |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "aggregate": [
    "192.168.0.0/16",
    "10.234.30.0/18"
  ],
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
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

