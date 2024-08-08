# Version 1.2.0 (2024-xx-xx) 
>**Note**:
This changelog is a work in progress. Update the date and ensure it has all the necessary information before sending a publication request to Devnet.

- cdFMC health monitoring: API endpoint to monitor the health of a cloud-delivered FMC (cdFMC).
- Provisioning: Provision a cdFMC and enable Multicloud Defense on a CDO tenant.
- MSSP Tenant management: Endpoints to add user groups, provision a cdFMC, and enable Multicloud Defense on a CDO tenant that is associated with a CDO MSSP Portal.
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