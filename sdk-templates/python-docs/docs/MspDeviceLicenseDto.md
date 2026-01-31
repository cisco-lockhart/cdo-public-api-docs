# MspDeviceLicenseDto


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**compliance_status** | **str** | Indicates whether the license used on the device is currently compliant with Smart Licensing restrictions. | [optional] 
**device_type** | [**EntityType**](EntityType.md) |  | [optional] 
**is_evaluation_license_expired** | **bool** | Indicates if the evaluation license has expired. This field is only relevant if the license status is EVALUATION. | [optional] 
**is_export_control_features_allowed** | **bool** | Indicates if export control features are allowed on the device. | [optional] 
**last_updated_time** | **datetime** | The time, in RFC 3339 format, that the license usage for this device was updated. | [optional] 
**license_status** | **str** | The license status of the device. | [optional] 
**licenses** | [**List[MspDeviceLicenseDetail]**](MspDeviceLicenseDetail.md) | The list of licenses enabled on the device. | [optional] 
**managed_tenant_uid** | **str** | The unique identifier, represented as a UUID, of the managed tenant associated with the device. | [optional] 
**model_number** | **str** | The hardware, or virtualized hardware platform, that the device is running on (ASA-only). This field can be missing in the case of a partially onboarded device. | [optional] 
**name** | **str** | The name of the device in CDO. Device names are unique in Security Cloud Control. | [optional] 
**serial** | **str** | The serial number of the device. | [optional] 
**smart_account** | [**SmartAccountDto**](SmartAccountDto.md) |  | [optional] 
**uid** | **str** | The unique identifier, represented as a UUID, of the device in Security Cloud Control. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.msp_device_license_dto import MspDeviceLicenseDto

# TODO update the JSON string below
json = "{}"
# create an instance of MspDeviceLicenseDto from a JSON string
msp_device_license_dto_instance = MspDeviceLicenseDto.from_json(json)
# print the JSON string representation of the object
print(MspDeviceLicenseDto.to_json())

# convert the object into a dict
msp_device_license_dto_dict = msp_device_license_dto_instance.to_dict()
# create an instance of MspDeviceLicenseDto from a dict
msp_device_license_dto_form_dict = msp_device_license_dto.from_dict(msp_device_license_dto_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


