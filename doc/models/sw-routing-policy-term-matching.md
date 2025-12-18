
# Sw Routing Policy Term Matching

zero or more criteria/filter can be specified to match the term, all criteria have to be met

## Structure

`SwRoutingPolicyTermMatching`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AsPath` | [`[]models.BgpAs`](../../doc/models/containers/bgp-as.md) | Optional | BGP AS, value in range 1-4294967294. Can be a Variable (e.g. `{{bgp_as}}` ) |
| `Community` | `[]string` | Optional | - |
| `Prefix` | `[]string` | Optional | zero or more criteria/filter can be specified to match the term, all criteria have to be met |
| `Protocol` | [`[]models.SwRoutingPolicyTermMatchingProtocolEnum`](../../doc/models/sw-routing-policy-term-matching-protocol-enum.md) | Optional | - |

## Example (as JSON)

```json
{
  "as_path": [
    "String1"
  ],
  "community": [
    "community4"
  ],
  "prefix": [
    "prefix7",
    "prefix8",
    "prefix9"
  ],
  "protocol": [
    "evpn",
    "ospf"
  ]
}
```

