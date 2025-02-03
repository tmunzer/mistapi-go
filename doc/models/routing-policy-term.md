
# Routing Policy Term

*This model accepts additional fields of type interface{}.*

## Structure

`RoutingPolicyTerm`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Action` | [`*models.RoutingPolicyTermAction`](../../doc/models/routing-policy-term-action.md) | Optional | When used as import policy |
| `Matching` | [`*models.RoutingPolicyTermMatching`](../../doc/models/routing-policy-term-matching.md) | Optional | zero or more criteria/filter can be specified to match the term, all criteria have to be met |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "action": {
    "accept": false,
    "add_community": [
      "add_community9"
    ],
    "add_target_vrfs": [
      "add_target_vrfs9",
      "add_target_vrfs8"
    ],
    "aggregate": [
      "aggregate1",
      "aggregate0",
      "aggregate9"
    ],
    "community": [
      "community8",
      "community9",
      "community0"
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

