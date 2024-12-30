class HTMLNode:
    # tag represents the HTML tag name
    # value represent a string the value of the HTML
    # children represent a list of HTMLNode objects that
    # represents children of this node
    # props represents a dict representing HTML attributes
    # {"href": "https://www.google.com"}
    def __init__(self, tag, value, children=None, props=None):
        if not isinstance(tag, str):
            raise TypeError("tag must be a string")
        if not (isinstance(value, str) or value is None):
            raise TypeError("text must be a string or None")
        if children is not None and not isinstance(children, list):
            raise TypeError("children must be a list or None")
        if props is not None and not isinstance(props, dict):
            raise TypeError("props must be a dictionary or None")

        self.tag = tag
        self.value = value
        self.children = children if children is None else []
        self.props = props if props is None else {}

    def __eq__(self, other):
        if not isinstance(other, HTMLNode):
            return False
        return (
            self.tag == other.tag,
            self.value == other.value,
            self.children == other.children,
            self.props == other.props,
        )

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
