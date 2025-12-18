
# Sw Routing Policy

## Structure

`SwRoutingPolicy`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Terms` | [`[]models.SwRoutingPolicyTerm`](../../doc/models/sw-routing-policy-term.md) | Optional | at least criteria/filter must be specified to match the term, all criteria have to be met<br><br>**Constraints**: *Minimum Items*: `1`, *Unique Items Required* |

## Example (as JSON)

```json
{
  "terms": [
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
      "name": "name8"
    },
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
      "name": "name8"
    },
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
      "name": "name8"
    }
  ]
}
```

