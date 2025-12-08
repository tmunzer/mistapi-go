
# Org Crl Import File

## Structure

`OrgCrlImportFile`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `File` | `*string` | Optional | a PEM or DER formatted CRL file |
| `Json` | `*string` | Optional | a JSON string with "name" field for CRL file issuer (optional) |

## Example (as JSON)

```json
{
  "file": "file4",
  "json": "json2"
}
```

