
# Utils Clear Session

to use five tuples to lookup the session to be cleared, all must be provided

## Structure

`UtilsClearSession`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ServiceName` | `*string` | Optional | ervice name, only supported in SSR |
| `SessionIds` | `[]uuid.UUID` | Optional | List of id of the sessions to be cleared |

## Example (as JSON)

```json
{
  "service_name": "internet-wan_and_lte",
  "session_ids": [
    "88776655-0123-4567-890a-112233445566"
  ]
}
```

