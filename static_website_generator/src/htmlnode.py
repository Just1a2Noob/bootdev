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
        self.children = children
        self.props = props

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
        for key, value in self.props.items():
            result += f'{key}="{value}" '
        return result.strip()


class LeafNode(HTMLNode):
    def __init__(self, tag, value, props=None):
        if not value:  # Ensure value is not empty
            raise ValueError("value cannot be empty")
        if value is str:
            if len(value) < 1:
                raise ValueError("value cannot be empty")

        super().__init__(tag, value, props=props)

    def __eq__(self, other):
        return (
            self.tag == other.tag,
            self.value == other.value,
            self.props == other.props,
        )

    def to_html(self):
        props = super().props_to_html()

        if self.tag is None:
            return self.value

        if props is None:
            return f"<{self.tag}>{self.value}</{self.tag}>"
        return f"<{self.tag} {props}>{self.value}</{self.tag}>"


class ParentNode(HTMLNode):
    def __init__(self, tag, children, props=None):
        # Check if tag is None
        if self.tag is None:
            raise ValueError("tag cannot be empty")
        # Check if children is none
        if self.children is None:
            raise ValueError("children cannot be empty")

        # Check if tag is just an empty string
        if self.tag is str:
            if self.len(self.tag) < 1:
                raise ValueError("tag cannot be empty")
        # Check if children is an empty list
        if self.children is list:
            if len(self.children) < 1:
                raise ValueError("children cannot be empty")
        super().__init__(tag, children, props)

    def to_html(self):
        result = ""
        for child in self.children:
            if child is LeafNode:
                result += child.to_html()
            else:
                raise ValueError("children inside list has to be type LeafNode")

        return f"<{self.tag}>{result}</{self.tag}>"
