#%%
with open('input/09-in.txt','r') as fin:
    data = fin.read().split('\n')
    data.pop()

players, *_, maxmarbles, _ = data[0].split()
players = int(players)
maxmarbles = int(maxmarbles)

#%%
import numpy as np
from collections import deque

def process(players, maxmarbles):
    marbles = deque([0])
    marbleidx = 0
    elves = [0 for i in range(players)]

    for i in range(maxmarbles):

        if (i+1) % 23 == 0:
            marbles.rotate(7)
            elves[i % players] += i+1 + marbles.pop()
            marbles.rotate(-1)

        else:
            marbles.rotate(-1)
            marbles.append(i+1)

    return max(elves)

#%%
print(process(players, maxmarbles))
print(process(players, maxmarbles*100))

#%%
