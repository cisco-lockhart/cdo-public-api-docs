# coding: utf-8

"""
    Cisco Security Cloud Control API

    Use the documentation to explore the endpoints Security Cloud Control has to offer

    The version of the OpenAPI document: 1.5.0
    Contact: cdo.tac@cisco.com
    Generated by OpenAPI Generator (https://openapi-generator.tech)

    Do not edit the class manually.
"""  # noqa: E501


from __future__ import annotations
import json
from enum import Enum
from typing_extensions import Self


class MerakiDeploymentMode(str, Enum):
    """
    (Meraki devices only) The deployment mode of the Meraki device.
    """

    """
    allowed enum values
    """
    ROUTED = 'ROUTED'
    PASSTHROUGH = 'PASSTHROUGH'

    @classmethod
    def from_json(cls, json_str: str) -> Self:
        """Create an instance of MerakiDeploymentMode from a JSON string"""
        return cls(json.loads(json_str))


