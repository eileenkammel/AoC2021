with open("puzzle-2input") as puzzle:
    movements = puzzle.read()
movements = movements.split("\n")


class Submarine():

    horizontal = 0
    depth = 0
    aim = 0

    def move_forward(self, unit):
        self.horizontal += unit
        self.depth += (self.aim * unit)

    def move_up(self, unit):
        self.aim -= unit

    def move_down(self, unit):
        self.aim += unit


sub = Submarine()

for movement in movements:
    direction, unit = movement.split()
    unit = int(unit)
    if direction == "forward":
        sub.move_forward(unit)
    if direction == "up":
        sub.move_up(unit)
    if direction == "down":
        sub.move_down(unit)

print(sub.depth * sub.horizontal)
