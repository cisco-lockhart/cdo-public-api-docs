# CompatibleVersionInfo

Information about a compatible ASA upgrade version and the devices that support it.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**asdm_image_url** | **str** | The ASDM image URL | [optional] 
**asdm_version** | **str** | The ASDM version | 
**compatible_devices** | [**List[CompatibleDevice]**](CompatibleDevice.md) | List of devices that are compatible with this upgrade version | 
**software_image_url** | **str** | The ASA software image URL | [optional] 
**software_version** | **str** | The software version | 

## Example

```python
from scc_firewall_manager_sdk.models.compatible_version_info import CompatibleVersionInfo

# TODO update the JSON string below
json = "{}"
# create an instance of CompatibleVersionInfo from a JSON string
compatible_version_info_instance = CompatibleVersionInfo.from_json(json)
# print the JSON string representation of the object
print(CompatibleVersionInfo.to_json())

# convert the object into a dict
compatible_version_info_dict = compatible_version_info_instance.to_dict()
# create an instance of CompatibleVersionInfo from a dict
compatible_version_info_form_dict = compatible_version_info.from_dict(compatible_version_info_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


