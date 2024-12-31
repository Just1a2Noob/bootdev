import unittest

from htmlnode import HTMLNode, LeafNode, ParentNode


class TestHTMLNode(unittest.TestCase):
    def test_HTMLNode(self):
        node = HTMLNode(
            "p",
            "This is a paragaph",
            [[HTMLNode("b", "bold text", [], {})]],
            {
                "href": "https://www.google.com",
                "target": "_blank",
            },
        )
        node2 = HTMLNode(
            "p",
            "This is a paragaph",
            [[HTMLNode("b", "bold text", [], {})]],
            {
                "href": "https://www.google.com",
                "target": "_blank",
            },
        )
        self.assertEqual(node, node2)

        node3 = HTMLNode("h1", "This is a header", None, None)
        node4 = HTMLNode("h1", "This is a header", None, None)
        self.assertEqual(node3, node4)

    def test_HTMLNode_None(self):
        node = HTMLNode(
            "p",
            "The others should be empty",
            None,
            None,
        )
        node2 = HTMLNode(
            "p",
            "The others should be empty",
            None,
            None,
        )
        self.assertEqual(node, node2)

    def test_HTMLNode_invalid_tag(self):
        with self.assertRaises(TypeError):
            HTMLNode(123, "Content")

    def test_HTMLNode_invalid_text(self):
        with self.assertRaises(TypeError) as cm:
            HTMLNode("p", 123, [], {})
        self.assertEqual(str(cm.exception), "text must be a string or None")

    def test_HTMLNode_invalid_attributes(self):
        with self.assertRaises(TypeError) as cm:
            HTMLNode("p", "Valid text", [], [])
        self.assertEqual(str(cm.exception), "props must be a dictionary or None")

    def test_HTMLNode_incomplete_data(self):
        node = HTMLNode("p", "Some text should be here")
        node2 = HTMLNode("p", "Some text should be here")
        self.assertEqual(node, node2)


class TestLeafNode(unittest.TestCase):
    def test_LeafNode_to_html(self):
        node = LeafNode("p", "This is a paragraph of text.")
        node2 = "<p>This is a paragraph of text.</p>"
        self.assertEqual(node.to_html(), node2)

        node3 = LeafNode("a", "Click me!", {"href": "https://www.google.com"})
        node4 = """<a href="https://www.google.com">Click me!</a>"""
        self.assertEqual(node3.to_html(), node4)

    def test_LeafNode_invalid_empty_value(self):
        with self.assertRaises(TypeError) as cm:
            LeafNode("p")
        self.assertEqual(
            str(cm.exception),
            "LeafNode.__init__() missing 1 required positional argument: 'value'",
        )

    def test_LeafNode_invalid_empty_string(self):
        with self.assertRaises(ValueError) as cm:
            LeafNode("p", "")
        self.assertEqual(str(cm.exception), "value cannot be empty")


class test_ParentNode(unittest.TestCase):
    def test_ParentNode_to_html(self):
        node = ParentNode(
            "p",
            [
                LeafNode("b", "Bold text"),
                LeafNode(None, "Normal text"),
                LeafNode("i", "italic text"),
                LeafNode(None, "Normal text"),
            ],
        )
        node2 = "<p><b>Bold text</b>Normal text<i>italic text</i>Normal text</p>"
        self.assertEqual(node.to_html(), node2)

    def test_ParentNode_invalid_children(self):
        with self.assertRaises(TypeError) as cm:
            ParentNode("p", "")
        self.assertEqual(str(cm.exception), "children must be list or None")

    def test_ParentNode_invalid_tag(self):
        with self.assertRaises(ValueError) as cm:
            ParentNode(None, [LeafNode("b", "Bold text")])
        self.assertEqual(str(cm.exception), "tag cannot be empty")


if __name__ == "__main__":
    unittest.main()
