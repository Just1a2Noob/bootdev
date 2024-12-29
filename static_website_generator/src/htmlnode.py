class HTMLNode:
    # tag represents the HTML tag name
    # value represent a string string the value of the HTML
    # children represent a list of HTMLNode objects that
    # represents children of this node
    # props represents a dict representing HTML attributes
    # {"href": "https://www.google.com"}
    def __init__(self, tag, value, children, props):
        self.tag = tag
        self.value = value
        self.children = children
        self.props = props

    def to_html(self):
        raise NotImplementedError

    def __repr__(self):
        return f"HTMLNode({self.tag}, {self.value}, {self.children}, {self.props})"

    def props_to_html(self):
        if self.props is None:
            return None

        result = ""
        for i in self.props:
            result += i + "=" + self.props[i] + " "

        return result


dummy = HTMLNode(
    "<p>",
    "This is a paragraph test",
    [],
    {
        "href": "https://www.google.com",
        "target": "_blank",
    },
)

print(dummy.__repr__())
