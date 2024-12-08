import os

current_dir = os.path.dirname(__file__)
file_path = os.path.join(current_dir, 'input.txt')


with open(file_path, 'r') as file:
    data = file.read().split("\n\n")

rules, updates = data
rules = {(*map(int, line.split("|")),) for line in rules.splitlines()}


class Sortable:

    def __init__(self, x):
        self.x = x

    def __eq__(self, other):
        return self.x == other.x

    def __lt__(self, other):
        return (other.x, self.x) not in rules


part_a = part_b = 0

for line in updates.splitlines():
    unsorted_update = [*map(int, line.split(","))]
    sorted_update = sorted(unsorted_update, key=Sortable)
    if sorted_update == unsorted_update:
        part_a += sorted_update[len(sorted_update) // 2]
    else:
        part_b += sorted_update[len(sorted_update) // 2]

print(part_a, part_b)