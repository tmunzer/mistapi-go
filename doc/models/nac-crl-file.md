
# Nac Crl File

## Structure

`NacCrlFile`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CreatedTime` | `*float64` | Optional | When the object has been created, in epoch |
| `Id` | `*string` | Optional | Unique ID for the uploaded CRL file, used to reference the file |
| `ModifiedTime` | `*float64` | Optional | When the object has been modified for the last time, in epoch |
| `Name` | `*string` | Optional | Issuer name for the CRL file |
| `Url` | `*string` | Optional | URL to download the uploaded CRL file |

## Example (as JSON)

```json
{
  "id": "a1ca26f3-44dd-4833-9a7b-97bbb2ab5230",
  "name": "SampleCertificateSigner",
  "url": "http://url/to/crl_file",
  "created_time": 107.38,
  "modified_time": 227.58
}
```

