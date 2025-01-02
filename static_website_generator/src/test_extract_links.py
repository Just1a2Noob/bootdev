import unittest

from extract_links import extract_markdown_images, extract_markdown_links


class test_extract_markdown_links(unittest.TestCase):
    def test_normality(self):
        node = """This is text with a link [to boot dev](https://www.boot.dev) 
        and [to youtube](https://www.youtube.com/@bootdotdev)"""
        node2 = [
            ("to boot dev", "https://www.boot.dev"),
            ("to youtube", "https://www.youtube.com/@bootdotdev"),
        ]
        self.assertEqual(extract_markdown_links(node), node2)

    def test_invalid_pattern(self):
        with self.assertRaises(ValueError) as cm:
            extract_markdown_links("Link with some [text](https://www.boot.dev")
        self.assertEqual(
            str(cm.exception),
            "Error extracting markdown links: Cannot find markdown links or malformed syntax",
        )


class test_extract_markdown_images(unittest.TestCase):
    def test_normality(self):
        node = """This is text with a ![rick roll](https://i.imgur.com/aKaOqIh.gif) 
        and ![obi wan](https://i.imgur.com/fJRm4Vk.jpeg)"""
        node2 = [
            ("rick roll", "https://i.imgur.com/aKaOqIh.gif"),
            ("obi wan", "https://i.imgur.com/fJRm4Vk.jpeg"),
        ]
        self.assertEqual(extract_markdown_images(node), node2)

    def test_invalid_pattern(self):
        with self.assertRaises(ValueError) as cm:
            extract_markdown_images(
                "This is a [cat](https://i.imgur.com/cat.gif) image"
            )
        self.assertEqual(
            str(cm.exception),
            "Error extracting markdown images: Cannot find markdown image links or malformed syntax",
        )


if __name__ == "__main__":
    unittest.main()
