class Guia:
    def __init__(self) -> None:
        self.lines = []
        self.vertical_lines = []
        self.horizontal_lines = []
        self.pivot_x_width = None
        self.pivot_y_height = None
        self.prancheta_height = None
        self.prancheta_width = None

    def set_pivot(self, pivot_x_width: int, pivot_y_height: int):
        self.pivot_x_width = pivot_x_width
        self.pivot_y_height = pivot_y_height

    def set_prancheta(self, width: int, height: int):
        self.prancheta_height = height
        self.prancheta_width = width

    def _cal_div(self, x, y):
        return (y - y % x) / x

    def _create_line(self, xi, yi, xii, yii):
        line = [(xi, yi), (xii, yii)]
        self.lines.append(line)
        return line

    def _create_lines(self, border_gap, total_lines, pivo_lenght, horizontal=0):
        for line in range(total_lines):
            if horizontal == 0:
                step_x = border_gap + (line * pivo_lenght)
                self.horizontal_lines.append(
                    self._create_line(step_x, 0, step_x, self.prancheta_height)
                )
            else:
                step_y = border_gap + (line * pivo_lenght)
                self.vertical_lines.append(
                    self._create_line(0, step_y, self.prancheta_width, step_y)
                )
        return self.lines

    def calculate(self):
        total_guias_vertical = int(
            self._cal_div(self.pivot_x_width, self.prancheta_width)
        )
        total_guias_horizontal = int(
            self._cal_div(self.pivot_y_height, self.prancheta_height)
        )

        border_gap_vertical = (self.prancheta_width % self.pivot_x_width) / 2
        border_gap_horizontal = (self.prancheta_height % self.pivot_y_height) / 2
        self._create_lines(
            border_gap_horizontal, total_guias_horizontal, self.pivot_y_height, 1
        )
        self._create_lines(
            border_gap_vertical, total_guias_vertical, self.pivot_x_width
        )

        ## line in end of horizontal gap
        self._create_line(
            0,
            self.prancheta_height - border_gap_horizontal,
            self.prancheta_width,
            self.prancheta_height - border_gap_horizontal,
        )

        ## line in end of vertical gap
        self._create_line(
            self.prancheta_width - border_gap_vertical,
            0,
            self.prancheta_width - border_gap_vertical,
            self.prancheta_height,
        )
        return

    def draw(self):
        img = Image.new("RGB", prancheta.size())
        draw = ImageDraw.Draw(img)
        for item in guia.lines:
            draw.line(item, fill=None, width=3)
        plt.imshow(img)
        plt.axis("off")  # Hide axis
        plt.show()
