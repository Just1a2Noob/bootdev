import re

from textnode import TextNode, TextType


def split_nodes_delimiter(old_nodes, delimiter, text_type):
    if type(old_nodes) is not list:
        raise TypeError("old_nodes must be a list")

    if delimiter is None or not isinstance(delimiter, str):
        raise ValueError("delimiter must be type str and cannot be empty")
    if text_type is None or not isinstance(text_type, TextType):
        raise ValueError("text_type must be class TextType and cannot be empty")

    result = []
    for node in old_nodes:
        if not isinstance(node, TextNode):
            raise ValueError("The list must contain TextNode")

        try:
            keyword = re.search(f"{delimiter}(.+?){delimiter}", node.text).group(1)
        except Exception:
            raise Exception("Invalid markdown syntax")
        raw_split = node.text.split(delimiter)
        for i in raw_split:
            if keyword in i:
                result.append(
                    TextNode(
                        text=keyword,
                        text_type=text_type,
                    )
                )
            else:
                result.append(
                    TextNode(
                        text=i,
                        text_type=node.text_type,
                    )
                )

    return result
