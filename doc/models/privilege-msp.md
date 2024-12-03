
# Privilege Msp

Privilieges settings

*This model accepts additional fields of type interface{}.*

## Structure

`PrivilegeMsp`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `OrgId` | `*uuid.UUID` | Optional | if `scope`==`org` |
| `OrgName` | `*string` | Optional | name of the org (for a site belonging to org) |
| `OrggroupId` | `*uuid.UUID` | Optional | if `scope`==`orggroup` |
| `Role` | [`models.PrivilegeMspRoleEnum`](../../doc/models/privilege-msp-role-enum.md) | Required | access permissions. enum: `admin`, `helpdesk`, `installer`, `read`, `write` |
| `Scope` | [`models.PrivilegeMspScopeEnum`](../../doc/models/privilege-msp-scope-enum.md) | Required | enum: `msp`, `org`, `orggroup` |
| `Views` | [`*models.AdminPrivilegeViewEnum`](../../doc/models/admin-privilege-view-enum.md) | Optional | Custom roles restrict Org users to specific UI views. This is useful for limiting UI access of Org users.<br><br>You can invite a new user or update existing users in your Org to this custom role. For this, specify view along with role when assigning privileges.<br><br>Below are the list of supported UI views. Note that this is UI only feature<br>Custom roles restrict Org users to specific UI views. This is useful for limiting UI access of Org users.<br><br>You can invite a new user or update existing users in your Org to this custom role. For this, specify `view` along with `role` when assigning privileges.<br><br>Below are the list of supported UI views. Note that this is UI only feature<br><br>\| UI View \| Required Role \| Description \|<br>\| --- \| --- \| --- \|<br>\| `reporting` \| `read` \| full access to all analytics tools \|<br>\| `marketing` \| `read` \| can view analytics and location maps \|<br>\| `super_observer` \| `read` \| can view all the organization except the subscription page \|<br>\| `location` \| `write` \| can view and manage location maps, can view analytics \|<br>\| `security` \| `write` \| can view and manage site labels, policies and security \|<br>\| `switch_admin` \| `helpdesk` \| can view and manage Switch ports, can view wired clients \|<br>\| `mxedge_admin` \| `admin` \| can view and manage Mist edges and Mist tunnels \|<br>\| `lobby_admin` \| `admin` \| full access to Org and Site Pre-shared keys \| |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "org_id": "00001be0-0000-0000-0000-000000000000",
  "org_name": "org_name0",
  "orggroup_id": "00000716-0000-0000-0000-000000000000",
  "role": "helpdesk",
  "scope": "org",
  "views": "super_observer",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

