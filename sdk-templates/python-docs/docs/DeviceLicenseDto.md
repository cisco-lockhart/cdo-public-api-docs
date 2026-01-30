# DeviceLicenseDto


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**compliance_status** | **str** | Indicates whether the license used on the device is currently compliant with Smart Licensing restrictions. | [optional] 
**device_type** | [**EntityType**](EntityType.md) |  | [optional] 
**is_evaluation_license_expired** | **bool** | Indicates if the evaluation license has expired. This field is only relevant if the license status is EVALUATION. | [optional] 
**is_export_control_features_allowed** | **bool** | Indicates if the license enables the device to offer export-controlled functionality. If this is set to false, Cisco software features that are subject to national security, foreign policy, and anti-terrorism laws and regulations are disabled. | [optional] 
**license_status** | **str** | The license status of the device. | [optional] 
**licenses** | **List[str]** | The list of licenses enabled on the device. | [optional] 
**model_number** | **str** | The hardware, or virtualized hardware platform, that the device is running on (ASA-only). This field can be missing in the case of a partially onboarded device. | [optional] 
**name** | **str** | The name of the device in CDO. Device names are unique in Security Cloud Control. | [optional] 
**serial** | **str** | The serial number of the device. | [optional] 
**smart_account** | [**SmartAccount**](SmartAccount.md) |  | [optional] 
**uid** | **str** | The unique identifier, represented as a UUID, of the device in Security Cloud Control. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.device_license_dto import DeviceLicenseDto

# TODO update the JSON string below
json = "{}"
# create an instance of DeviceLicenseDto from a JSON string
device_license_dto_instance = DeviceLicenseDto.from_json(json)
# print the JSON string representation of the object
print(DeviceLicenseDto.to_json())

# convert the object into a dict
device_license_dto_dict = device_license_dto_instance.to_dict()
# create an instance of DeviceLicenseDto from a dict
device_license_dto_form_dict = device_license_dto.from_dict(device_license_dto_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


