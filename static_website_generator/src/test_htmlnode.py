import unittest

from htmlnode import HTMLNode


# TODO: Create some test cases for HTMLNode check for edge cases
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
        with self.assertRaises(TypeError) as cm:
            HTMLNode(123, "Valid text", [], {})
        self.assertEqual(str(cm.exception), "tag must be a string")

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


if __name__ == "__main__":
    unittest.main()
