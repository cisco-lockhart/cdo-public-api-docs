# DeviceLicenseDetail


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**end_date** | **datetime** |  | [optional] 
**license_uid** | **str** |  | [optional] 
**name** | **str** |  | [optional] 
**start_date** | **datetime** |  | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.device_license_detail import DeviceLicenseDetail

# TODO update the JSON string below
json = "{}"
# create an instance of DeviceLicenseDetail from a JSON string
device_license_detail_instance = DeviceLicenseDetail.from_json(json)
# print the JSON string representation of the object
print(DeviceLicenseDetail.to_json())

# convert the object into a dict
device_license_detail_dict = device_license_detail_instance.to_dict()
# create an instance of DeviceLicenseDetail from a dict
device_license_detail_form_dict = device_license_detail.from_dict(device_license_detail_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


