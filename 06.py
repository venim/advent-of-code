#%%
with open('input/06-in.txt','r') as fin:
    data = fin.read().split('\n')
    data.pop()
# data = '''1, 1
# 1, 6
# 8, 3
# 3, 4
# 5, 5
# 8, 9'''.split('\n')
# data

#%%
coordinates = []
maxx = 0
maxy = 0
for i,v in enumerate(data):
    v = list(map(int,v.split(', ')))
    coordinates.append( {'x':v[0], 'y':v[1]})
    if v[0] > maxx:
        maxx = v[0]
    if v[1] > maxy:
        maxy = v[1]

def distance(coordinates, r, c,):
    total = 0
    shortest = None
    for i, cell in enumerate(coordinates):
        distance = abs(cell['x']-r) + abs(cell['y']-c)
        total += distance
        if shortest == None or distance < shortest[1]:
            shortest = [i, distance]
        elif distance == shortest[1]:
            shortest += [i, distance]

    if len(shortest) == 2:
        shortest = shortest[0]
    else:
        shortest = -1

    return shortest, total

import numpy as np
Part1 = np.zeros((maxx+1, maxy+1))
Part2 = np.zeros((maxx+1, maxy+1))

for r in range(maxx+1):
    for c in range(maxy+1):
        Part1[r,c], Part2[r,c] = distance(coordinates,r,c)
    
noninfinite = set(range(len(coordinates)))
for v in np.concatenate((Part1[0], Part1[:,0], Part1[-1], Part1[:,-1])):
    if v in noninfinite:
        noninfinite.remove(v)

maxarea = 0
for i in list(noninfinite):
    area = len(Part1[np.where(Part1 == i)])
    if area > maxarea:
        maxarea = area

print(maxarea)
print(len(Part2[np.where(Part2 < 10000)]))

#%%
