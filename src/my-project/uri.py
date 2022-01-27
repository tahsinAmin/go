while True:
    everyone, problem = map(int, input().split())

    if everyone == problem == 0:
        break

    characteristics = 4
    arr = []

    for i in range(everyone):
        arr.append([int(j) for j in input().split()])
        characteristics -= (1 if (arr[i].count(1) == 0) else 0)
        characteristics -= (1 if (arr[i].count(0) == 0) else 0)

    problemsAnswered = [0]*problem
    for i in range(problem):
        for j in range(everyone):
            problemsAnswered[j] += 1 if arr[j][i] == 1 else 0

    for i in problemsAnswered:
        characteristics -= 1 if i == 0 else 0
        characteristics -= 1 if i == problem else 0

    print(characteristics, '\n', problemsAnswered)
