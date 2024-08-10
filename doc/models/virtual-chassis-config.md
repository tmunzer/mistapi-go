
# Virtual Chassis Config

Virtual Chassis

## Structure

`VirtualChassisConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Locating` | `*bool` | Optional | - |
| `Members` | [`[]models.VirtualChassisConfigMember`](../../doc/models/virtual-chassis-config-member.md) | Optional | - |
| `Preprovisioned` | `*bool` | Optional | To create the Virtual Chassis in Pre-Provisioned mode<br>**Default**: `false` |

## Example (as JSON)

```json
{
  "preprovisioned": false,
  "locating": false,
  "members": [
    {
      "locating": false,
      "mac": "mac2",
      "member": 176,
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
      "member": 176,
      "vc_ports": [
        "vc_ports2",
        "vc_ports3",
        "vc_ports4"
      ],
      "vc_role": "master"
    }
  ]
}
```

