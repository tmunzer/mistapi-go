
# Virtual Chassis Port

## Structure

`VirtualChassisPort`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Members` | [`[]models.ConfigVcPortMember`](../../doc/models/config-vc-port-member.md) | Required | **Constraints**: *Minimum Items*: `1`, *Unique Items Required* |
| `Op` | [`models.VirtualChassisPortOperationEnum`](../../doc/models/virtual-chassis-port-operation-enum.md) | Required | **Constraints**: *Minimum Length*: `1` |

## Example (as JSON)

```json
{
  "members": [
    {
      "member": 188.64,
      "vc_ports": [
        "vc_ports2",
        "vc_ports3",
        "vc_ports4"
      ]
    }
  ],
  "op": "set"
}
```

