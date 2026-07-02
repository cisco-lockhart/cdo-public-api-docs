# AsaChangeItemDto


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**deployed_config** | **str** | The configuration deployed to the device (or the configuration that was attempted to be deployed to the device, in the case of a failed deployment). Note: this field is set to null if the device already had this configuration, or when snapshot capture failed for this device. | [optional] 
**deployment_run_uid** | **str** | The unique identifier, represented as a UUID, of the deployment run the change item is associated with. | [optional] 
**deployment_type** | **str** | Discriminator identifying this change item as ASA. | [optional] 
**device_uid** | **str** | The unique identifier, represented as a UUID, of the device the change item is associated with. | [optional] 
**initial_config** | **str** | The configuration on the device prior to deployment. Note: this field is set to null if snapshot capture failed for this device. | [optional] 
**uid** | **str** | The unique identifier, represented as a UUID, of this change item. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.asa_change_item_dto import AsaChangeItemDto

# TODO update the JSON string below
json = "{}"
# create an instance of AsaChangeItemDto from a JSON string
asa_change_item_dto_instance = AsaChangeItemDto.from_json(json)
# print the JSON string representation of the object
print(AsaChangeItemDto.to_json())

# convert the object into a dict
asa_change_item_dto_dict = asa_change_item_dto_instance.to_dict()
# create an instance of AsaChangeItemDto from a dict
asa_change_item_dto_form_dict = asa_change_item_dto.from_dict(asa_change_item_dto_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


