
# Org Apitoken

Org API Token

**Note:**
`privilege` field is required to create the object, but may not be
returned in the POST API Response (only in the afterward GET)

## Structure

`OrgApitoken`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CreatedBy` | `models.Optional[string]` | Optional | email of the token creator / null if creator is deleted |
| `CreatedTime` | `*float64` | Optional | - |
| `Id` | `*uuid.UUID` | Optional | - |
| `Key` | `*string` | Optional | - |
| `LastUsed` | `models.Optional[float64]` | Optional | - |
| `Name` | `string` | Required | name of the token |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `Privileges` | [`[]models.PrivilegeOrg`](../../doc/models/privilege-org.md) | Optional | list of privileges the token has on the orgs/sites<br>**Constraints**: *Minimum Items*: `1`, *Maximum Items*: `10`, *Unique Items Required* |
| `SrcIps` | `[]string` | Optional | list of allowed IP addresses from where the token can be used from. At most 10 IP addresses can be specified, cannot be changed once the API Token is created. |

## Example (as JSON)

```json
{
  "created_by": "user@mycorp.com",
  "created_time": 1626875902.0,
  "id": "497f6eca-6276-4993-bfeb-53ecbbba6f08",
  "key": "1qkb...QQCL",
  "last_used": 1690115110.0,
  "name": "org_token_xyz",
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "privileges": [
    {
      "role": "admin",
      "scope": "org"
    }
  ],
  "src_ips": [
    "63.3.56.0/24",
    "63.3.55.4"
  ]
}
```

