import unittest

from textnode import TextNode, TextType


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


if __name__ == "__main__":
    unittest.main()
