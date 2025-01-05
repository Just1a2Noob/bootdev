import unittest

from htmlnode import HTMLNode, LeafNode, ParentNode
from markdown_blocks import (
    block_to_block_type,
    markdown_to_blocks,
    markdown_to_htmlnode,
    text_to_children,
    text_to_code,
    text_to_list,
)


class test_markdown_to_blocks(unittest.TestCase):
    # BUG: I shouldn't have to do setUp but if I dont use it
    # It doesn't works
    def setUp(self):
        self.markdown_to_blocks = markdown_to_blocks

    def test_basic_markdown_structure(self):
        node = """
This is **bolded** paragraph

This is another paragraph with *italic* text and `code` here
This is the same paragraph on a new line

* This is a list
* with items
"""
        expected_output = [
            "This is **bolded** paragraph",
            "This is another paragraph with *italic* text and `code` here\nThis is the same paragraph on a new line",
            "* This is a list\n* with items",
        ]
        self.assertEqual(self.markdown_to_blocks(node), expected_output)

    def test_empty_string(self):
        self.assertEqual(self.markdown_to_blocks(""), [])

    def test_single_block(self):
        node = "This is a single block of text"
        expected_output = ["This is a single block of text"]
        self.assertEqual(self.markdown_to_blocks(node), expected_output)

    def test_multiple_empty_lines(self):
        node = """
This is **bolded** paragraph




This is another paragraph with *italic* text and `code` here
This is the same paragraph on a new line

* This is a list
* with items
"""
        expected_output = [
            "This is **bolded** paragraph",
            "This is another paragraph with *italic* text and `code` here\nThis is the same paragraph on a new line",
            "* This is a list\n* with items",
        ]
        self.assertEqual(self.markdown_to_blocks(node), expected_output)

    def test_whitespace_handling(self):
        node = "  Block with spaces  \n\n   Another block   "
        expected_output = ["Block with spaces", "Another block"]
        self.assertEqual(self.markdown_to_blocks(node), expected_output)


class test_block_to_block_type(unittest.TestCase):
    def test_header_block(self):
        node = "# This should be a header"
        expected_output = "header"
        self.assertEqual(block_to_block_type(node), expected_output)

    def test_code_block(self):
        node = "```print('hello world')```"
        expected_output = "code"
        self.assertEqual(block_to_block_type(node), expected_output)

    def test_unordered_list(self):
        node1 = "* List1\n* List2\n* List3"
        expected_output1 = "unordered_list"
        self.assertEqual(block_to_block_type(node1), expected_output1)

        node2 = "- List1\n- List2\n- List3"
        expected_output2 = "unordered_list"
        self.assertEqual(block_to_block_type(node2), expected_output2)

    def test_ordered_list(self):
        node = "1. List1\n2. List2\n3. List3"
        expected_output = "ordered_list"
        self.assertEqual(block_to_block_type(node), expected_output)

    def test_qoute_block(self):
        node = ">[!NOTE]\n>Very important note."
        expected_output = "quote"
        self.assertEqual(block_to_block_type(node), expected_output)

    def test_basic_text(self):
        node = "This should just be a **normal** *text*"
        expected_output = "paragraph"
        self.assertEqual(block_to_block_type(node), expected_output)

    def test_unclosed_delimiter(self):
        node = "```print('hello world')``"
        expected_output = "paragraph"
        self.assertEqual(block_to_block_type(node), expected_output)


class test_text_to_children(unittest.TestCase):
    def test_inline_input(self):
        node = "This **should** be some text.\nAnd this is another line with *another* text."
        expected_output = [
            LeafNode(None, "This "),
            LeafNode("b", "should"),
            LeafNode(None, "be some text\nAnd this is another line with "),
            LeafNode("i", "another"),
            LeafNode(None, " text."),
        ]
        self.assertEqual(text_to_children(node), expected_output)

    def test_paragraph_input(self):
        node = "This is just a normal paragraph with nothing in it."
        expected_output = [
            LeafNode("p", "This is just a normal paragraph with nothing in it")
        ]
        self.assertEqual(text_to_children(node), expected_output)


class test_text_to_list(unittest.TestCase):
    def test_ordered_list(self):
        node = "1. list1\n2. list2\n3. list3"
        expected_output = [
            LeafNode("li", "list1"),
            LeafNode("li", "list2"),
            LeafNode("li", "list3"),
        ]
        self.assertEqual(text_to_list(node, "ordered_list"), expected_output)

    def test_unorderd_list(self):
        node1 = "- list1\n- list2\n- list3"
        expected_output1 = [
            LeafNode("li", "list1"),
            LeafNode("li", "list2"),
            LeafNode("li", "list3"),
        ]
        self.assertEqual(text_to_list(node1, "unordered_list"), expected_output1)

        node2 = "* list1\n* list2\n* list3"
        expected_output2 = [
            LeafNode("li", "list1"),
            LeafNode("li", "list2"),
            LeafNode("li", "list3"),
        ]
        self.assertEqual(text_to_list(node2, "unordered_list"), expected_output2)

    def test_mixed_unordered_list(self):
        node = "- list1\n* list2\n- list3"
        expected_output = []
        self.assertEqual(text_to_list(node, "ordered_list"), expected_output)

    def test_unordered_list_with_bold_italic(self):
        node = "* **bolded** text\n* text of an *italic*\n* Maybe a *mix* of **both**"
        expected_output = [
            ParentNode(
                "li",
                [
                    LeafNode("b", "bolded"),
                    LeafNode(None, " text"),
                ],
            ),
            ParentNode(
                "li",
                [
                    LeafNode(None, "text of an "),
                    LeafNode("i", "italic"),
                ],
            ),
            ParentNode(
                "li",
                [
                    LeafNode(None, "Maybe a "),
                    LeafNode("i", "mix"),
                    LeafNode(None, " of "),
                    LeafNode("b", "both"),
                ],
            ),
        ]
        self.assertEqual(text_to_list(node, "unordered_list"), expected_output)


class test_text_to_code(unittest.TestCase):
    def test_valid_input(self):
        node = "```python\nprint('hello world')\n# This should print hello world```"
        expected_output = ParentNode(
            "pre",
            [
                LeafNode(
                    "```",
                    "python\nprint('hello world')\n# This should print hello world",
                )
            ],
        )
        self.assertEqual(text_to_code(node), expected_output)


class test_markdown_to_htmlnode(unittest.TestCase):
    def test_valid_input(self):
        node = """
        This **code** block, should explain how you should use functions to print hello world

        ```python
        text = 'Hello World!'
        print(text)
        ```

        In this code there are a few things going on:

        1. We assigned the variable named 'text' with a str,
        2. In that *string* it contains `Hello World`,
        3. We use print to show the **output** of variable text.

        > Assigning the string to a variable isn't necessary. You can just directly print('Hello World!')
        """

        expected_output = HTMLNode(
            "div",
            None,
            [
                ParentNode(
                    "p",
                    [
                        LeafNode(None, "This "),
                        LeafNode("b", "code"),
                        LeafNode(
                            None,
                            " block, should explain how you should use functions to print hello world",
                        ),
                    ],
                ),
                ParentNode(
                    "pre",
                    [LeafNode("code", "python\ntext = 'Hello World'\nprint(text)")],
                ),
                LeafNode("p", "In this code there are a few things going on:"),
                ParentNode(
                    "ol",
                    [
                        LeafNode(
                            "li", "We assigned the variable named 'text' with astr,"
                        ),
                        ParentNode(
                            "li",
                            [
                                LeafNode(None, "In that "),
                                LeafNode("b", "string"),
                                LeafNode(None, " it contains "),
                                LeafNode("code", "Hello World"),
                                LeafNode(None, ","),
                            ],
                        ),
                        ParentNode(
                            "li",
                            [
                                LeafNode(None, "We use print to show the "),
                                LeafNode("b", "output"),
                                LeafNode(None, " of variable text."),
                            ],
                        ),
                    ],
                ),
                LeafNode(
                    "blockquote",
                    "Assigning the string to a variable isn't necessary. YOu can just directly print('Hello World!)",
                ),
            ],
        )

        self.assertEqual(markdown_to_htmlnode(node), expected_output)


if __name__ == "__main__":
    unittest.main()
