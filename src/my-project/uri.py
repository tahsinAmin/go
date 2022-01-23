while True:
    everyone, problem = map(int, input().split())

    if everyone == problem == 0:
        break

    arr = []
    nobody_solved_all = True
    cnt_every_problem_solved_by_at_least_one = 0
    no_problem_solved_by_everyone = True
    cnt_if_solved_by_everyone = 0
    everyone_solved_atleast_one = True

    for i in range(everyone):
        arr.append([int(j) for j in input().split()])
        everyone_solved_atleast_one = False if arr[i].count(1) == 0 else True
        nobody_solved_all = False if arr[i].count(0) == 0 else True

    for i in range(problem):
        is_every_problem_solved_by_at_least_one = False
        if no_problem_solved_by_everyone:
            no_problem_solved_by_everyone = True
        cnt_if_solved_by_everyone = 0
        for j in range(everyone):
            if arr[j][i] == 1:

                if no_problem_solved_by_everyone:
                    cnt_if_solved_by_everyone += 1
                    if cnt_if_solved_by_everyone == everyone:
                        no_problem_solved_by_everyone = False

                if not is_every_problem_solved_by_at_least_one:
                    is_every_problem_solved_by_at_least_one = True
                    cnt_every_problem_solved_by_at_least_one += 1

    score = (1 if nobody_solved_all else 0) + \
        (1 if cnt_every_problem_solved_by_at_least_one == problem else 0) + \
        (1 if no_problem_solved_by_everyone else 0) + \
        (1 if everyone_solved_atleast_one else 0)

    print(score)


# while True:
#     n, m = map(int, input().split())
#     if n == m == 0:
#         break
#     carac = 4
#     competitor = []
#     for i in range(n):
#         competitor.append([int(x) for x in input().split()])

#     for i in range(n):
#         p = competitor[i].count(1)
#         if p == m:
#             carac -= 1
#             break

#     problema = [0] * m
#     for i in range(n):
#         for j in range(m):
#             if competitor[i][j] == 1:
#                 problema[j] += 1
#     if problema.count(0) > 0:
#         carac -= 1

#     for j in range(m):
#         if problema[j] == n:
#             carac -= 1
#             break

#     for i in range(n):
#         p = competitor[i].count(0)
#         if p == m:
#             carac -= 1
#             break

#     print(carac)
