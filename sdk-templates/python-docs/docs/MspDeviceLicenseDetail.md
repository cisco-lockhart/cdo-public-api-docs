# MspDeviceLicenseDetail


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**end_date** | **datetime** |  | [optional] 
**license_uid** | **str** |  | [optional] 
**name** | **str** |  | [optional] 
**start_date** | **datetime** |  | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.msp_device_license_detail import MspDeviceLicenseDetail

# TODO update the JSON string below
json = "{}"
# create an instance of MspDeviceLicenseDetail from a JSON string
msp_device_license_detail_instance = MspDeviceLicenseDetail.from_json(json)
# print the JSON string representation of the object
print(MspDeviceLicenseDetail.to_json())

# convert the object into a dict
msp_device_license_detail_dict = msp_device_license_detail_instance.to_dict()
# create an instance of MspDeviceLicenseDetail from a dict
msp_device_license_detail_form_dict = msp_device_license_detail.from_dict(msp_device_license_detail_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


