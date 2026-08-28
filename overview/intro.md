<seotitle>Cloud Firewall Manager, Firewall Cloud Manager, Security Policy, Network Management, Deploy, Upgrade, Security Cloud Control, Cisco Security Cloud Control, Cisco Security Cloud Control API, Cisco Secure Firewall, SCC Firewall Manager, Security Cloud Control Firewall Manager SDK, Security Cloud Control Firewall Manager CLI, Security Cloud Control Firewall Manager Terraform Provider, Security Cloud Control Firewall Manager Ansible Collection</seotitle>

# Cisco Security Cloud Control Firewall Manager API Documentation

Cisco Security Cloud Control Firewall Manager exposes a rich REST API. This document describes the semantics of the REST API. Backwards compatibility of the API is guaranteed using a versioning system.

## What can you do with it?
- Manage your devices, services, and device managers
- Deploy changes to devices at scale
- Manage your objects
- Perform complex searches across Security Cloud Control
- Execute queries across Security Cloud Control and cloud-delivered Firewall Management Center (cdFMC)
- Monitor Remote Access Virtual Private Network (RA VPN) sessions and Multi-factor Authentication (MFA) events
- Monitor your changelog
- Execute commands across your entire fleet of devices
- Build your own dashboard as a Managed Service Provider

## Automating Security Cloud Control Firewall Manager

You do not have to call the API directly. We build and support a Python SDK, a CLI, an Ansible collection, a Terraform provider, and an agent plugin for Claude Code and Codex on top of these endpoints. They are all listed on [Automation Tools and AI Agents](/docs/cisco-security-cloud-control-firewall-manager/automation-tools-and-ai-agents/).

## FedRAMP Availability

All Security Cloud Control Firewall Manager APIs are available on FedRAMP (https://manage.secure.cisco/api/rest) and IL5 (https://manage.securitydod.cisco/api/rest), with the exception of the MSP APIs.