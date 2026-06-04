# FtdSingleDeviceDeploymentInput

Per-device deployment configuration. Cannot be combined with the deprecated `deviceUids` field. Each entry identifies a device and optionally selects which policy types to deploy. Set `selectedPolicyTypes` to null or include FULL_DEPLOY to perform a full deployment for that device. Each entry must have a unique uid.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**selected_policy_types** | **List[str]** | (Optional) The policy types to deploy to this device. Use this to only deploy selected policies to the device, while skipping deployment of other parts of the device&#39;s configuration on SCC Firewall Management. Use individual policy types or group shortcuts (ACCESS_GROUP, UZTNA_GROUP). Set to null, or include FULL_DEPLOY, to perform a full deployment for this device. | [optional] 
**uid** | **str** | The unique identifier, represented as a UUID, of the device to deploy to. | 

## Example

```python
from scc_firewall_manager_sdk.models.ftd_single_device_deployment_input import FtdSingleDeviceDeploymentInput

# TODO update the JSON string below
json = "{}"
# create an instance of FtdSingleDeviceDeploymentInput from a JSON string
ftd_single_device_deployment_input_instance = FtdSingleDeviceDeploymentInput.from_json(json)
# print the JSON string representation of the object
print(FtdSingleDeviceDeploymentInput.to_json())

# convert the object into a dict
ftd_single_device_deployment_input_dict = ftd_single_device_deployment_input_instance.to_dict()
# create an instance of FtdSingleDeviceDeploymentInput from a dict
ftd_single_device_deployment_input_form_dict = ftd_single_device_deployment_input.from_dict(ftd_single_device_deployment_input_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


