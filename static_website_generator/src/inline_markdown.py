import re

from textnode import TextNode, TextType


def split_nodes_delimiter(old_nodes, delimiter, text_type):
    """Split TextNode's based on delimiter and text_type.

    Args:
        old_nodes (list): The list of TextNode's.
        delimiter (str): The delimiter for markdown syntax.
        text_type (TextType): The text type of TextNode.

    Returns:
        list: The list of nodes split based on delimiter.
    """
    if not isinstance(old_nodes, list):
        raise TypeError("old_nodes must be type of list")

    if not isinstance(text_type, TextType):
        raise ValueError("text_type must be a class of TextType")

    results = []
    for node in old_nodes:
        if not isinstance(node, TextNode):
            raise ValueError("nodes must only contain TextNode")

        if node.text_type != TextType.TEXT:
            results.append(node)
            continue

        split_nodes = []
        sections = node.text.split(delimiter)
        if len(sections) % 2 == 0:
            raise ValueError("Invalid markdown, bold section not closed")
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
    """Converts a list of TextNode containing alt and url
        to TextNode(alt, TextType.IMAGE, url).

    Args:
        old_nodes (list): List containing TextNode's.

    Return:
        list: The list of TextNode's and
            any TextNode containing alt and url are changed.
    """
    if not isinstance(old_nodes, list):
        raise ValueError("Input must be type list")

    results = []
    for node in old_nodes:
        # Check if node is TextNode
        if not isinstance(node, TextNode):
            raise TypeError("The list must only contain TextNode")

        if node.text_type != TextType.TEXT:
            results.append(node)
            continue

        original_text = node.text
        matches = extract_markdown_images(original_text)
        if len(matches) == 0:
            results.append(node)
            continue

        current_position = 0
        for alt, link in matches:
            image_markup = f"![{alt}]({link})"
            start_index = original_text.find(image_markup, current_position)

            # Add any text before the image as a TEXT node
            if start_index > current_position:
                prefix_text = original_text[current_position:start_index]
                if prefix_text:  # Only check if text exists
                    results.append(TextNode(prefix_text, TextType.TEXT))

            # Add the image node
            results.append(TextNode(alt, TextType.IMAGE, link))

            current_position = start_index + len(image_markup)

        # Add any remaining text after the last image
        if current_position < len(original_text):
            remaining_text = original_text[current_position:]
            if remaining_text:  # Only check if text exists
                results.append(TextNode(remaining_text, TextType.TEXT))

    return results


def split_nodes_link(old_nodes):
    """Converts a list of TextNode containing alt and url
        to TextNode(alt, TextType.LINK, url).

    Args:
        old_nodes (list): List containing TextNode's.

    Return:
        list: The list of TextNode's and
            any TextNode containing alt and url are changed.
    """
    if not isinstance(old_nodes, list):
        raise ValueError("Input must be type list")

    results = []
    for node in old_nodes:

        # Check if node is TextNode
        if not isinstance(node, TextNode):
            raise TypeError("The list must only contain TextNode")

        if node.text_type != TextType.TEXT:
            results.append(node)
            continue

        original_text = node.text
        matches = extract_markdown_links(original_text)
        if len(matches) == 0:
            results.append(node)
            continue

        current_position = 0
        for alt, link in matches:
            link_markup = f"[{alt}]({link})"
            start_index = original_text.find(link_markup, current_position)

            # Add any text before the link as a TEXT node
            if start_index > current_position:
                prefix_text = original_text[current_position:start_index]
                if prefix_text:  # Only check if text exists
                    results.append(TextNode(prefix_text, TextType.TEXT))

            # Add the link node
            results.append(TextNode(alt, TextType.LINK, link))

            current_position = start_index + len(link_markup)

        # Add any remaining text after the last link
        if current_position < len(original_text):
            remaining_text = original_text[current_position:]
            if remaining_text:  # Only check if text exists
                results.append(TextNode(remaining_text, TextType.TEXT))

    return results


def extract_markdown_images(text):
    """Finds all markdown images in text

    Args:
        text (str): string containing markdown image syntax

    Return:
        list: list containing tuple of alt and link
    """
    matches = re.findall(r"!\[([^\[\]]*)\]\(([^()\s]+)\)", text)
    return matches


def extract_markdown_links(text):
    """Finds all markdown links in text.

    Args:
        text (str): string containing markdown links syntax.

    Return:
        list: list containing tuple of alt and link.
    """
    matches = re.findall(r"(?<!\!)\[([^\[\]]+)\]\(([^()\s]+)\)", text)
    return matches


def text_to_textnodes(text):
    """Converts text to TextNodes.

    Args:
        text (str): String with markdown syntax.

    Return:
        list: List of TextNode's, where the TextType is
            based on markdown syntax of the string.
    """
    nodes = [TextNode(text, TextType.TEXT)]
    nodes = split_nodes_delimiter(nodes, "**", TextType.BOLD)
    nodes = split_nodes_delimiter(nodes, "*", TextType.ITALIC)
    nodes = split_nodes_delimiter(nodes, "`", TextType.CODE)
    nodes = split_nodes_image(nodes)
    nodes = split_nodes_link(nodes)

    return nodes
