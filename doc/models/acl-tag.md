
# Acl Tag

## Structure

`AclTag`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `GbpTag` | `*int` | Optional | required if<br><br>- `type`==`dynamic_gbp` (gbp_tag received from RADIUS)<br>- `type`==`static_gbp` (applying gbp tag against matching conditions) |
| `Macs` | `[]string` | Optional | required if<br><br>- `type`==`mac`<br>- `type`==`static_gbp` if from matching mac |
| `Network` | `*string` | Optional | if:<br><br>* `type`==`mac` (optional. default is `any`)<br>* `type`==`subnet` (optional. default is `any`)<br>* `type`==`network`<br>* `type`==`resource` (optional. default is `any`)<br>* `type`==`static_gbp` if from matching network (vlan)' |
| `RadiusGroup` | `*string` | Optional | required if:<br><br>* `type`==`radius_group`<br>* `type`==`static_gbp`<br>  if from matching radius_group |
| `Specs` | [`[]models.AclTagSpec`](../../doc/models/acl-tag-spec.md) | Optional | if `type`==`resource`<br>empty means unrestricted, i.e. any |
| `Subnets` | `[]string` | Optional | if<br><br>- `type`==`subnet`<br>- `type`==`resource` (optional. default is `any`)<br>- `type`==`static_gbp` if from matching subnet |
| `Type` | [`models.AclTagTypeEnum`](../../doc/models/acl-tag-type-enum.md) | Required | enum: `any`, `dynamic_gbp`, `mac`, `network`, `radius_group`, `resource`, `static_gbp`, `subnet` |

## Example (as JSON)

```json
{
  "gbp_tag": 60,
  "macs": [
    "macs9",
    "macs0"
  ],
  "network": "network2",
  "radius_group": "radius_group8",
  "specs": [
    {
      "port_range": "port_range8",
      "protocol": "protocol6"
    },
    {
      "port_range": "port_range8",
      "protocol": "protocol6"
    }
  ],
  "type": "static_gbp"
}
```

