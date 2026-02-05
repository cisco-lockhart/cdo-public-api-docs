# PublicThousandEyesOutageDto


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**affected_domains** | **List[str]** |  | [optional] 
**alert_id** | **str** |  | [optional] 
**application_name** | **str** |  | [optional] 
**duration_seconds** | **int** |  | [optional] 
**end** | **datetime** |  | [optional] 
**errors** | **List[str]** |  | [optional] 
**outage_type** | **str** |  | [optional] 
**regions** | **List[str]** |  | [optional] 
**severity** | **str** |  | [optional] 
**start** | **datetime** |  | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.public_thousand_eyes_outage_dto import PublicThousandEyesOutageDto

# TODO update the JSON string below
json = "{}"
# create an instance of PublicThousandEyesOutageDto from a JSON string
public_thousand_eyes_outage_dto_instance = PublicThousandEyesOutageDto.from_json(json)
# print the JSON string representation of the object
print(PublicThousandEyesOutageDto.to_json())

# convert the object into a dict
public_thousand_eyes_outage_dto_dict = public_thousand_eyes_outage_dto_instance.to_dict()
# create an instance of PublicThousandEyesOutageDto from a dict
public_thousand_eyes_outage_dto_form_dict = public_thousand_eyes_outage_dto.from_dict(public_thousand_eyes_outage_dto_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


