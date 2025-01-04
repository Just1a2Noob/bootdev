from enum import Enum

from htmlnode import HTMLNode, LeafNode


class TextType(Enum):
    """A class to repressent text types of TextNode.

    This class allows you to assing text types to TextNode's.
    """

    TEXT = "text"
    BOLD = "bold"
    ITALIC = "italic"
    CODE = "code"
    LINK = "link"
    IMAGE = "image"

    # Additional classes
    O_LIST = "ordered list"
    U_LIST = "unordered list"
    QUOTE = "quote"
    CODE_BLOCK = "code block"


class TextNode:
    """A class to represent different types of inline text.

    This class acts as an intermediate representation for
    parsing Markdown text, and outputting HTML.

    Attributes:
        text (str): The text content of the node
        text_type (TextType): The type of text this node contains
        url (str): The url of the link or imag, if the text is link.
                    Default is None.
    """

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
    """Converts a TextNode to LeafNode

    Args:
        text_node (TextNode): TextNode

    Returns:
        LeafNode: Changed TextNode to LeafNode with appropriate type
    """

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
