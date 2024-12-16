# OnPremFmcInfo

(Devices managed by on-prem FMC only) Information on the on-prem FMC managing this device.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**uid** | **str** | The unique identifier, represented as a UUID, of the on-prem FMC that manages this device. | [optional] 
**device_record_on_fmc** | [**FmcDeviceRecord**](FmcDeviceRecord.md) |  | [optional] 
**link** | **str** | The endpoint to access this resource from. | [optional] 
**location** | **str** | The fully-qualified domain name or IP address of the on-prem FMC managing this device. | [optional] 

## Example

```python
from cdo_sdk_python.models.on_prem_fmc_info import OnPremFmcInfo

# TODO update the JSON string below
json = "{}"
# create an instance of OnPremFmcInfo from a JSON string
on_prem_fmc_info_instance = OnPremFmcInfo.from_json(json)
# print the JSON string representation of the object
print(OnPremFmcInfo.to_json())

# convert the object into a dict
on_prem_fmc_info_dict = on_prem_fmc_info_instance.to_dict()
# create an instance of OnPremFmcInfo from a dict
on_prem_fmc_info_form_dict = on_prem_fmc_info.from_dict(on_prem_fmc_info_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


