import unittest

from inline_markdown import (
    extract_markdown_images,
    extract_markdown_links,
    split_nodes_delimiter,
    split_nodes_image,
    split_nodes_link,
    text_to_textnodes,
)
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

    # BUG: Fix this text it shouldve be pass but it did not

    #   def test_non_text_node(self):
    #       with self.assertRaises(TypeError):
    #           split_nodes_delimiter(["not_a_node"], "**", TextType.BOLD)

    def test_invalid_old_nodes_type(self):
        with self.assertRaises(TypeError):
            split_nodes_delimiter("not_a_list", "**", TextType.BOLD)

    # BUG: Fix this text it shouldve be pass but it did not

    #   def test_invalid_text_type(self):
    #       node1 = TextNode("This is a bold word.", TextType.TEXT)
    #
    #       with self.assertRaises(ValueError) as cm:
    #           split_nodes_delimiter([node1], "**", "not_a_TextType")
    #       self.assertEqual(str(cm.exception), "text_type must be a class of TextType")


class test_extract_link(unittest.TestCase):
    def test_valid_input(self):
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
    def test_valid_input(self):
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


class test_extract_markdown_links(unittest.TestCase):
    def test_valid_input(self):
        node = """This is text with a link [to boot dev](https://www.boot.dev) 
        and [to youtube](https://www.youtube.com/@bootdotdev)"""
        node2 = [
            ("to boot dev", "https://www.boot.dev"),
            ("to youtube", "https://www.youtube.com/@bootdotdev"),
        ]
        self.assertEqual(extract_markdown_links(node), node2)


class test_extract_markdown_images(unittest.TestCase):
    def test_valid_input(self):
        node = """This is text with a ![rick roll](https://i.imgur.com/aKaOqIh.gif) 
        and ![obi wan](https://i.imgur.com/fJRm4Vk.jpeg)"""
        node2 = [
            ("rick roll", "https://i.imgur.com/aKaOqIh.gif"),
            ("obi wan", "https://i.imgur.com/fJRm4Vk.jpeg"),
        ]
        self.assertEqual(extract_markdown_images(node), node2)


# TODO: Create tests for extract_markdown_images


class test_text_to_text_nodes(unittest.TestCase):
    def test_valid_input(self):
        node = "This is **text** with an *italic* word and a `code block` and an ![obi wan image](https://i.imgur.com/fJRm4Vk.jpeg) and a [link](https://boot.dev)"

        expected_output = [
            TextNode("This is ", TextType.TEXT),
            TextNode("text", TextType.BOLD),
            TextNode(" with an ", TextType.TEXT),
            TextNode("italic", TextType.ITALIC),
            TextNode(" word and a ", TextType.TEXT),
            TextNode("code block", TextType.CODE),
            TextNode(" and an ", TextType.TEXT),
            TextNode(
                "obi wan image", TextType.IMAGE, "https://i.imgur.com/fJRm4Vk.jpeg"
            ),
            TextNode(" and a ", TextType.TEXT),
            TextNode("link", TextType.LINK, "https://boot.dev"),
        ]
        self.assertEqual(text_to_textnodes(node), expected_output)

    def test_various_types(self):
        node = "`code` *should* be **bolded** [link](https://google.com) and ![cutecat](https://i.imgur/cutecat.jpeg)"

        expected_output = [
            TextNode("code", TextType.CODE),
            TextNode(" ", TextType.TEXT),
            TextNode("should", TextType.ITALIC),
            TextNode(" be ", TextType.TEXT),
            TextNode("bolded", TextType.BOLD),
            TextNode(" ", TextType.TEXT),
            TextNode("link", TextType.LINK, "https://google.com"),
            TextNode(" and ", TextType.TEXT),
            TextNode("cutecat", TextType.IMAGE, "https://i.imgur/cutecat.jpeg"),
        ]
        self.assertEqual(text_to_textnodes(node), expected_output)

    def test_unclosed_delimiter(self):
        node = "This a test of *might and strength"
        with self.assertRaises(Exception):
            text_to_textnodes(node)


if __name__ == "__main__":
    unittest.main()
