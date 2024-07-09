
# Service

Applications used for the Gateway configurations

## Structure

`Service`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Addresses` | `[]string` | Optional | if `type`==`custom`, ip subnets |
| `AppCategories` | `[]string` | Optional | when `type`==`app_categories`<br>list of application categories are available through /api/v1/const/app_categories |
| `AppSubcategories` | `[]string` | Optional | when `type`==`app_categories`<br>list of application categories are available through /api/v1/const/app_subcategories |
| `Apps` | `[]string` | Optional | when `type`==`apps`<br>list of applications are available through:<br><br>- /api/v1/const/applications,<br>- /api/v1/const/gateway_applications<br>- /insight/top_app_by-bytes?wired=true |
| `CreatedTime` | `*int` | Optional | - |
| `Description` | `*string` | Optional | - |
| `Dscp` | `*int` | Optional | when `traffic_type`==`custom` |
| `FailoverPolicy` | [`*models.ServiceFailoverPolicyEnum`](../../doc/models/service-failover-policy-enum.md) | Optional | **Default**: `"revertable"` |
| `Hostnames` | `[]string` | Optional | if `type`==`custom`, web filtering |
| `Id` | `*uuid.UUID` | Optional | - |
| `MaxJitter` | `*int` | Optional | when `traffic_type`==`custom`, for uplink selection |
| `MaxLatency` | `*int` | Optional | when `traffic_type`==`custom`, for uplink selection |
| `MaxLoss` | `*int` | Optional | when `traffic_type`==`custom`, for uplink selection |
| `ModifiedTime` | `*int` | Optional | - |
| `Name` | `*string` | Optional | - |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `SleEnabled` | `*bool` | Optional | whether to enable measure SLE<br>**Default**: `false` |
| `Specs` | [`[]models.ServiceSpec`](../../doc/models/service-spec.md) | Optional | - |
| `SsrRelaxedTcpStateEnforcement` | `*bool` | Optional | **Default**: `false` |
| `TrafficClass` | [`*models.ServiceTrafficClassEnum`](../../doc/models/service-traffic-class-enum.md) | Optional | when `traffic_type`==`custom`<br>**Default**: `"best_effort"` |
| `TrafficType` | `*string` | Optional | values from `/api/v1/consts/traffic_types`<br><br>* when `type`==`apps`, we'll choose traffic_type automatically<br>* when `type`==`addresses` or `type`==`hostnames`, you can provide your own settings (optional)<br>**Default**: `"data_best_effort"` |
| `Type` | [`*models.ServiceTypeEnum`](../../doc/models/service-type-enum.md) | Optional | **Default**: `"custom"` |
| `Urls` | `[]string` | Optional | when `type`==`urls no need for spec as URL can encode the ports being used` |

## Example (as JSON)

```json
{
  "app_categories": [
    "Sports"
  ],
  "app_subcategories": [
    "Shopping"
  ],
  "apps": [
    "office365",
    "okta"
  ],
  "failover_policy": "revertable",
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "sle_enabled": false,
  "ssr_relaxed_tcp_state_enforcement": false,
  "traffic_class": "best_effort",
  "traffic_type": "data_best_effort",
  "type": "custom",
  "addresses": [
    "addresses8",
    "addresses9",
    "addresses0"
  ],
  "created_time": 242
}
```

