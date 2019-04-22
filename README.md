# Masternode Hosting Solution

This is a solution which allows for "fake" masternode hosting.

## Installation and Usage
```
git clone https://gitlab.com/jackkdev/phantom-hosting-api.git
cd phantom-hosting-api
dep ensure
go run main.go
```

## Making requests
### Generate a masternode configuration string

```http request
POST http://localhost:8000/generateconfig
```
#### Request Body
```json
{
  "port": 9998,
  "genkey": "75eqvNfaEfkd3YTwQ3hMwyxL2BgNSrqHDgWc6jbUh4Gdtnro2Wo",
  "txid": "f8a3e39da2d13e10736a77940a2a78823e30e3ac40140f0a0b1ec31d07989aef",
  "tx_index": 1
}
```
#### Request Response
```json
{
    "success": true,
    "data": "331720b1-6d69-404c-b84e-932642c93e92 [5a67:ae46:afa1:fd29:35a:2b37:dd1d:b138]:9998 75eqvNfaEfkd3YTwQ3hMwyxL2BgNSrqHDgWc6jbUh4Gdtnro2Wo f8a3e39da2d13e10736a77940a2a78823e30e3ac40140f0a0b1ec31d07989aef 1 1555938586",
    "error": ""
}
```