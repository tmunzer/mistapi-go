
# Privilege Msp

Privilieges settings

## Structure

`PrivilegeMsp`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `OrgId` | `*uuid.UUID` | Optional | if `scope`==`org` |
| `OrgName` | `*string` | Optional | name of the org (for a site belonging to org) |
| `OrggroupId` | `*uuid.UUID` | Optional | if `scope`==`orggroup` |
| `Role` | [`models.PrivilegeMspRoleEnum`](../../doc/models/privilege-msp-role-enum.md) | Required | access permissions |
| `Scope` | [`models.PrivilegeMspScopeEnum`](../../doc/models/privilege-msp-scope-enum.md) | Required | - |
| `Views` | [`*models.PrivilegeMspViewEnum`](../../doc/models/privilege-msp-view-enum.md) | Optional | Custom roles restrict Org users to specific UI views. This is useful for limiting UI access of Org users.<br><br>You can invite a new user or update existing users in your Org to this custom role. For this, specify view along with role when assigning privileges.<br><br>Below are the list of supported UI views. Note that this is UI only feature<br>Custom roles restrict Org users to specific UI views. This is useful for limiting UI access of Org users.<br><br>You can invite a new user or update existing users in your Org to this custom role. For this, specify `view` along with `role` when assigning privileges.<br><br>Below are the list of supported UI views. Note that this is UI only feature<br><br>\| UI View \| Description \|<br>\| --- \| --- \|<br>\| `reporting` \| full access to all analytics tools \|<br>\| `marketing` \| can view analytics and location maps \|<br>\| `location` \| can view and manage location maps \|<br>\| `security` \| can view and manage WLAN, rogues and authentication \|<br>\| `switch_admin` \| can view and manage Switch ports \|<br>\| `mxedge_admin` \| can view and manage Mist edges and Mist tunnels \|<br>\| `lobby_admin` \| full access to Org and Site Pre-shared keys \| |

## Example (as JSON)

```json
{
  "org_id": "00002258-0000-0000-0000-000000000000",
  "org_name": "org_name6",
  "orggroup_id": "00002672-0000-0000-0000-000000000000",
  "role": "admin",
  "scope": "orggroup",
  "views": "mxedge_admin"
}
```
