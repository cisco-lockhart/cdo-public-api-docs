# Inventory


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**devices** | [**List[Entity]**](Entity.md) | Devices that match the search term. | [optional] 
**managers** | [**List[Entity]**](Entity.md) | Device Managers that match the search term. | [optional] 
**services** | [**List[Entity]**](Entity.md) | Cloud Services that match the search term. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.inventory import Inventory

# TODO update the JSON string below
json = "{}"
# create an instance of Inventory from a JSON string
inventory_instance = Inventory.from_json(json)
# print the JSON string representation of the object
print(Inventory.to_json())

# convert the object into a dict
inventory_dict = inventory_instance.to_dict()
# create an instance of Inventory from a dict
inventory_form_dict = inventory.from_dict(inventory_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


