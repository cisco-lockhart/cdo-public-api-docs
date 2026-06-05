# OnPremFmcInfo


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**link** | **str** | The endpoint to access this resource from. | [optional] 
**location** | **str** | The fully-qualified domain name or IP address of the on-prem FMC managing this device. | [optional] 
**uid** | **str** | The unique identifier, represented as a UUID, of the on-prem FMC that manages this device. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.on_prem_fmc_info import OnPremFmcInfo

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


