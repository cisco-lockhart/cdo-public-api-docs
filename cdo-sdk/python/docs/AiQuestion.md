# AiQuestion


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**content** | **str** | The content of the message. | [optional] 

## Example

```python
from cdo_sdk_python.models.ai_question import AiQuestion

# TODO update the JSON string below
json = "{}"
# create an instance of AiQuestion from a JSON string
ai_question_instance = AiQuestion.from_json(json)
# print the JSON string representation of the object
print(AiQuestion.to_json())

# convert the object into a dict
ai_question_dict = ai_question_instance.to_dict()
# create an instance of AiQuestion from a dict
ai_question_form_dict = ai_question.from_dict(ai_question_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


