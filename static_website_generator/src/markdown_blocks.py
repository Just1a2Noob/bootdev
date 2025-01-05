import re

from htmlnode import HTMLNode, LeafNode, ParentNode
from inline_markdown import text_to_textnodes
from textnode import TextNode, TextType, text_node_to_html_node


def markdown_to_blocks(markdown):
    """Split a markdown document into block.

    Args:
        markdown (str): A document that uses markdown syntax.

    Returns:
        list: A list containing the blocks from the markdown document.
    """
    blocks = markdown.split("\n\n")

    results = []
    for block in blocks:
        if block == "":
            continue
        block = block.strip()
        results.append(block)

    return results


def block_to_block_type(block):
    """Determines the block type of a block.

    Args:
        block (str): A string/block of markdown text.

    Returns:
        str: A string representing the type of block it is.
    """
    header_pattern = r"^#"
    code_pattern = r"^```[\s\S]*```$"
    qoute_pattern = r"^>"
    unordered_list_pattern = r"^[*-]\s"
    ordered_list_pattern = r"^\d+\."

    patterns = [
        (header_pattern, "header"),
        (code_pattern, "code"),
        (qoute_pattern, "quote"),
        (unordered_list_pattern, "unordered_list"),
        (ordered_list_pattern, "ordered_list"),
    ]

    found = 0
    for pattern, block_type in patterns:
        if re.search(pattern, block):
            found = 1
            return block_type

    if found == 0:
        return "paragraph"


def text_to_children(text):
    """Converts any text to a list of LeafNode's.

    Args:
        text (str): A string that uses markdown syntax.

    Returns:
        list: A list containing LeafNode's from the text
                with the appropriate type.
    """
    text_nodes = text_to_textnodes(text)

    children = []
    for node in text_nodes:
        children.append(text_node_to_html_node(node))

    return children


def text_to_list(text, list_type):
    """Converts a text to a list of ParentNode with tag 'li'.

    Args:
        text (str): A string containing markdown syntax.
        list_type (str): A string representing the type
                        of list is being used.

    Returns:
        list: A list of ParentNode's with the tag 'li'.
    """

    # if list_type is invalid pattern is empty
    pattern = ""
    if list_type == "unordered_list":
        pattern = r"^[*-]\s(.*)$"
    if list_type == "ordered_list":
        pattern = r"^\d+\. +(.+)$"

    # re.MULTILINE  flag ensures that the ^ anchor
    # matches the start of every line, not just the start of the entire string.
    items = re.findall(pattern, text, flags=re.MULTILINE)

    results = []
    for item in items:
        children = text_to_children(item)
        if len(children) > 1:
            results.append(ParentNode(tag="li", children=children))
        else:
            results.append(LeafNode(tag="li", value=item))

    return results


def text_to_code(text):
    """Converts text to ParentNode with code block tags.

    Args:
        text (str): A string with markdown syntax.

    Returns:
        ParentNode: A ParentNode with the 'pre' tag and
                    within that ParentNode contains another ParentNode
                    with 'code' tag.
    """
    return ParentNode(tag="pre", children=[LeafNode(tag="```", value=text)])


def markdown_to_htmlnode(markdown):
    """Converts a full markdown document into a single parent HTMLNode.

    Args:
        markdown (str): A string in a format of markdown document.

    Returns:
        HTMLNode: A single HTMLNode contaning many 'children'.
    """
    md_blocks = markdown_to_blocks(markdown)

    block_type_to_tag = {
        "paragraph": "p",
        "code": "```",
        "unordered_list": "ul",
        "ordered_list": "ol",
        "quote": "blockquote",
    }

    nodes = []
    for block in md_blocks:
        text_type = block_to_block_type(block)

        if text_type == "code":
            nodes.append(text_to_code(block))

        if text_type == "unordered_list" or text_type == "ordered_list":
            list_node = text_to_list(block, text_type)
            nodes.append(
                ParentNode(tag=block_type_to_tag[text_type], children=list_node)
            )

        # If its not code/list type it just appends
        else:
            children = text_to_children(block)

            if len(children) > 1:
                nodes.append(
                    ParentNode(tag=block_type_to_tag[text_type], children=children)
                )
            else:
                nodes.append(LeafNode(tag=block_type_to_tag[text_type], value=block))

    results = HTMLNode(tag="div", value=None, children=nodes, props=None)

    return results


node = """
This is **bolded** paragraph

```python
print(hello_world)
```


This is another paragraph with *italic* text and `code` here
This is the same paragraph on a new line

* This is a list
* with items
"""
# print(markdown_to_htmlnode(node))
