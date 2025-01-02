import unittest

from split_delimiter import split_nodes_delimiter
from textnode import TextNode, TextType


class test_split_delimiter(unittest.TestCase):
    def test_normality(self):
        node1 = split_nodes_delimiter(
            [
                TextNode("This is text with a `code block` word", TextType.TEXT),
            ],
            "`",
            TextType.CODE,
        )
        node2 = [
            TextNode("This is text with a ", TextType.TEXT),
            TextNode("code block", TextType.CODE),
            TextNode(" word", TextType.TEXT),
        ]
        self.assertEqual(node1, node2)

    def test_syntax_error(self):
        with self.assertRaises(Exception) as cm:
            split_nodes_delimiter(
                [TextNode("Lorem *Ipsum lorek", TextType.TEXT)],
                "*",
                TextType.ITALIC,
            )
        self.assertEqual(str(cm.exception), "Invalid markdown syntax")

    def test_invalid_delimiter(self):
        with self.assertRaises(ValueError) as cm:
            split_nodes_delimiter(
                [TextNode("Lorem Ipsum lorek", TextType.TEXT)],
                [],
                TextType.ITALIC,
            )
        self.assertEqual(
            str(cm.exception), "delimiter must be type str and cannot be empty"
        )

    def test_invalid_text_type(self):
        with self.assertRaises(ValueError) as cm:
            split_nodes_delimiter(
                [TextNode("Lorem Ipsum lorek", TextType.TEXT)],
                "*",
                "italic",
            )
        self.assertEqual(
            str(cm.exception), "text_type must be class TextType and cannot be empty"
        )

    def test_invalid_old_nodes(self):
        with self.assertRaises(ValueError) as cm:
            split_nodes_delimiter(
                ["Lorem ipsum lorek", TextType.TEXT],
                "*",
                TextType.ITALIC,
            )
        self.assertEqual(str(cm.exception), "The list must contain TextNode")


if __name__ == "__main__":
    unittest.main()
