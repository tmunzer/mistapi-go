
# Acl Tag

Resource tags (`type`==`resource` or `type`==`gbp_resource`) can only be used in `dst_tags`

*This model accepts additional fields of type interface{}.*

## Structure

`AclTag`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `EtherTypes` | `[]string` | Optional | ARP / IPv6. Default is `any` |
| `GbpTag` | `*int` | Optional | Required if<br><br>- `type`==`dynamic_gbp` (gbp_tag received from RADIUS)<br>- `type`==`gbp_resource`<br>- `type`==`static_gbp` (applying gbp tag against matching conditions) |
| `Macs` | `[]string` | Optional | Required if<br><br>- `type`==`mac`<br>- `type`==`static_gbp` if from matching mac |
| `Network` | `*string` | Optional | If:<br><br>* `type`==`mac` (optional. default is `any`)<br>* `type`==`subnet` (optional. default is `any`)<br>* `type`==`network`<br>* `type`==`resource` (optional. default is `any`)<br>* `type`==`static_gbp` if from matching network (vlan) |
| `PortUsage` | `*string` | Optional | Required if `type`==`port_usage` |
| `RadiusGroup` | `*string` | Optional | Required if:<br><br>* `type`==`radius_group`<br>* `type`==`static_gbp`<br>  if from matching radius_group |
| `Specs` | [`[]models.AclTagSpec`](../../doc/models/acl-tag-spec.md) | Optional | If `type`==`resource`, `type`==`radius_group`, `type`==`port_usage` or `type`==`gbp_resource`. Empty means unrestricted, i.e. any |
| `Subnets` | `[]string` | Optional | If<br><br>- `type`==`subnet`<br>- `type`==`resource` (optional. default is `any`)<br>- `type`==`static_gbp` if from matching subnet |
| `Type` | [`models.AclTagTypeEnum`](../../doc/models/acl-tag-type-enum.md) | Required | enum:<br><br>* `any`: matching anything not identified<br>* `dynamic_gbp`: from the gbp_tag received from RADIUS<br>* `gbp_resource`: can only be used in `dst_tags`<br>* `mac`<br>* `network`<br>* `port_usage`<br>* `radius_group`<br>* `resource`: can only be used in `dst_tags`<br>* `static_gbp`: applying gbp tag against matching conditions<br>* `subnet`' |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "ether_types": [
    "ether_types4"
  ],
  "gbp_tag": 132,
  "macs": [
    "macs5",
    "macs6"
  ],
  "network": "network4",
  "port_usage": "port_usage4",
  "type": "any",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

