import unittest

from split_delimiter import split_nodes_delimiter, split_nodes_image, split_nodes_link
from textnode import TextNode, TextType


class test_split_delimiter(unittest.TestCase):
    # TODO: Add 1 more test with multiple delimiters
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


class test_extract_link(unittest.TestCase):
    def test_normality(self):
        node = TextNode(
            "This is text with a link [to boot dev](https://www.boot.dev) and [to youtube](https://www.youtube.com/@bootdotdev)",
            TextType.TEXT,
        )
        node2 = [
            TextNode("This is text with a link ", TextType.TEXT),
            TextNode("to boot dev", TextType.LINK, "https://www.boot.dev"),
            TextNode(" and ", TextType.TEXT),
            TextNode(
                "to youtube", TextType.LINK, "https://www.youtube.com/@bootdotdev"
            ),
        ]
        self.assertEqual(split_nodes_link([node]), node2)

    def test_invalid_text_node(self):
        with self.assertRaises(ValueError) as cm:
            split_nodes_link(
                ["This is a youtube [link](https://www.youtube.com)", TextType.TEXT]
            )
        self.assertEqual(str(cm.exception), "The list must contain TextNode")

    def test_invalid_old_nodes(self):
        with self.assertRaises(TypeError) as cm:
            split_nodes_link(
                {"link": "https://www.youtube.com", "text_type": TextType.TEXT}
            )
        self.assertEqual(str(cm.exception), "old_nodes must be a list")


class test_extract_images(unittest.TestCase):
    def test_normality(self):
        node = TextNode(
            "This is an image of a ![cute cat](https://www.imgur.com/cat) she is very cute",
            TextType.TEXT,
        )
        node2 = [
            TextNode("This is an image of a ", TextType.TEXT),
            TextNode("cute cat", TextType.LINK, "https://www.imgur.com/cat"),
            TextNode(" she is very cute", TextType.TEXT),
        ]
        self.assertEqual(split_nodes_image([node]), node2)


if __name__ == "__main__":
    unittest.main()
