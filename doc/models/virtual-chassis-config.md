
# Virtual Chassis Config

Virtual Chassis

*This model accepts additional fields of type interface{}.*

## Structure

`VirtualChassisConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Locating` | `*bool` | Optional | - |
| `Members` | [`[]models.VirtualChassisConfigMember`](../../doc/models/virtual-chassis-config-member.md) | Optional | - |
| `Preprovisioned` | `*bool` | Optional | To create the Virtual Chassis in Pre-Provisioned mode<br><br>**Default**: `false` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "preprovisioned": false,
  "locating": false,
  "members": [
    {
      "locating": false,
      "mac": "mac2",
      "member_id": 58,
      "vc_ports": [
        "vc_ports2",
        "vc_ports3",
        "vc_ports4"
      ],
      "vc_role": "master"
    },
    {
      "locating": false,
      "mac": "mac2",
      "member_id": 58,
      "vc_ports": [
        "vc_ports2",
        "vc_ports3",
        "vc_ports4"
      ],
      "vc_role": "master"
    }
  ],
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

