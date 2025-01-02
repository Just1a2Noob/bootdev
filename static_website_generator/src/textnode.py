from enum import Enum

from htmlnode import HTMLNode, LeafNode


class TextType(Enum):
    TEXT = "text"
    BOLD = "bold"
    ITALIC = "italic"
    CODE = "code"
    LINK = "link"
    IMAGE = "image"


class TextNode:
    def __init__(self, text, text_type, url=None):
        self.text = text
        self.text_type = text_type
        self.url = url

    def __eq__(self, other):
        return (
            self.text_type == other.text_type
            and self.text == other.text
            and self.url == other.url
        )

    def __repr__(self):
        return f"TextNode({self.text}, {self.text_type}, {self.url})"


def text_node_to_html_node(text_node):
    tag_converter = {
        "text": None,
        "bold": "b",
        "italic": "i",
        "code": "code",
        "link": "a",
        "img": "img",
    }

    if isinstance(type(text_node), TextNode):
        raise TypeError("input must be of TextNode class")

    # Creating href link if there is any
    props = None
    if text_node.url is not None:
        props = {"href": text_node.url}

    result = LeafNode(
        tag=tag_converter[text_node.text_type.value],
        value=text_node.text,
        props=props,
    )
    return result
