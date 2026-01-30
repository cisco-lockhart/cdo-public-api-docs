# Changelog


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**entity_name** | **str** | The name of the device, manager or service to which the Change Log refers. | [optional] 
**entity_uid** | **str** | The unique identifier, represented as a UUID, of the device, manager or service to which the Change Log refers. | [optional] 
**events** | [**List[Event]**](Event.md) | The events recorded in this Change Log. | [optional] 
**last_event_date** | **datetime** | The time (UTC; represented using the RFC-3339 standard) at which the last event in the changelog occurred. You can perform range queries on this endpoint using the syntax &#x60;[lastEventDate:YYYY-MM-ddTHH:mm:ssZ TO lastEventDate:YYYY-MM-ddTHH:mm:ssZ]&#x60;: please note that this will need to be URL-encoded. Please note that the timeRange request parameter is mapped to the lastEventDate field. For further details on the timeRange parameter, refer to the Changelogs API documentation. | [optional] 
**status** | **str** | The status of the Change Log. | [optional] 
**uid** | **str** | The unique identifier, represented as a UUID, of the Change Log. | 

## Example

```python
from scc_firewall_manager_sdk.models.changelog import Changelog

# TODO update the JSON string below
json = "{}"
# create an instance of Changelog from a JSON string
changelog_instance = Changelog.from_json(json)
# print the JSON string representation of the object
print(Changelog.to_json())

# convert the object into a dict
changelog_dict = changelog_instance.to_dict()
# create an instance of Changelog from a dict
changelog_form_dict = changelog.from_dict(changelog_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


