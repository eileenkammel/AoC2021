with open("puzzle1-1input") as puzzleinput:
    raw_input = puzzleinput.read()
measures = raw_input.split()

measures_as_int = list(map(int, measures))

counter = 0

def compare(previous, current):
    if current > previous:
        global counter
        counter += 1

for i in range(len(measures_as_int)-1):
    compare(measures_as_int[i], measures_as_int[i+1])

print(counter)