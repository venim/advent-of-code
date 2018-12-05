#%%
with open('01-in.txt','r') as fin:
    data = fin.readlines()

#%%
result = 0
frequencies = set([result])
found = False
while(not found):
    for d in data:
        result = eval('%d%s' %(result, d.strip('\n')))
        if result not in frequencies:
            frequencies.add(result)
        else:
            print(result)
            found = True
            break
#print(result)
