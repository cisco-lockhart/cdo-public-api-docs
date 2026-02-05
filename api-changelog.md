# Version 1.17.0 (2026-02-05)

## Added

### Device Upgrades
- Endpoint to [upgrade multiple ASA devices](https://developer.cisco.com/docs/cisco-security-cloud-control-firewall-manager/upgrade-multiple-asa-devices/) in a single request (`POST /v1/inventory/devices/asas/upgrades/trigger`).
- Endpoint to [get compatible upgrade versions for multiple ASAs](https://developer.cisco.com/docs/cisco-security-cloud-control-firewall-manager/get-upgrade-versions-compatible-with-multiple-asas/) (`GET /v1/inventory/devices/asas/upgrades/versions`).

### Inventory
- Endpoint to [export devices](https://developer.cisco.com/docs/cisco-security-cloud-control-firewall-manager/export-devices/) to CSV (`POST /v1/inventory/devices/export`).

### MSP
- Endpoint to [upgrade multiple ASA devices across managed tenants](https://developer.cisco.com/docs/cisco-security-cloud-control-firewall-manager/upgrade-multiple-asas-across-multiple-tenants/) (`POST /v1/msp/inventory/devices/asas/upgrades/trigger`).
- Endpoint to [get compatible ASA upgrade versions for MSP-managed devices](https://developer.cisco.com/docs/cisco-security-cloud-control-firewall-manager/get-compatible-asa-upgrade-versions/) (`GET /v1/msp/inventory/devices/asas/upgrades/versions`).
- Endpoint to [bulk modify MSP-managed devices](https://developer.cisco.com/docs/cisco-security-cloud-control-firewall-manager/modify-msp-managed-devices/) (`PATCH /v1/msp/inventory/devices/bulk`).

### cdFMC API
- Added AI Defense policy and profile endpoints, and access-policy AI Defense settings.
- Added DNS policy rule endpoints and identity policy category/rule endpoints.
- Added external authentication configuration and operational endpoints (LDAP/RADIUS config objects, fetch/apply).
- Added Secure Access integration endpoints (regions, tunnel configurations, transcripts).
- Added platform endpoints for content updates and device upgrade info.
- Added analysis and operational endpoints for filters, VPN tunnel status details, virtual access interfaces, and VPN tunnel status lookup.
- Added new object endpoints for SIDNS lists/downloads, SIDNS feeds, sinkholes, skip servers, realm sequences, and EVE exception lists, plus access-policy advanced logging, EVE settings, and decryption standard mode configuration.
- Active session queries now support additional filters (search, session duration, domain, user risk, device risk) and include expanded examples.
- DNS policies, identity policies, SIDNS feeds, and sinkholes now support POST/PUT/DELETE where previously only GET was available.


## Improvements

### Device Upgrades
- [Stage ASA device upgrades](https://developer.cisco.com/docs/cisco-security-cloud-control-firewall-manager/upgrade-asa-device/) to get the latest version on the device without performing the actual upgrade.
- Upgrade compatibility and upgrade run responses now include error details when retrieval or execution fails.

### MSP
- [MSP-managed tenant](https://developer.cisco.com/docs/cisco-security-cloud-control-firewall-manager/mspmanagedtenant/) records now include cdFMC type, tenant pay type, and SCC organization UID.

### Object Management
- Object issue handling now supports ignore flags for duplicates and unused issues.

## Deprecations

- The JWKS endpoint is now marked deprecated (`GET /.well-known/jwks.json`).
- The MSP FTD upgrade runs endpoint is now deprecated; use `GET /v1/msp/inventory/devices/upgrades/runs` instead of `GET /v1/msp/inventory/devices/ftds/upgrades/runs`.
- The cdFMC AIOps configure endpoint was removed from the spec (`/v1/cdfmc/api/fmc_config/v1/domain/{domainUUID}/integration/aiops/configure`).

# Version 1.16.0 (2025-11-14)

## Added

### Events
- Endpoint to [get event search reports](https://developer.cisco.com/docs/cisco-security-cloud-control-firewall-manager/get-event-search-reports/) containing the results of event log searches executed using the Report feature in the Event Logging page.

### Health Metrics
- Endpoint to get [device health metric aggregations](https://developer.cisco.com/docs/cisco-security-cloud-control-firewall-manager/get-device-health-metric-aggregations/) for all managed devices, supporting CPU, memory, and disk metrics.
- Endpoint to get [detailed device health metrics](https://developer.cisco.com/docs/cisco-security-cloud-control-firewall-manager/get-device-detailed-list-for-aggregations/ to retrieve filtered device lists corresponding to aggregation metrics with threshold filtering (CRITICAL, WARNING, OK).
- Endpoint to get [interface health metric aggregations](https://developer.cisco.com/docs/cisco-security-cloud-control-firewall-manager/get-interface-health-metric-aggregations/) for all managed devices, supporting link status, overruns, underruns, dropped packets, L2 decode drops, and input/output errors.
- Endpoint to get [interface health metric error aggregations](https://developer.cisco.com/docs/cisco-security-cloud-control-firewall-manager/get-interface-health-metric-error-aggregations/) providing summary of interface errors across all managed devices.

### Device Upgrades
- Endpoint to [get device upgrade runs](https://developer.cisco.com/docs/cisco-security-cloud-control-firewall-manager/get-device-upgrade-runs/) to list all FTD device upgrade runs in the tenant.
- Endpoint to [get a device upgrade run](https://developer.cisco.com/docs/cisco-security-cloud-control-firewall-manager/get-device-upgrade-run/) by UID with detailed completion status for each device.
- Endpoint to [modify a device upgrade run](https://developer.cisco.com/docs/cisco-security-cloud-control-firewall-manager/modify-device-upgrade-run/) to update upgrade run properties.
- Endpoint to [delete a device upgrade run](https://developer.cisco.com/docs/cisco-security-cloud-control-firewall-manager/delete-device-upgrade-run/) by UID.

### MSP
- Endpoint to [calculate compatible upgrade versions for multiple MSP-managed devices](https://developer.cisco.com/docs/cisco-security-cloud-control-firewall-manager/calculate-compatible-ftd-upgrade-versions/) across different tenants (up to 50 devices).
- Endpoint to [upgrade multiple MSP-managed FTD devices](https://developer.cisco.com/docs/cisco-security-cloud-control-firewall-manager/upgrade-multiple-ftds-across-multiple-tenants/) across different tenants simultaneously.
- Endpoint to [get MSP upgrade runs](https://developer.cisco.com/docs/cisco-security-cloud-control-firewall-manager/get-msp-device-upgrade-runs/) to monitor upgrade operations across managed tenants.
- Endpoint to [get a specific MSP upgrade run](https://developer.cisco.com/docs/cisco-security-cloud-control-firewall-manager/get-msp-device-upgrade-run/) by UID.
- Endpoint to [get MSP upgrade run attribute values](https://developer.cisco.com/docs/cisco-security-cloud-control-firewall-manager/get-distinct-attribute-values-for-msp-upgrade-runs/) to retrieve distinct values for filtering upgrade runs.

## Improvements

### Device Upgrades
- Device upgrade operations now support maintenance window configuration. You can specify `ignoreMaintenanceWindow` parameter when upgrading [single FTD devices](https://developer.cisco.com/docs/cisco-security-cloud-control-firewall-manager/upgrade-ftd-device/) and [multiple FTD devices](https://developer.cisco.com/docs/cisco-security-cloud-control-firewall-manager/upgrade-multiple-ftd-devices/).

### MSP
- The [Get MSP-managed Devices](https://developer.cisco.com/docs/cisco-security-cloud-control-firewall-manager/get-msp-managed-devices/) endpoint now provides additional filterable fields including `hardwareModels` and `modelNumbers` for better device inventory management.

### Object Management
- The [Delete Object](https://developer.cisco.com/docs/cisco-security-cloud-control-firewall-manager/delete-object/) endpoint now supports a `forcedDelete` parameter to delete objects that are in use by other objects but are not associated with a device.

### API Consistency
- The [Get Users](https://developer.cisco.com/docs/cisco-security-cloud-control-firewall-manager/get-users/) endpoint now has a maximum limit of 50 results (reduced from 200) for better performance.

### Authentication & Tenants
- The authentication token response now includes `tenantName` field providing the name of the enterprise (tenant) in Security Cloud Control.
- Tenant settings now include `objectSharingEnabled` field to indicate if Object Sharing is enabled for the tenant.

### High Availability
- HA node information now includes `role` field for cdFMC-managed FTDs, indicating whether a node is ACTIVE or STANDBY in an HA pair.

### Transactions
- Improved transaction field descriptions for better clarity across all transaction-related responses.

# Version 1.15.0 (2025-09-05)

## Added

- *21* new endpoints
  to [view and manage ASA interfaces](https://developer.cisco.com/docs/cisco-security-cloud-control-firewall-manager/get-asa-etherchannel-interfaces/).

## Improvements

- Significantly improved and expanded examples in
  the [Getting Started](https://developer.cisco.com/docs/cisco-security-cloud-control-firewall-manager/getting-started/)
  section.
- You can choose to stage an upgrade instead of applying the upgrade, when upgrading
  a [single FTD device](https://developer.cisco.com/docs/cisco-security-cloud-control-firewall-manager/upgrade-ftd-device/)
  or [multiple FTD devices](https://developer.cisco.com/docs/cisco-security-cloud-control-firewall-manager/upgrade-multiple-ftd-devices/)
  using the `stageUpgrade` parameter.
    - If `stageUpgrade` is set to true, the image will be downloaded on to the device and readiness
      checks will be performed. However, the upgrade will not be applied to the device.
- The [Get MSP-managed Devices API](https://developer.cisco.com/docs/cisco-security-cloud-control-firewall-manager/get-msp-managed-devices/) endpoint
now provides additional information on HA devices and FTD clusters.
- The [Get Devices](https://developer.cisco.com/docs/cisco-security-cloud-control-firewall-manager/get-devices/) and [Get MSP-managed Devices](https://developer.cisco.com/docs/cisco-security-cloud-control-firewall-manager/get-msp-managed-devices/) endpoints now also provide additional information on certificate expiry dates.

# Version 1.14.0 (2025-07-25)

## Added

- Multiple new API endpoints to collect and monitor health metrics for ASA devices.
    - Endpoints
      to [enable](https://developer.cisco.com/docs/cisco-security-cloud-control/opt-in-to-asa-health-metrics/)
      and [disable](https://developer.cisco.com/docs/cisco-security-cloud-control/opt-out-of-asa-health-metrics/)
      ASA health metrics collection.
    - Endpoint to
      get [device health metrics](https://developer.cisco.com/docs/cisco-security-cloud-control/get-time-series-health-metrics-for-one-or-more-asa-devices/) (
      such as CPU usage, memory usage etc), and
      an [interface metrics](https://developer.cisco.com/docs/cisco-security-cloud-control/get-time-series-interface-metrics-for-an-asa-device/)
      for ASAs across multiple time ranges.
- Multiple new API endpoints to query and export MSSP portal inventory.
    - Endpoints to get
      managed [devices](https://developer.cisco.com/docs/cisco-security-cloud-control/get-msp-managed-devices/), [device managers](https://developer.cisco.com/docs/cisco-security-cloud-control/get-msp-managed-device-managers/), [templates](https://developer.cisco.com/docs/cisco-security-cloud-control/get-msp-managed-templates/),
      and [cloud services](https://developer.cisco.com/docs/cisco-security-cloud-control/get-msp-managed-cloud-services/)
      across all tenants in an MSSP portal.
    - Endpoints to get an individual
      managed [device](https://developer.cisco.com/docs/cisco-security-cloud-control/get-msp-managed-device-by-uid/), [device manager](https://developer.cisco.com/docs/cisco-security-cloud-control/get-msp-managed-device-manager-by-uid/), [template](https://developer.cisco.com/docs/cisco-security-cloud-control/get-msp-managed-template-by-uid/),
      or [cloud service](https://developer.cisco.com/docs/cisco-security-cloud-control/get-msp-managed-cloud-service-by-uid/)
      across any of the tenants managed by an MSSP portal.
    - Endpoints to
      export [managed devices](https://developer.cisco.com/docs/cisco-security-cloud-control/export-msp-managed-devices/), [device managers](https://developer.cisco.com/docs/cisco-security-cloud-control/export-msp-managed-device-managers/), [templates](https://developer.cisco.com/docs/cisco-security-cloud-control/export-msp-managed-templates/),
      and [cloud services](https://developer.cisco.com/docs/cisco-security-cloud-control/export-msp-managed-cloud-services/)
      across all tenants in an MSSP portal to CSV.
- Endpoint
  to [get upgrade versions](https://developer.cisco.com/docs/cisco-security-cloud-control/get-upgrade-versions-compatible-with-multiple-ftds/)
  compatible with multiple FTDs.

# Version 1.13.0 (2025-05-23)

## Added

- New endpoint
  to [deploy changes to up to 50 FTDs simultaneously](https://developer.cisco.com/docs/cisco-security-cloud-control/cdfmc-managed-ftds-only-deploy-changes-to-multiple-ftd-devices/).

## Improvements

- You can now upgrade up to 50 FTDs simultaneously using
  the [Upgrade multiple FTD devices](https://developer.cisco.com/docs/cisco-security-cloud-control/upgrade-multiple-ftd-devices/)
  endpoint.
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
