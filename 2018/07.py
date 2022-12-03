#%%
with open('input/07-in.txt','r') as fin:
    data = fin.read().split('\n')
    data.pop()

#%%
prereq = {}
steps = set()

for d in data:
    _, pre, *_, step, _, _ = d.split(' ')
    steps |= {pre, step}
    if step not in prereq:
        prereq[step] = {pre}
    else:
        prereq[step].add(pre)

#%%
instructions = ''

part1 = set(steps)
prereq1 = {k:set(v) for k,v in prereq.items()}
while(part1):
    potential = sorted([step for step in part1 if step not in prereq1.keys()])
    step = potential[0]
    part1.remove(step)

    for key in list(prereq1.keys()):
        if step in prereq1[key]:
            prereq1[key].remove(step)
        if not prereq1[key]:
            del(prereq1[key])
    instructions += step

instructions

#%%

t = 0
workers = [None for i in range(5)]
part2 = set(steps)
prereq2 = {k:set(v) for k,v in prereq.items()}
queue = []
finished = ''
while(part2):

    potential = sorted([step for step in part2 if step not in prereq2.keys() and step not in [worker[1] for worker in workers if worker != None]])
    queue += potential

    part2 -= set(queue)

    for i, worker in enumerate(workers):
        if worker == None:
            if queue:
                step = queue.pop(0)
                worker = workers[i] = (ord(step)-65+60+t, step)
                
        if worker != None:
            if worker[0] == t:
                finished += worker[1]
                workers[i] = None

                for key in list(prereq2.keys()):
                    if worker[1] in prereq2[key]:
                        prereq2[key].remove(worker[1])
                    if not prereq2[key]:
                        del(prereq2[key])


    t += 1
workers[0][0]+1

#%%
