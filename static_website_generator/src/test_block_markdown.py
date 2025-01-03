import unittest

from markdown_blocks import markdown_to_blocks


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


if __name__ == "__main__":
    unittest.main()
