
# Routing Policy Term

## Structure

`RoutingPolicyTerm`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Action` | [`*models.RoutingPolicyTermAction`](../../doc/models/routing-policy-term-action.md) | Optional | when used as import policy |
| `Matching` | [`*models.RoutingPolicyTermMatching`](../../doc/models/routing-policy-term-matching.md) | Optional | zero or more criteria/filter can be specified to match the term, all criteria have to be met |

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
    "community": [
      "community8",
      "community9",
      "community0"
    ],
    "exclude_as_path": [
      "exclude_as_path4"
    ]
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
    ]
  }
}
```

