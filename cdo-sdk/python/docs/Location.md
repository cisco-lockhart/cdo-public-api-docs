# Location

The location information of the device associated with the RA VPN session.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**city** | **str** | The city where the device associated with the RA VPN session is located. | [optional] 
**subdivision** | **str** | The subdivision (e.g., state or province) where the device is located. | [optional] 
**country** | **str** | The country where the device associated with the RA VPN session is located. | [optional] 

## Example

```python
from cdo_python_sdk.models.location import Location

# TODO update the JSON string below
json = "{}"
# create an instance of Location from a JSON string
location_instance = Location.from_json(json)
# print the JSON string representation of the object
print(Location.to_json())

# convert the object into a dict
location_dict = location_instance.to_dict()
# create an instance of Location from a dict
location_form_dict = location.from_dict(location_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


