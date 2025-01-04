import re


def markdown_to_blocks(markdown):
    blocks = markdown.split("\n\n")

    results = []
    for block in blocks:
        if block == "":
            continue
        block = block.strip()
        results.append(block)

    return results


def block_to_block_type(block):
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
