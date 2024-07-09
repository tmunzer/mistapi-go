
# Gateway Port Mirroring Port Mirror

## Structure

`GatewayPortMirroringPortMirror`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `FamilyType` | `*string` | Optional | - |
| `IngressPortIds` | `[]string` | Optional | - |
| `OutputPortId` | `*string` | Optional | - |
| `Rate` | `*int` | Optional | - |
| `RunLength` | `*int` | Optional | **Constraints**: `>= 0` |

## Example (as JSON)

```json
{
  "output_port_id": "ge-0/0/5",
  "family_type": "family_type0",
  "ingress_port_ids": [
    "ingress_port_ids2"
  ],
  "rate": 108,
  "run_length": 0
}
```

