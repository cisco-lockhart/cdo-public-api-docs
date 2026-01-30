# SCC Firewall Manager SDK

Python SDK for the Cisco Security Cloud Control Firewall Manager API.

## Overview

Use this SDK to interact with the Security Cloud Control Firewall Manager API endpoints
programmatically.

## Requirements

- Python 3.8+

## Installation

```bash
pip install scc-firewall-manager-sdk
```

## Quick Start

### Authentication

The SDK uses Bearer token (JWT) authentication. You can obtain a token from the Security Cloud
Control console.

```python
from scc_firewall_manager_sdk import ApiClient, Configuration
import os

configuration = Configuration(
    host="https://api.us.security.cisco.com/firewall",
    access_token=os.environ["BEARER_TOKEN"]
)
```

### Regional Endpoints

| Region | Host                                          |
|--------|-----------------------------------------------|
| US     | `https://api.us.security.cisco.com/firewall`  |
| EU     | `https://api.eu.security.cisco.com/firewall`  |
| APJ    | `https://api.apj.security.cisco.com/firewall` |
| AU     | `https://api.au.security.cisco.com/firewall`  |
| IN     | `https://api.in.security.cisco.com/firewall`  |
| UAE    | `https://api.uae.security.cisco.com/firewall` |

### Example Usage

#### List Devices

```python
from scc_firewall_manager_sdk import Configuration, ApiClient, InventoryApi
import os

configuration = Configuration(
    host="https://api.us.security.cisco.com/firewall", # Replace `us` with your region
    access_token=os.environ["BEARER_TOKEN"]
)

with ApiClient(configuration) as api_client:
    inventory_api = InventoryApi(api_client)
    devices = inventory_api.get_devices()
    for device in devices.items:
        print(f"Device: {device.name} - {device.device_type}")
```

### API Reference

See the sidebar for documentation on all available API endpoints and models. This SDK, and the
Python docs are auto-generated. To see the original OpenAPI spec,
visit [Cisco Devnet](https://developer.cisco.com/docs/cisco-security-cloud-control-firewall-manager/introduction/).

### Support

For issues and questions, contact [cdo.tac@cisco.com](mailto:cdo.tac@cisco.com).
