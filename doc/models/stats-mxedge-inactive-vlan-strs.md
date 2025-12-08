
# Stats Mxedge Inactive Vlan Strs

Inactive wired/L2TP VLANs. Entries can be individual VLANs or ranges.

## Structure

`StatsMxedgeInactiveVlanStrs`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `L2tp` | `[]string` | Optional | Inactive L2TP VLANs. Entries can be individual VLANs or ranges. |
| `Wired` | `[]string` | Optional | Inactive wired VLANs. Entries can be individual VLANs or ranges. |

## Example (as JSON)

```json
{
  "wired": [
    "100",
    "102-106"
  ],
  "l2tp": [
    "l2tp1"
  ]
}
```

