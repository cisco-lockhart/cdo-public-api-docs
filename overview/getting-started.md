# Getting Started

The following sections introduce you to the resources in the Security Cloud Control API, and provide instructions to make your first API request. Examples will be shown using the command line with cURL.

## Supported RESTful Operations
The resources in the Security Cloud Control API support one or more of the `GET`, `POST`, `PATCH`, `PUT`, and `DELETE` operations. 

>**Note**:
Not all of the resources support all of these operations. 


- **GET**: There are two kinds of `GET` operations in Security Cloud Control: list `GET` operations, which return a paged list of resources, and individual `GET` operations which return a single resource by UID. In the list endpoints, it is possible to search and sort by a subset of the fields in the resource.
- **POST**: `POST` operations are used in Security Cloud Control to either create a resource or trigger an asynchronous operation. All POST operations with an action verb trigger asynchronous operations. Asynchronous operations can be tracked using the Security Cloud Control transaction API.
- **PATCH**: `PATCH` operations are used in Security Cloud Control to modify resources.
- **PUT**: `PUT` operations are used by the Cloud-delivered FMC API endpoints to modify resources.
- **DELETE**: `DELETE` operations are used in Security Cloud Control to delete resources. All `DELETE` operations in the Security Cloud Control API are idempotent.

All responses from the Security Cloud Control API are in the JSON format. Additionally, all payloads must be sent to the Security Cloud Control API as JSON.

## API Resources
At a high level, the Security Cloud Control API supports operations on the following resources:
- **Inventory Management** - Manage devices, device managers, cloud services, and templates in your Security Cloud Control tenant.
- **Connector Management** - Add, remove, and view information on connectors used to communicate with devices in your Security Cloud Control tenant.
- **Cloud-delivered FMC** - Manage the Cloud-delivered FMC in your Security Cloud Control tenant (if present).
- **Object Management** - Manage firewall policy objects in your Security Cloud Control tenant.
- **User Management** - Manage users in your Security Cloud Control tenant.
- **Tenant Management** - Manage your Security Cloud Control tenant, or Managed Service Portal.
- **Remote Access Monitoring** - View and manage the Remote Access Virtual Private Network (RA VPN) sessions, and Multi-factor Authentication (MFA) events in your Security Cloud Control tenant.
- **Search** - Perform searches across all of the resources in your Security Cloud Control tenant.
- **Changelogs** - View a detailed history of all changes made to your Security Cloud Control tenant.
- **Change Requests** - Manage change requests, which can be used to associate changes made to your security policy on Security Cloud Control with external change-management systems.
- **Transaction Management** - Security Cloud Control transactions are used to track the progress of asynchronous operations triggered using the API on your Security Cloud Control tenant. For example, use Security Cloud Control transactions to monitor the progress of deploying configuration to an ASA.
- **Meta endpoints** - Meta information about Security Cloud Control itself.

## API User Prerequisites

To use the Security Cloud Control API, you must have an API token. We recommend creating an [API-only user](https://docs.defenseorchestrator.com/c-secure-device-connector-sdc.html#!t-create-api-only-users.html) in your Security Cloud Control tenant, and generating a token for it.

In order to perform the operation described in this page, the API-only user should at least have the **Read Only** role. However, different endpoints in the Security Cloud Control API can require admin or super-admin privileges.

## Base URI

Cisco Security Cloud Control is deployed in multiple regions, and the API is available in all of these. The server URL for each region is as follows:

- US: https://api.us.security.cisco.com/firewall (use this if the Security Cloud Control URL you visit is https://www.defenseorchestrator.com or https://us.manage.security.cisco.com)
- EU: https://api.eu.security.cisco.com/firewall (use this if the Security Cloud Control URL you visit is https://www.defenseorchestrator.eu or https://eu.manage.security.cisco.com)
- APJ: https://api.apj.security.cisco.com/firewall (use this if the Security Cloud Control URL you visit is https://apj.cdo.cisco.com or https://apj.manage.security.cisco.com)
- Australia: https://api.au.security.cisco.com/firewall (use this if the Security Cloud Control URL you visit is https://aus.cdo.cisco.com or https://aus.manage.security.cisco.com)
- India: https://api.in.security.cisco.com/firewall (use this if the Security Cloud Control URL you visit is https://in.cdo.cisco.com or https://in.manage.security.cisco.com)

>**Note**:
If you were using the old base URIs (`https://edge.<region>.cdo.cisco.com`), they will continue to work. But please switch your automation to use the new URLs.

## 1. Authorization

In addition to the base URL, an Authorization header must be added to every API request with the following format.
```
Authorization: Bearer $API_TOKEN
```
where `$API_TOKEN` is an API token for your Security Cloud Control tenant.

## 2. Get a list of devices

### Request

Use this `CuRL` command to make a request to retrieve the list of devices from the `/api/rest/v1/inventory/devices` endpoint.

```curl
curl -X GET \                                                                                                             08:48:01
--url https://edge.us.cdo.cisco.com/api/rest/v1/inventory/devices \
--header "Authorization: Bearer $API_TOKEN"
```

### Response

The actual response returned will be different depending on the devices in your Security Cloud Control tenant, but it should look something like this.

```json
{
  "count": 2,
  "limit": 50,
  "offset": 0,
  "items": [
    {
      "uid": "667cc77e-1bbd-4af8-9fa3-8211c75e818c",
      "name": "Richmond",
      "deviceType": "ASA",
      "connectorType": "SDC",
      "connectorUid": "0184ed35-53f7-465a-9ba6-12b814773018",
      "address": "x.y.z.a:443",
      "serial": "xxxxx",
      "chassisSerial": "xxxxxx",
      "softwareVersion": "9.15(1)",
      "connectivityState": "ONLINE",
      "configState": "NOT_SYNCED",
      "conflictDetectionState": "NO_CONFLICTS",
      "asdmVersion": "7.15(1)150",
      "asaFailoverMode": "OFF",
      "asaLicenseEntitlements": {
        "Encryption-DES": "Enabled",
        "AnyConnect Essentials": "Disabled",
        "Encryption-3DES-AES": "Enabled",
        "Cluster": "Disabled",
        "Maximum VLANs": "50",
        "Inside Hosts": "Unlimited",
        "Other VPN Peers": "250",
        "Shared License": "Disabled",
        "AnyConnect for Cisco VPN Phone": "Disabled",
        "Botnet Traffic Filter": "Enabled",
        "AnyConnect Premium Peers": "0",
        "Failover": "Active/Standby",
        "Advanced Endpoint Assessment": "Disabled",
        "Total TLS Proxy Sessions": "0",
        "Carrier": "Disabled",
        "AnyConnect for Mobile": "Disabled",
        "Total VPN Peers": "250",
        "Security Contexts": "0"
      },
      "state": "DONE",
      "stateMachineDetails": {
        "identifier": "deviceProcessRequestsStateMachine"
      },
      "labels": {
        "groupedLabels": {
          "Region": [
            "West",
            "North"
          ],
          "SecureX": [
            "WF001"
          ],
          "GROUP1": [
            "bar"
          ]
        },
        "ungroupedLabels": [
          "foo"
        ]
      }
    },
    {
      "uid": "1c41932d-6310-437b-ba22-9a97e08d8cbb",
      "name": "Clarksville",
      "deviceType": "ASA",
      "connectorType": "SDC",
      "connectorUid": "0184ed35-53f7-465a-9ba6-12b814773018",
      "address": "a.b.c.d:8443",
      "serial": "xxxxx",
      "chassisSerial": "xxxxxxxx",
      "softwareVersion": "9.16(3)19",
      "connectivityState": "ONLINE",
      "configState": "SYNCED",
      "conflictDetectionState": "NO_CONFLICTS",
      "asdmVersion": "7.18(1)152",
      "asaFailoverMode": "OFF",
      "asaLicenseEntitlements": {
        "Encryption-DES": "Enabled",
        "AnyConnect Essentials": "Disabled",
        "Encryption-3DES-AES": "Enabled",
        "Cluster": "Disabled",
        "Maximum VLANs": "50",
        "Inside Hosts": "Unlimited",
        "Other VPN Peers": "250",
        "Shared License": "Disabled",
        "AnyConnect for Cisco VPN Phone": "Enabled",
        "Botnet Traffic Filter": "Enabled",
        "AnyConnect Premium Peers": "250",
        "Failover": "Active/Standby",
        "Advanced Endpoint Assessment": "Enabled",
        "Total TLS Proxy Sessions": "500",
        "Carrier": "Enabled",
        "AnyConnect for Mobile": "Enabled",
        "Total VPN Peers": "250",
        "Security Contexts": "0"
      },
      "state": "DONE",
      "stateMachineDetails": {
        "identifier": "deviceCreateObjectsStateMachine"
      },
      "labels": {
        "groupedLabels": {
          "FLOW": [
            "BACKUPS"
          ]
        }
      }
    }
  ]
}

```


The response returned is a JSON object that has some metadata on the request itself, including the:
* **count**: The number of items in the collection. If this number is greater than `limit`, the number of items returned will be less than the total number of items in the list.
* **limit**: The maximum number of items returned by this request.
* **offset**: The offset from which the list was retrieved.
* **items:** The items in the response.

>**Note**:
The `limit` and `offset` values can be set as query parameters to the request, and can be used to page through a list. By default, `limit` is set to `50`, and `offset` to `0`.


## Try it out with Postman
Prefer to use Postman rather than code or the command line? Check out our [Postman Collection](#!postman-collection/postman-collection).