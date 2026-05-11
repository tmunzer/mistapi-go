
# Org E911 Report

E911 AP BSSID report status for the organization

## Structure

`OrgE911Report`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Detail` | `*string` | Optional | Human-readable description of the action taken |
| `LastGenerated` | `*int` | Optional | Unix timestamp of when the report file was last generated. Only present when `status` is `available`. |
| `Status` | [`*models.OrgE911ReportStatusEnum`](../../doc/models/org-e911-report-status-enum.md) | Optional | Current status of E911 report generation. enum: `disabled`, `scheduled`, `available` |
| `Url` | `*string` | Optional | Presigned URL to download the CSV file. Only present when `status` is `available`. |

## Example (as JSON)

```json
{
  "detail": "detail2",
  "last_generated": 248,
  "status": "available",
  "url": "url0"
}
```

