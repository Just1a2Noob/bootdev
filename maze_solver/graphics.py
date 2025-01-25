from tkinter import BOTH, Canvas, Tk


class Window:
    def __init__(self, width, height):
        self.root = Tk()
        self.root.title("Maze Solver")

        self.canvas = Canvas(self.root, bg="white", width=width, height=height)
        self.canvas.pack(fill=BOTH, expand=1)

        self.running = False

        self.root.protocol("WM_DELETE_WINDOW", self.close)

    def redraw(self):
        self.root.update_idletasks()
        self.root.update()

    def draw_line(self, line, fill_color="black"):
        line.draw(self.canvas, fill_color)

    def wait_for_close(self):
        self.running = True
        while self.running == True:
            self.redraw()
        print("window closed...")

    def close(self):
        self.running = False
        self.root.quit()


class Point:
    def __init__(self, x1, y1):
        self.x = x1
        self.y = y1


class Line:
    def __init__(self, point1, point2):
        self.point1 = point1
        self.point2 = point2

    def draw(self, canvas, fill_color="black"):
        canvas.create_line(
            self.point1.x,
            self.point1.y,
            self.point2.x,
            self.point2.y,
            fill=fill_color,
            width=2,
        )
