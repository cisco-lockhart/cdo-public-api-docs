# Version 1.5.0 (2024-10-08)

## Added
- MSSP Tenant management: 
  - Endpoint to list the tenants managed by an MSSP portal.
  - Endpoint to generate the API token for an API-only user in a tenant managed by an MSSP portal.

## Improvements
- MSSP Tenant management: It is now possible to create API-only users in a tenant managed by the MSSP portal.
- Changelogs: It is now possible to filter changelogs by date, and also retrieve changelogs over the last 5, 15, 30, and 60 minutes.
- ASA device management: The API now returns the chassis serial number of secondary devices in ASA HA Pairs.
- cdFMC Health Monitoring: 
  - The rate limit for the endpoint has been raised from 1 request per minute to 2 requests per minute.
  - The health monitoring endpoint returns device HA status and Power Supply Unit (PSU) information.

# Version 1.4.0 (2024-09-12)

## Added
- MSSP tenant management: Endpoint to provision cloud-delivered FMC (cdFMC) on a CDO tenant that is associated with a CDO MSSP Portal.
- Inventory management: Endpoint to provision a FTD using the Zero-Touch Provisioning (ZTP) workflow.

# Version 1.3.0 (2024-08-22) 

## Added
- cdFMC health monitoring: API endpoint to monitor the health of a cloud-delivered FMC (cdFMC).
- Provisioning: Provision a cdFMC and enable Multicloud Defense on a CDO tenant.
- MSSP Tenant management: Endpoints to add user groups, and enable Multicloud Defense on a CDO tenant that is associated with a CDO MSSP Portal.
- Improvements to the devices endpoint: The **Get Devices** endpoint now provides information on the redundancy mode of a firewall, and provides additional High Availability (HA) and clustering information for FTD HA Pairs and clusters.

# Version 1.2.0 (2024-07-19)

## Added
- ASA Access Groups and Access Rules management: API endpoints to manage ASA Access Groups and Access Rules.
- MSSP Tenant management: API endpoints for MSSP portal super-admin users to create a new tenant and add it to the CDO MSSP portal and add existing tenants to the portal.
- Active Directory (AD) groups management: API endpoints to create, remove, modify, list, and fetch Active Directory groups on a CDO tenant.

# Version 1.1.0 (2024-05-30)

## Added
- AI Assistant: API endpoints to start a conversation with the AI assistant, ask questions, and list conversations and messages.

# Version 1.0.0 (2024-05-23)

## Added
- Secure Device Connector (SDC) management: API endpoints to create and delete Secure Device Connectors (SDCs).
- ASA Command Line Interface (CLI): API endpoints to execute arbitrary CLIs and CLI macros across one or more devices (ASAs only).
- RA VPN session management: API endpoints to view, refresh, and terminate RA VPN sessions on RA VPN headend devices.
- MFA event monitoring: API endpoints to view MFA events on Duo Admin Panels.
- Object Management: Object Management endpoints are now accessible from `v1` endpoints like the rest of CDO's endpoints.

# Version 0.1.0 (2024-04-02)

## Added

This is the first major release of the CDO API, and exposes a large number of CDO features. Some of the highlights include:
- Simplified Inventory management: API endpoints to add, delete, view, and perform operations on devices, device managers, and cloud services available on the CDO inventory page.
- Integrated cdFMC/CDO APIs: The cdFMC REST APIs can be used using the same URL as the CDO's APIs.
- User and tenant management: API endpoints to view tenant information, and add, remove, and modify user roles using the CDO API.
- Change log Monitoring: API endpoints that can be used to see who has been making changes to devices, and to manage change requests.
- Object Management: Add, edit, delete, and view objects in CDO using these API endpoints.