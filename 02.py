#%% [markdown]
# # Open input data file
#%%
with open('02-in.txt','r') as fin:
    data = fin.read().split('\n')

#%% [markdown]
# # Part 1

#%%
winners = {
    2: set([]),
    3: set([])
}
for ID in data:
    letters = {}
    for letter in ID:
        if letter in letters.keys():
            letters[letter] += 1
        else:
            letters[letter] = 1
    
    letters = {k:v for k,v in letters.items() if v==2 or v==3}
    for k,v in letters.items():
        winners[v].add(ID)

#%% [markdown]
# ## Answer:
print(len(winners[2])*len(winners[3]))

#%% [markdown]
# # Part 2
#%%
best = None
for index, ID in enumerate(data):
    for ID2 in data[index+1:]:
        differences = 0
        answer = ''
        for i in range(len(ID)):
            if ID[i] != ID2[i]:
                differences += 1
            else:
                answer += ID[i]

        if best == None or differences <  best[0]:
            best = (differences, answer)

#%% [markdown]
# ## Answer:
print(best)

#%%
