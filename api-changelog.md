# Version 1.13.0 (2025-05-16)

## Improvements
- You can now upgrade up to 50 FTDs simultaneously using the [Upgrade multiple FTD devices](https://developer.cisco.com/docs/cisco-security-cloud-control/upgrade-multiple-ftd-devices/) endpoint.
- It is now possible to store arbitrary metadata in a device record on SCC Firewall Manager. This
  metadata is not used by the API, but can be used by other applications to store information
  related to the device.

# Version 1.12.0 (2025-05-06)

## Added

- Endpoint
  to [rebuild cache for cdFMC upgrade packages](https://developer.cisco.com/docs/cisco-security-cloud-control/update-cache-of-compatible-upgrade-packages-for-all-ftds/)
  on a tenant.

# Version 1.11.0 (2025-04-27)

## Added

- Endpoint
  to [upgrade multiple (up to 10) FTDs simultaneously](https://developer.cisco.com/docs/cisco-security-cloud-control/upgrade-multiple-ftd-devices/)
- Endpoint
  to [read configuration from multiple ASAs](https://developer.cisco.com/docs/cisco-security-cloud-control/read-configurations-for-multiple-asa-devices/)
  into SCC Firewall Manager.
- Endpoint
  to [accept certificates on multiple ASAs](https://developer.cisco.com/docs/cisco-security-cloud-control/accept-certificates-for-multiple-asa-devices/)
  into SCC Firewall Manager.
- Endpoint
  to [delete multiple users](https://developer.cisco.com/docs/cisco-security-cloud-control/delete-multiple-users-from-security-cloud-control-firewall-manager-tenant/)
  on a tenant.

## Fixes

- Fixed regression that prevented searches on users by name working.
- Fixed performance issues when triggering deploys on FTDs managed by cdFMCs with large numbers of
  devices.
- Fixed performance issues when triggering upgrades on FTDs managed by cdFMCs with large numbers of
  devices.

# Version 1.10.0 (2025-04-01)

## Added

- The Cloud-delivered Firewall Management Center (cdFMC) API docs are now available as part of the
  Security Cloud Control Firewall Manager API documentation.
  See [here](https://developer.cisco.com/docs/cisco-security-cloud-control/cdfmc-api-overview/) for
  more information.

## Improvements

- Improved the accuracy and quality of S2S VPN tunnel and RA VPN session information provided by the
  cdFMC health monitoring endpoint.

# Version 1.9.0 (2025-03-17)

## Added

- Endpoints to upgrade FTD devices
    - Identify which versions are compatible with a given FTD device
    - Upgrade FTD device to compatible version
- Endpoint to deploy changes to FTD devices
- Endpoint to fetch API-only users on a tenant
- Endpoint to fetch targets associated with an object

# Version 1.8.0 (2025-02-05)

## Improvements

- You can now access the Security Cloud Control Firewall APIs from
  `https://api.<your-region>.security.cisco.com/firewall`. The old URLs (
  `https://edge.<region>.cdo.cisco.com/api/rest` and
  `https://<region>.manage.security.cisco.com/api/rest` will continue to work).
- Support added to perform zero-touch provisioning of FTDs using cdFMC templates.

# Version 1.7.0 (2025-01-24)

## Added

- Endpoints to upgrade ASA devices.
- Endpoints to execute arbitrary API endpoints on an on-prem FMC.
- Endpoint to fetch the raw config of an ASA.

## Improvements

- MSSP tenant management:
    - When creating a SCC tenant, you can now optionally provide a sales order number.
- Get Devices endpoint:
    - Provide High Availability (HA) and Clustering information for on-prem FMC-managed FTDs. Note:
      Unlike with a cdFMC-managed FTD HA pair or cluster, each node of an on-prem FMC-managed FTD HA
      pairs and cluster is represented as a separate device record.
        - Note: Full HA and clustering information for on-prem FMC is currently gated behind a
          feature flag.
    - Add links from FMC-managed devices to the FMC, and access policies managed by the FMC.
    - Improve searchability of the endpoint by adding multiple new searchable fields.
- List VPN Sessions endpoint:
    - It is now possible to search VPN sessions by `loginTime` and `lastActiveTime`.

## Fixes

- Multiple fixes to correct the value of counts displayed in list endpoints.

## Deprecations

- The `uidOnFmc` field is deprecated and will be removed in v2 of the API. It has been supplanted by
  the `cdFmcInfo.deviceRecordOnFmc` and `onPremFmcInfo.deviceRecordOnFmc` fields, which contain
  additional information.
- The `deviceRole` field is deprecated and will be removed in v2 of the API. It has been supplanted
  by the `deviceRoles` field, which is a list and can represent the multiple roles a device can
  take.

# Version 1.6.0 (2024-11-08)

## Welcome to Cisco Security Cloud Control

Cisco Defense Orchestrator is now Security Cloud Control, our unified security management solution
for Cisco Firewalls, Multicloud Defense, and Hypershield.
See [the FAQ](https://www.cisco.com/c/en/us/products/collateral/security/defense-orchestrator/security-cloud-control-faq.html)
to get answers to commonly asked questions.

The existing endpoints, at `<region-name>.cdo.cisco.com`, will continue to work. If you are already
a user of the API, you do not need to change anything.

## Added

- MSSP tenant management:
    - Endpoint to remove a tenant from the MSSP portal.
    - Endpoint to get the list of users associated with a tenant managed by an MSSP portal.
    - Endpoints to remove users and user groups associated with a tenant managed by an MSSP portal.

# Version 1.5.0 (2024-10-08)

## Added

- MSSP Tenant management:
    - Endpoint to list the tenants managed by an MSSP portal.
    - Endpoint to generate the API token for an API-only user in a tenant managed by an MSSP portal.

## Improvements

- MSSP Tenant management: It is now possible to create API-only users in a tenant managed by the
  MSSP portal.
- Changelogs: It is now possible to filter changelogs by date, and also retrieve changelogs over the
  last 5, 15, 30, and 60 minutes.
- ASA device management: The API now returns the chassis serial number of secondary devices in ASA
  HA Pairs.
- cdFMC Health Monitoring:
    - The rate limit for the endpoint has been raised from 1 request per minute to 2 requests per
      minute.
    - The health monitoring endpoint returns device HA status and Power Supply Unit (PSU)
      information.

# Version 1.4.0 (2024-09-12)

## Added

- MSSP tenant management: Endpoint to provision cloud-delivered FMC (cdFMC) on a CDO tenant that is
  associated with a CDO MSSP Portal.
- Inventory management: Endpoint to provision a FTD using the Zero-Touch Provisioning (ZTP)
  workflow.

# Version 1.3.0 (2024-08-22)

## Added

- cdFMC health monitoring: API endpoint to monitor the health of a cloud-delivered FMC (cdFMC).
- Provisioning: Provision a cdFMC and enable Multicloud Defense on a CDO tenant.
- MSSP Tenant management: Endpoints to add user groups, and enable Multicloud Defense on a CDO
  tenant that is associated with a CDO MSSP Portal.
- Improvements to the devices endpoint: The **Get Devices** endpoint now provides information on the
  redundancy mode of a firewall, and provides additional High Availability (HA) and clustering
  information for FTD HA Pairs and clusters.

# Version 1.2.0 (2024-07-19)

## Added

- ASA Access Groups and Access Rules management: API endpoints to manage ASA Access Groups and
  Access Rules.
- MSSP Tenant management: API endpoints for MSSP portal super-admin users to create a new tenant and
  add it to the CDO MSSP portal and add existing tenants to the portal.
- Active Directory (AD) groups management: API endpoints to create, remove, modify, list, and fetch
  Active Directory groups on a CDO tenant.

# Version 1.1.0 (2024-05-30)

## Added

- AI Assistant: API endpoints to start a conversation with the AI assistant, ask questions, and list
  conversations and messages.

# Version 1.0.0 (2024-05-23)

## Added

- Secure Device Connector (SDC) management: API endpoints to create and delete Secure Device
  Connectors (SDCs).
- ASA Command Line Interface (CLI): API endpoints to execute arbitrary CLIs and CLI macros across
  one or more devices (ASAs only).
- RA VPN session management: API endpoints to view, refresh, and terminate RA VPN sessions on RA VPN
  headend devices.
- MFA event monitoring: API endpoints to view MFA events on Duo Admin Panels.
- Object Management: Object Management endpoints are now accessible from `v1` endpoints like the
  rest of CDO's endpoints.

# Version 0.1.0 (2024-04-02)

## Added

This is the first major release of the CDO API, and exposes a large number of CDO features. Some of
the highlights include:

- Simplified Inventory management: API endpoints to add, delete, view, and perform operations on
  devices, device managers, and cloud services available on the CDO inventory page.
- Integrated cdFMC/CDO APIs: The cdFMC REST APIs can be used using the same URL as the CDO's APIs.
- User and tenant management: API endpoints to view tenant information, and add, remove, and modify
  user roles using the CDO API.
- Change log Monitoring: API endpoints that can be used to see who has been making changes to
  devices, and to manage change requests.
- Object Management: Add, edit, delete, and view objects in CDO using these API endpoints.
