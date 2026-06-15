# MspSecureClientPlatform

The OS and CPU architecture combinations to upgrade. If omitted, all platforms available for the specified version are uploaded to each device.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**cpu_architecture** | **str** | The CPU architecture. One of: ARM64, X86_64, UNIVERSAL. | 
**os** | **str** | The operating system. One of: LINUX, MACOS, WINDOWS. | 

## Example

```python
from scc_firewall_manager_sdk.models.msp_secure_client_platform import MspSecureClientPlatform

# TODO update the JSON string below
json = "{}"
# create an instance of MspSecureClientPlatform from a JSON string
msp_secure_client_platform_instance = MspSecureClientPlatform.from_json(json)
# print the JSON string representation of the object
print(MspSecureClientPlatform.to_json())

# convert the object into a dict
msp_secure_client_platform_dict = msp_secure_client_platform_instance.to_dict()
# create an instance of MspSecureClientPlatform from a dict
msp_secure_client_platform_form_dict = msp_secure_client_platform.from_dict(msp_secure_client_platform_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


