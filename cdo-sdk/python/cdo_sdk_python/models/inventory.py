# coding: utf-8

"""
    Cisco Defense Orchestrator API

    Use the interactive documentation to explore the endpoints CDO has to offer

    The version of the OpenAPI document: 0.0.1
    Contact: cdo.tac@cisco.com
    Generated by OpenAPI Generator (https://openapi-generator.tech)

    Do not edit the class manually.
"""  # noqa: E501


from __future__ import annotations
import pprint
import re  # noqa: F401
import json

from pydantic import BaseModel, ConfigDict, Field
from typing import Any, ClassVar, Dict, List, Optional
from cdo_sdk_python.models.entity import Entity
from typing import Optional, Set
from typing_extensions import Self

class Inventory(BaseModel):
    """
    Results from the CDO inventory that match the search term.
    """ # noqa: E501
    devices: Optional[List[Entity]] = Field(default=None, description="Devices that match the search term.")
    managers: Optional[List[Entity]] = Field(default=None, description="Device Managers that match the search term.")
    services: Optional[List[Entity]] = Field(default=None, description="Cloud Services that match the search term.")
    __properties: ClassVar[List[str]] = ["devices", "managers", "services"]

    model_config = ConfigDict(
        populate_by_name=True,
        validate_assignment=True,
        protected_namespaces=(),
    )


    def to_str(self) -> str:
        """Returns the string representation of the model using alias"""
        return pprint.pformat(self.model_dump(by_alias=True))

    def to_json(self) -> str:
        """Returns the JSON representation of the model using alias"""
        # TODO: pydantic v2: use .model_dump_json(by_alias=True, exclude_unset=True) instead
        return json.dumps(self.to_dict())

    @classmethod
    def from_json(cls, json_str: str) -> Optional[Self]:
        """Create an instance of Inventory from a JSON string"""
        return cls.from_dict(json.loads(json_str))

    def to_dict(self) -> Dict[str, Any]:
        """Return the dictionary representation of the model using alias.

        This has the following differences from calling pydantic's
        `self.model_dump(by_alias=True)`:

        * `None` is only added to the output dict for nullable fields that
          were set at model initialization. Other fields with value `None`
          are ignored.
        """
        excluded_fields: Set[str] = set([
        ])

        _dict = self.model_dump(
            by_alias=True,
            exclude=excluded_fields,
            exclude_none=True,
        )
        # override the default output from pydantic by calling `to_dict()` of each item in devices (list)
        _items = []
        if self.devices:
            for _item in self.devices:
                if _item:
                    _items.append(_item.to_dict())
            _dict['devices'] = _items
        # override the default output from pydantic by calling `to_dict()` of each item in managers (list)
        _items = []
        if self.managers:
            for _item in self.managers:
                if _item:
                    _items.append(_item.to_dict())
            _dict['managers'] = _items
        # override the default output from pydantic by calling `to_dict()` of each item in services (list)
        _items = []
        if self.services:
            for _item in self.services:
                if _item:
                    _items.append(_item.to_dict())
            _dict['services'] = _items
        return _dict

    @classmethod
    def from_dict(cls, obj: Optional[Dict[str, Any]]) -> Optional[Self]:
        """Create an instance of Inventory from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "devices": [Entity.from_dict(_item) for _item in obj["devices"]] if obj.get("devices") is not None else None,
            "managers": [Entity.from_dict(_item) for _item in obj["managers"]] if obj.get("managers") is not None else None,
            "services": [Entity.from_dict(_item) for _item in obj["services"]] if obj.get("services") is not None else None
        })
        return _obj

