
# Sw Routing Policy Term

## Structure

`SwRoutingPolicyTerm`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Actions` | [`*models.SwRoutingPolicyTermAction`](../../doc/models/sw-routing-policy-term-action.md) | Optional | When used as import policy |
| `Matching` | [`*models.SwRoutingPolicyTermMatching`](../../doc/models/sw-routing-policy-term-matching.md) | Optional | zero or more criteria/filter can be specified to match the term, all criteria have to be met |
| `Name` | `string` | Required | - |

## Example (as JSON)

```json
{
  "actions": {
    "accept": false,
    "community": [
      "community4",
      "community5"
    ],
    "local_preference": "String5",
    "prepend_as_path": [
      "prepend_as_path9",
      "prepend_as_path8",
      "prepend_as_path7"
    ]
  },
  "matching": {
    "as_path": [
      "String3"
    ],
    "community": [
      "community4"
    ],
    "prefix": [
      "prefix5",
      "prefix6",
      "prefix7"
    ],
    "protocol": [
      "bgp",
      "direct"
    ]
  },
  "name": "name4"
}
```

