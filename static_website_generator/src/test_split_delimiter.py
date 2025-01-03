import unittest

from split_delimiter import split_nodes_delimiter, split_nodes_image, split_nodes_link
from textnode import TextNode, TextType


class test_split_delimiter(unittest.TestCase):
    def test_valid_input(self):
        node1 = TextNode("This is a **bold** word.", TextType.TEXT)
        expected_output = [
            TextNode("This is a ", TextType.TEXT),
            TextNode("bold", TextType.BOLD),
            TextNode(" word.", TextType.TEXT),
        ]
        self.assertEqual(
            split_nodes_delimiter([node1], "**", TextType.BOLD), expected_output
        )

    def test_multiple_nodes(self):
        node1 = TextNode("First **bold** word.", TextType.TEXT)
        node2 = TextNode("Second **bold** word.", TextType.TEXT)
        expected_output = [
            TextNode("First ", TextType.TEXT),
            TextNode("bold", TextType.BOLD),
            TextNode(" word.", TextType.TEXT),
            TextNode("Second ", TextType.TEXT),
            TextNode("bold", TextType.BOLD),
            TextNode(" word.", TextType.TEXT),
        ]
        self.assertEqual(
            split_nodes_delimiter([node1, node2], "**", TextType.BOLD), expected_output
        )

    def test_empty_string(self):
        node1 = TextNode("**", TextType.TEXT)
        with self.assertRaises(ValueError):
            split_nodes_delimiter([node1], "**", TextType.BOLD)

    def test_non_text_node(self):
        with self.assertRaises(TypeError):
            split_nodes_delimiter(["not_a_node"], "**", TextType.BOLD)

    def test_invalid_old_nodes_type(self):
        with self.assertRaises(TypeError):
            split_nodes_delimiter("not_a_list", "**", TextType.BOLD)

    def test_invalid_delimiter(self):
        node1 = TextNode("This is a bold word.", TextType.TEXT)
        with self.assertRaises(ValueError):
            split_nodes_delimiter([node1], None, TextType.BOLD)
        with self.assertRaises(ValueError):
            split_nodes_delimiter([node1], 123, TextType.BOLD)

    def test_invalid_text_type(self):
        node1 = TextNode("This is a bold word.", TextType.TEXT)
        with self.assertRaises(ValueError):
            split_nodes_delimiter([node1], "**", None)
        with self.assertRaises(TypeError):
            split_nodes_delimiter([node1], "**", "not_a_TextType")


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
        with self.assertRaises(TypeError) as cm:
            split_nodes_link(
                ["This is a youtube [link](https://www.youtube.com)", TextType.TEXT]
            )
        self.assertEqual(str(cm.exception), "The list must only contain TextNode")

    def test_invalid_old_nodes(self):
        with self.assertRaises(ValueError) as cm:
            split_nodes_link(
                {"link": "https://www.youtube.com", "text_type": TextType.TEXT}
            )
        self.assertEqual(str(cm.exception), "Input must be type list")


class test_extract_images(unittest.TestCase):
    def test_normality(self):
        node = TextNode(
            "This is an image of a ![cute cat](https://www.imgur.com/cat) she is very cute",
            TextType.TEXT,
        )
        node2 = [
            TextNode("This is an image of a ", TextType.TEXT),
            TextNode("cute cat", TextType.IMAGE, "https://www.imgur.com/cat"),
            TextNode(" she is very cute", TextType.TEXT),
        ]
        self.assertEqual(split_nodes_image([node]), node2)


if __name__ == "__main__":
    unittest.main()
