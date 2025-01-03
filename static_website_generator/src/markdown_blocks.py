def markdown_to_blocks(markdown):
    blocks = markdown.split("\n\n")

    results = []
    for block in blocks:
        if block == "":
            continue
        block = block.strip()
        results.append(block)

    return results


document = " # This is a heading\n\n\
This is a paragraph of text. It has some **bold** and *italic* words inside of it.\n\n\
* This is the first list item in a list block\n\
* This is a list item\n\
* This is another list item"

print(markdown_to_blocks(document))
