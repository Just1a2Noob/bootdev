import re


def extract_markdown_images(text):
    try:
        matches = re.findall(r"!\[([^\[\]]*)\]\(([^()\s]+)\)", text)
        if not matches:
            raise ValueError("Cannot find markdown image links or malformed syntax")
    except Exception as e:
        raise ValueError(f"Error extracting markdown images: {e}")
    return matches


def extract_markdown_links(text):
    try:
        matches = re.findall(r"(?<!\!)\[([^\[\]]+)\]\(([^()\s]+)\)", text)
        if not matches:
            raise ValueError("Cannot find markdown links or malformed syntax")
    except Exception as e:
        raise ValueError(f"Error extracting markdown links: {e}")
    return matches
