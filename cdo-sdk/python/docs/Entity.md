# Entity

Cloud Services that match the search term.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**uid** | **str** | The unique identifier, represented as a UUID, of the entity in Security Cloud Control. | [optional] 
**name** | **str** | The name of the entity in CDO. Device names are unique in Security Cloud Control. | [optional] 
**type** | [**EntityType**](EntityType.md) |  | [optional] 
**address** | **str** | The address of the entity. | [optional] 
**matching_configurations** | **List[str]** | (ASAs only) Parts of the ASA device configuration that match the search term. | [optional] 

## Example

```python
from cdo_sdk_python.models.entity import Entity

# TODO update the JSON string below
json = "{}"
# create an instance of Entity from a JSON string
entity_instance = Entity.from_json(json)
# print the JSON string representation of the object
print(Entity.to_json())

# convert the object into a dict
entity_dict = entity_instance.to_dict()
# create an instance of Entity from a dict
entity_form_dict = entity.from_dict(entity_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


