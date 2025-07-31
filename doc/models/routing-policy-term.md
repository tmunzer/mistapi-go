
# Routing Policy Term

*This model accepts additional fields of type interface{}.*

## Structure

`RoutingPolicyTerm`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Actions` | [`*models.RoutingPolicyTermAction`](../../doc/models/routing-policy-term-action.md) | Optional | When used as import policy |
| `Matching` | [`*models.RoutingPolicyTermMatching`](../../doc/models/routing-policy-term-matching.md) | Optional | zero or more criteria/filter can be specified to match the term, all criteria have to be met |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "actions": {
    "accept": false,
    "add_community": [
      "add_community5",
      "add_community6",
      "add_community7"
    ],
    "add_target_vrfs": [
      "add_target_vrfs1"
    ],
    "community": [
      "community4",
      "community5"
    ],
    "exclude_as_path": [
      "exclude_as_path0",
      "exclude_as_path1",
      "exclude_as_path2"
    ],
    "exampleAdditionalProperty": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "matching": {
    "as_path": [
      "as_path2"
    ],
    "community": [
      "community4"
    ],
    "network": [
      "network7",
      "network8",
      "network9"
    ],
    "prefix": [
      "prefix5",
      "prefix6",
      "prefix7"
    ],
    "protocol": [
      "protocol5",
      "protocol6"
    ],
    "exampleAdditionalProperty": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

