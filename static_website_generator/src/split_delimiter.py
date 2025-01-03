import re

from extract_links import extract_markdown_images, extract_markdown_links
from textnode import TextNode, TextType


def split_nodes_delimiter(old_nodes, delimiter, text_type):
    if not isinstance(old_nodes, list):
        raise TypeError("old_nodes must be a list")

    if delimiter is None or not isinstance(delimiter, str):
        raise ValueError("delimiter must be type str and cannot be empty")
    if text_type is None:
        raise ValueError("text_type cannot be empty")

    if not isinstance(text_type, TextType):
        raise TypeError("text_type must be of class TextType")

    results = []
    for node in old_nodes:
        if not isinstance(node, TextNode):
            raise TypeError("Elements in old_nodes must be of type TextNode")

        split_nodes = []
        sections = node.text.split(delimiter)
        if len(sections) % 2 == 0:
            raise ValueError("Invalid markdown syntax")
        for i in range(len(sections)):
            if sections[i] == "":
                continue
            if i % 2 == 0:
                split_nodes.append(TextNode(sections[i], TextType.TEXT))
            else:
                split_nodes.append(TextNode(sections[i], text_type))
        results.extend(split_nodes)
    return results


def split_nodes_image(old_nodes):
    if not isinstance(old_nodes, list):
        raise ValueError("Input must be type list")

    results = []
    for node in old_nodes:
        original_text = node.text
        matches = extract_markdown_images(original_text)
        if len(matches) == 0:
            results.append(node)
            continue
        for alt, link in matches:
            sections = original_text.split(f"![{alt}]({link})", 1)

            if len(sections) != 2:
                raise ValueError("Invalid markdown, image selection is not closed")

            if sections[0] != "":
                results.append(TextNode(sections[0], TextType.TEXT))

            results.append(TextNode(alt, TextType.IMAGE, link))

            # The statement below changes the original text
            # Meaning for each loop it reduces the problem
            original_text = sections[1]

    # appends the remaining text at the end
    if original_text != "":
        results.append(TextNode(original_text, TextType.TEXT))

    return results


def split_nodes_link(old_nodes):
    if not isinstance(old_nodes, list):
        raise ValueError("Input must be type list")

    results = []
    for node in old_nodes:

        if not isinstance(node, TextNode):
            raise TypeError("The list must only contain TextNode")

        original_text = node.text
        matches = extract_markdown_links(original_text)
        if len(matches) == 0:
            results.append(node)
            continue
        for alt, link in matches:
            sections = original_text.split(f"[{alt}]({link})", 1)

            if len(sections) != 2:
                raise ValueError("Invalid markdown, image selection is not closed")

            if sections[0] != "":
                results.append(TextNode(sections[0], TextType.TEXT))

            results.append(TextNode(alt, TextType.LINK, link))

            # The statement below changes the original text
            # Meaning for each loop it reduces the problem
            original_text = sections[1]

    # appends the remaining text at the end
    if original_text != "":
        results.append(TextNode(original_text, TextType.TEXT))

    return results
