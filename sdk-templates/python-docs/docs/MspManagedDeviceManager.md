# MspManagedDeviceManager


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**address** | **str** | The address of the device, in &lt;code&gt;host:port&lt;/code&gt; format. Security Cloud Control connects to the device at this address. | [optional] 
**connectivity_state** | [**ConnectivityState**](ConnectivityState.md) |  | [optional] 
**device_manager_type** | [**EntityType**](EntityType.md) |  | [optional] 
**fmc_domain_uid** | **str** | (FMC-only) The Domain UID of the FMC. | [optional] 
**managed_tenant_display_name** | **str** | The display name of the managed tenant in CDO. | [optional] 
**managed_tenant_name** | **str** | The name of the managed tenant in CDO. | [optional] 
**managed_tenant_region** | **str** | The region of the managed tenant in CDO. This is the region where the device manager is located. | [optional] 
**managed_tenant_uid** | **str** | The unique identifier, represented as a UUID, of the managed tenant in Security Cloud Control that this device manager belongs to. | [optional] 
**name** | **str** | The name of the device manager in CDO. Device manager names are unique in Security Cloud Control. | 
**software_version** | **str** | The version of the software running on the device. | [optional] 
**tenant_uid** | **str** |  | [optional] 
**uid** | **str** | The unique identifier, represented as a UUID, of the device manager in Security Cloud Control. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.msp_managed_device_manager import MspManagedDeviceManager

# TODO update the JSON string below
json = "{}"
# create an instance of MspManagedDeviceManager from a JSON string
msp_managed_device_manager_instance = MspManagedDeviceManager.from_json(json)
# print the JSON string representation of the object
print(MspManagedDeviceManager.to_json())

# convert the object into a dict
msp_managed_device_manager_dict = msp_managed_device_manager_instance.to_dict()
# create an instance of MspManagedDeviceManager from a dict
msp_managed_device_manager_form_dict = msp_managed_device_manager.from_dict(msp_managed_device_manager_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


