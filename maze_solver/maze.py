import random
import time

from cell import Cell


class Maze:
    def __init__(
        self,
        x1,
        y1,
        num_rows,
        num_cols,
        cell_size_x,
        cell_size_y,
        win=None,
    ):
        self._cells = []
        self._x1 = x1
        self._y1 = y1
        self.num_rows = num_rows
        self.num_cols = num_cols
        self.cell_size_x = cell_size_x
        self.cell_size_y = cell_size_y
        self._win = win

        self._create_cells()
        self._break_entrance_and_exit()

    def _create_cells(self):
        for i in range(self.num_cols):
            col_cells = []
            for j in range(self.num_rows):
                col_cells.append(Cell(self._win))
            self._cells.append(col_cells)

        for i in range(self.num_cols):
            for j in range(self.num_rows):
                self._draw_cell(i, j)

    def _draw_cell(self, i, j):
        if self._win is None:
            return

        x1 = self._x1 + i * self.cell_size_x
        y1 = self._y1 + j * self.cell_size_y
        x2 = x1 + self.cell_size_x
        y2 = y1 + self.cell_size_y

        self._cells[i][j].draw(x1, y1, x2, y2)
        self._animate()

    def _break_entrance_and_exit(self):
        self._cells[0][0].has_top_wall = False
        self._draw_cell(0, 0)
        self._cells[self.num_cols - 1][self.num_rows - 1].has_bottom_wall = False
        self._draw_cell(self.num_cols - 1, self.num_rows - 1)

    # Helper function for _break_walls_r
    def _index_finder(self, lists, cell):
        for sub_list in lists:
            if char in sub_list:
                return (sub_list.index(cell), lists.index(sub_list))
        raise ValueError("Cell is not within the list")

    def _break_walls_r(self, i, j):
        self._cells[i][j].visited = True

        to_visit = []

        # Neighbours
        left = [i - 1, j]
        right = [i + 1, j]
        up = [i, j - 1]
        down = [i, j + 1]
        neighbours = [left, right, down, up]
        for n in neighbours:
            try:
                to_visit.append(self._cells[n[0]][n[1]])
            except IndexError:
                continue

        # Check if it has any neighbours at all
        if len(to_visit) < 1:
            self._draw_cell(i, j)
            return
        else:
            curr_cell = self._cells[i][j]
            chosen_cell = random.choice(to_visit)
            chosen_cell_index = self._index_finder(self._cells, chosen_cell)

            # if conditions for breaking walls:
            if i > chosen_cell_index[0]:
                curr_cell.has_left_wall = False
                chosen_cell.has_right_wall = False
            if i < chosen_cell_index[0]:
                curr_cell.has_right_wall = False
                chosen_cell.has_left_wall = False

            if j > chosen_cell_index[1]:
                curr_cell.has_bottom_wall = False
                chosen_cell.has_top_wall = False
            if j < chosen_cell_index[1]:
                curr_cell.has_top_wall = False
                chosen_cell.has_bottom_wall = False

            self._break_walls_r(chosen_cell[0], chosen_cell[1])

    def _animate(self):
        if self._win is None:
            return
        self._win.redraw()
        time.sleep(0.05)
