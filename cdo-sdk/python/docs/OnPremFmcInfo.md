# OnPremFmcInfo

(Devices managed by on-prem FMC only) Information on the on-prem FMC managing this device.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **str** | The name of the device in CDO. Device names are unique in CDO. | [optional] 
**address** | **str** | The address of the on-prem FMC managing this device, in &#x60;host:port&#x60; format. | [optional] 

## Example

```python
from openapi_client.models.on_prem_fmc_info import OnPremFmcInfo

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


