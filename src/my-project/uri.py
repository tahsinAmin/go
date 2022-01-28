while True:
    everyone, problem = map(int, input().split())

    if everyone == problem == 0:
        break

    characteristics = 4
    arr = []

    for i in range(everyone):
        arr.append([int(j) for j in input().split()])

        characteristics -= 1 if (arr[i].count(1) == problem) else 0  # 1
        characteristics -= 1 if (arr[i].count(1) < 1) else 0  # 4

    problemsAnswered = [0]*problem
    for i in range(problem):
        for j in range(everyone):
            problemsAnswered[i] += 1 if arr[j][i] == 1 else 0

    characteristics -= 1 if 0 in problemsAnswered else 0  # 2
    characteristics -= 1 if everyone in problemsAnswered else 0  # 3

    print(characteristics, '\n', problemsAnswered)


# 3 3
# 1 1 1
# 1 1 1
# 1 1 1
