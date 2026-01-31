# SmartAccountDto


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**parent** | **str** | The name of the smart account used to license the device. | [optional] 
**virtual** | **str** | The name of the virtual account, within the smart account, used to license the device | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.smart_account_dto import SmartAccountDto

# TODO update the JSON string below
json = "{}"
# create an instance of SmartAccountDto from a JSON string
smart_account_dto_instance = SmartAccountDto.from_json(json)
# print the JSON string representation of the object
print(SmartAccountDto.to_json())

# convert the object into a dict
smart_account_dto_dict = smart_account_dto_instance.to_dict()
# create an instance of SmartAccountDto from a dict
smart_account_dto_form_dict = smart_account_dto.from_dict(smart_account_dto_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


