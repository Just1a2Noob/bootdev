import re

from extract_links import extract_markdown_images, extract_markdown_links
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

        # TODO: This can only find one link change it to re.findall
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


def split_nodes_image(old_nodes):
    if type(old_nodes) is not list:
        raise TypeError("old_nodes must be a list")

    result = []
    for node in old_nodes:
        if not isinstance(node, TextNode):
            raise ValueError("The list must contain TextNode")

        images = {}
        for i in extract_markdown_images(node.text):
            images[i[0]] = i[1]

        original_list = re.split(r"!\[([^\[\]]*)\]\(([^()\s]+)\)", node.text)
        max_length = len(original_list)
        for key in images:
            j = 0
            while j < max_length:
                if key in original_list[j]:
                    result.append(TextNode(key, TextType.LINK, images[key]))
                    original_list.remove(key)
                    original_list.remove(images[key])
                    max_length -= 2
                    continue
                else:
                    result.append(TextNode(original_list[j], TextType.TEXT))
                    original_list.remove(original_list[j])

    return result


def split_nodes_link(old_nodes):
    if type(old_nodes) is not list:
        raise TypeError("old_nodes must be a list")

    result = []
    for node in old_nodes:
        if not isinstance(node, TextNode):
            raise ValueError("The list must contain TextNode")

        images = {}
        for i in extract_markdown_links(node.text):
            images[i[0]] = i[1]

        original_list = re.split(r"(?<!\!)\[([^\[\]]+)\]\(([^()\s]+)\)", node.text)
        max_length = len(original_list)
        for key in images:
            j = 0
            while j < max_length:
                if key in original_list[j]:
                    result.append(TextNode(key, TextType.LINK, images[key]))
                    original_list.remove(key)
                    original_list.remove(images[key])
                    max_length -= 2
                    break
                else:
                    result.append(TextNode(original_list[j], TextType.TEXT))
                    original_list.remove(original_list[j])

    return result


# TODO: Both split functions cannot handle the below test case
# Fix both functions so that the test case below works
node = TextNode(
    "This is an image of a ![cute cat](https://www.imgur.com/cat) she is very cute",
    TextType.TEXT,
)
print(split_nodes_image([node]))
