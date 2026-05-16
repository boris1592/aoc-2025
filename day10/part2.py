from z3 import Int, Optimize

all_buttons = []
targets = []

with open('input') as file:
    for line in file:
        if len(line) == 0:
            continue

        parts = line.split()
        buttons = []

        for button in parts[1:-1]:
            buttons.append(list(map(int, button[1:-1].split(','))))

        all_buttons.append(buttons)
        targets.append(list(map(int, parts[-1][1:-1].split(','))))

total = 0

for group in range(len(all_buttons)):
    buttons = all_buttons[group]
    target = targets[group]
    variables = []
    s = Optimize()

    for i in range(len(buttons)):
        variables.append(Int(f'x{i}'))
        s.add(variables[i] >= 0)
 
    for var, val in enumerate(target):
        expr = -val

        for i, btn in enumerate(buttons):
            if var not in btn:
                continue

            expr = expr + variables[i]

        s.add(expr == 0)

    sum_expr = 0

    for var in variables:
        sum_expr = sum_expr + var

    s.minimize(sum_expr)
    print(s.check())
    model = s.model()

    for var in variables:
        total += model[var].as_long()

print(total)
