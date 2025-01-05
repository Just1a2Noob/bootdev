class HTMLNode:
    """A class representing a 'node' in a HTML document

    The HTMLNode class will represent a "node" in an HTML document tree
    (like a <p> tag and its contents, or an <a> tag and its contents)
    and is purpose-built to render itself as HTML.

    Attributes:
        tag (str): A string representing the HTML tag name (e.g."p", "h1", etc).
        value (str): A string representing the value of the HTML tag.
        children (list): A list of HTMLNode objects representing the children
                            of this node.
        props (dict): A dictionary of key-value pairs representing the attributes
                        of HTML tag.
    """

    def __init__(self, tag, value, children=None, props=None):
        if not (isinstance(tag, str) or tag is None):
            raise TypeError("tag must be a string or None")
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
            return ""
        props_html = ""
        for prop in self.props:
            props_html += f' {prop}="{self.props[prop]}"'
        return props_html


class LeafNode(HTMLNode):
    """A class to represent a single HTML tag with no children
        with the type of HTMLNode

    Attributes:
        tag (str): A string representing the LeafNode tag name (e.g."p", "h1", etc).
        value (str): A string representing the value of the HTML tag.
        props (dict): A dictionary of key-value pairs representing the attributes
                        of LeafNode tag.
    """

    def __init__(self, tag, value, props=None):
        if value is None:  # Ensure value is not empty
            raise ValueError("value cannot be empty")

        super().__init__(tag, value, props=props)

    def __repr__(self):
        return f"LeafNode({self.tag}, {self.value}, {self.props})"

    def __eq__(self, other):
        return (
            self.tag == other.tag,
            self.value == other.value,
            self.props == other.props,
        )

    def to_html(self):
        if self.value is None:
            raise ValueError("Invalid HTML: no value")
        if self.tag is None:
            return self.value
        return f"<{self.tag}{self.props_to_html()}>{self.value}</{self.tag}>"


class ParentNode(HTMLNode):
    """A class to represent the nesting of HTML nodes inside one another

    This class is used for any HTML node that isn't a 'leaf' node
    (i.e. it has children) is a 'parent' node.

    Attributes:
        tag (str): A string representing the HTML tag name (e.g."p", "h1", etc).
        children (list): A list of HTMLNode objects representing the children
                            of this node.
        props (dict): A dictionary of key-value pairs representing the attributes
                        of ParentNode tag.
    """

    def __init__(self, tag, children, props=None):
        # Check if tag is None
        if tag is None:
            raise ValueError("tag cannot be empty")
        # Check if children is none
        if children is None:
            raise ValueError("children cannot be empty")

        # Check if tag is just an empty string
        if tag is str:
            if self.len(tag) < 1:
                raise ValueError("tag cannot be empty")

        # Check if children is an empty list
        if children is not None and not isinstance(children, list):
            raise TypeError("children must be list or None")
        else:
            if len(children) < 1:
                raise ValueError("children cannot be empty")

        super().__init__(tag=tag, value=None, children=children, props=props)

    def __repr__(self):
        return f"ParentNode({self.tag}, {self.children}), {self.props}"

    def to_html(self):
        if self.tag is None:
            raise ValueError("Invalid HTML: no tag")
        if self.children is None:
            raise ValueError("Invalid HTML: no children")
        children_html = ""
        for child in self.children:
            children_html += child.to_html()

        return f"<{self.tag}{self.props_to_html()}>{children_html}</{self.tag}>"
