# Go client for managing Citrix DaaS resources

## About

The Citrix DaaS Rest Client for GO enables managing Citrix DaaS resources programmatically in GO applications.

## Installation
- From github

```
go get github.com/citrix/citrix-daas-rest-go
```

## Usage
The Citrix DaaS Go client contains the following two packages:
1. `github.com/citrix/citrix-daas-rest-go/citrixorchestration`: SDK which contains the client model and api functions for Citrix DaaS API. This is auto-generated by OpenAPI Generator and do not require any manual changes.
2. `github.com/citrix/citrix-daas-rest-go/client`: Rest client configurable with Citrix DaaS customer details and handles authentication.

Instantiate a client using `NewCitrixDaasClient`. The following variables are required to configure Citrix DaaS client:

|              | Cloud                                                                                                                        | On-Premises                                                                                                            |
|--------------|------------------------------------------------------------------------------------------------------------------------------|------------------------------------------------------------------------------------------------------------------------|
| authUrl      | Citrix Cloud authentication URL, i.e. `https://api-us.cloud.com/cctrustoauth2/{customerId}/tokens/clients` for US customers. | On-premises trust service URL, i.e. `https://{deliveryControllerHostname}/citrix/orchestration/api/techpreview/tokens` |
| hostname     | Citrix Cloud DaaS service hostname, i.e. `{customerId}.xendesktop.net`                                                       | Delivery Controller Hostname / IP address                                                                              |
| customerId   | Cloud Customer Id                                                                                                            | `CitrixOnPremises`                                                                                                     |
| clientId     | Citrix Cloud API Key clientId                                                                                                | Domain Admin Username                                                                                                  |
| clientSecret | Citrix Cloud API Key clientSecret                                                                                            | Domain Admin Password                                                                                                  |
| onPremise    | `false`                                                                                                                      | `true`                                                                                                                 |


## Example

```
package main

import (
        "context"

        citrixorchestration "github.com/citrix/citrix-daas-rest-go/citrixorchestration"
        citrixclient "github.com/citrix/citrix-daas-rest-go/client"
)

func main() {
        // Create a new Citrix API client for cloud customer 83czxoqlpepv
        authUrl := "https://api-us.cloud.com/cctrustoauth2/83czxoqlpepv/tokens/clients"
        hostname := "83czxoqlpepv.xendesktop.net"
        customerId := "83czxoqlpepv"
        clientId := "{apiKeyClientId}"
        clientSecret := "{apiKeyClientSecret}"
        onPremise := false
        client, err := citrixclient.NewCitrixDaasClient(authUrl, hostname, customerId, clientId, clientSecret, onPremise)
        
        // Create GET zone request with zoneId and siteId
        zoneId := "8994379b-8585-4717-9765-632992e738d3"
        siteId := "8509b2e6-ff02-48bb-aaa3-07b84f09b9f4"
        getZoneRequest := client.ApiClient.ZonesTPApi.ZonesTPGetZone(ctx, zoneId, customerId, siteId)

        // Get bearer token and put into authorization header
        token, _ := client.SignIn()
        getZoneRequest = getZoneRequest.Authorization(token)

        // Execute request and get response
        zone, _, err := getZoneRequest.Execute()
        if err != nil {
            return
        }
}

```
