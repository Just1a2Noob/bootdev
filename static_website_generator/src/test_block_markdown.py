import unittest

from markdown_blocks import block_to_block_type, markdown_to_blocks


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


if __name__ == "__main__":
    unittest.main()
