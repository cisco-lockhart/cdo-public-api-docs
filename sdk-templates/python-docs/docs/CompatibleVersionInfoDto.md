# CompatibleVersionInfoDto

Information about a compatible upgrade version and the devices that support it.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**compatible_devices** | [**List[CompatibleDeviceDto]**](CompatibleDeviceDto.md) | List of devices that are compatible with this upgrade version | 
**is_suggested_version** | **bool** | A boolean value, indicating whether this version is a suggested version to upgrade to. | 
**release_date** | **datetime** | The release date of this software version | [optional] 
**software_version** | **str** | The software version | 
**upgrade_type** | **str** | The type of the upgrade | 

## Example

```python
from scc_firewall_manager_sdk.models.compatible_version_info_dto import CompatibleVersionInfoDto

# TODO update the JSON string below
json = "{}"
# create an instance of CompatibleVersionInfoDto from a JSON string
compatible_version_info_dto_instance = CompatibleVersionInfoDto.from_json(json)
# print the JSON string representation of the object
print(CompatibleVersionInfoDto.to_json())

# convert the object into a dict
compatible_version_info_dto_dict = compatible_version_info_dto_instance.to_dict()
# create an instance of CompatibleVersionInfoDto from a dict
compatible_version_info_dto_form_dict = compatible_version_info_dto.from_dict(compatible_version_info_dto_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


