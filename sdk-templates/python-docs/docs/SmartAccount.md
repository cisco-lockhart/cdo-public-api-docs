# SmartAccount


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**parent** | **str** | The name of the smart account used to license the device. | [optional] 
**virtual** | **str** | The name of the virtual account, within the smart account, used to license the device | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.smart_account import SmartAccount

# TODO update the JSON string below
json = "{}"
# create an instance of SmartAccount from a JSON string
smart_account_instance = SmartAccount.from_json(json)
# print the JSON string representation of the object
print(SmartAccount.to_json())

# convert the object into a dict
smart_account_dict = smart_account_instance.to_dict()
# create an instance of SmartAccount from a dict
smart_account_form_dict = smart_account.from_dict(smart_account_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


