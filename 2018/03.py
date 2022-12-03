#%% [markdown]
# # Advent of Code 2018
# ## Day 03
#%%
# Open input data file
with open('input/03-in.txt','r') as fin:
    data = fin.read().split('\n')

#%% 
# Use regular expressions to parse each input
# Add the elfID to the fabric dictionary for each cell
import re
parser = re.compile('#(\d*) @ (\d*),(\d*): (\d*)x(\d*)')

fabric = {}
elfIDs = set()
for elf in data:
    result = parser.search(elf)
    if result == None:
        print(elf)
        continue
    ID, x, y, dx, dy = map(int,result.groups())
    elfIDs.add(ID)
    for r in range(x, x+dx):
        for c in range(y, y+dy):
            cell = (r,c)
            if cell in fabric.keys():
                fabric[cell].add(ID)
            else:
                fabric[cell] = set([ID])

# If there's more than one ID in a cell, add it to the overlap count 
# Remove from elfID set for Part 2 
overlap = 0
for cell, IDs in fabric.items():
    if len(IDs) > 1:
        overlap += 1
        for ID in IDs:
            if ID in elfIDs:
                elfIDs.remove(ID)

#%%
# Part 1 Answer:
overlap

#%%
# Part 2 Answer
elfIDs


#%%
