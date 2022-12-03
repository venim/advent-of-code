#%%
with open('input/04-in.txt','r') as fin:
    data = fin.read().split('\n')

data = sorted(data)

#%%

import re
import numpy as np
parser = re.compile('\[\d+-\d+-\d+ \d+:(\d+)\] (Guard #(\d+) )?(.*)')

guards = {}
guard = 0
asleep = 0

for line in data:
    result = parser.search(line)
    if result == None:
        continue
    minute, _, guardID, action = result.groups()
    minute = int(minute)

    if guardID != None:
        guard = guardID
        if guardID not in guards:
            guards[guardID] = np.zeros((60))
    elif action == 'falls asleep':
        asleep = minute
    elif action == 'wakes up':
        for i in range(asleep-1, minute):
            guards[guard][i] += 1

sleepiestGuard = max(guards.items(), key=lambda item: len(item[1][item[1]>0]))[0]

print(np.argmax(guards[sleepiestGuard]) * int(sleepiestGuard))

sleepiestMinute = max(guards.items(), key = lambda item: max(item[1]))[0]
print((np.argmax(guards[sleepiestMinute])+1) * int(sleepiestMinute))

#%%
