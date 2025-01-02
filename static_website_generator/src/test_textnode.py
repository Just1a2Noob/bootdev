import unittest

from htmlnode import LeafNode
from textnode import TextNode, TextType, text_node_to_html_node


class TestTextNode(unittest.TestCase):
    def test_textnode_with_invalid_text_type(self):
        with self.assertRaises(AttributeError) as cm:
            TextNode("This is a text node", TextType.RANDOM)
        self.assertEqual(
            str(cm.exception),
            "type object 'TextType' has no attribute 'RANDOM'",
        )

    def test_eq(self):
        node = TextNode("This is a text node", TextType.BOLD)
        node2 = TextNode("This is a text node", TextType.BOLD)
        self.assertEqual(node, node2)

        node3 = TextNode("This should be none", None)
        node4 = TextNode("This should be none", None)
        self.assertEqual(node3, node4)


class test_text_node_to_html_node(unittest.TestCase):
    def test_url_props(self):
        node = TextNode("Lorem Ipsum", TextType.ITALIC)
        node2 = LeafNode("i", "Lorem Ipsum", None)
        self.assertEqual(text_node_to_html_node(node), node2)

        node3 = TextNode("Lorem Ipsum", TextType.ITALIC, "www.google.com")
        node4 = LeafNode("i", "Lorem Ipsum", {"href": "www.google.com"})
        self.assertEqual(text_node_to_html_node(node3), node4)


if __name__ == "__main__":
    unittest.main()
