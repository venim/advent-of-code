#%%
with open('05-in.txt','r') as fin:
    data = fin.read().split('\n')

data = data[0]
#%%
def reduce(data):
    import re
    parser = re.compile('([a-zA-Z])\\1', re.I)
    idx = 0

    try:
        while (1):
            match = next(parser.finditer(data[idx:]))
            start = match.start()
            group = match.group(0)
            if group[0] != group[1]:
                data = data[:idx+start] + data[idx+start+2:]
                if idx > 0:
                    idx -= 1
            else:
                idx += 1
    except StopIteration:
        return data:

part1 = reduce(data)
minimum = len(part1)
print(minimum)

#%%
for unit in 'abcdefghijklmnopqrstuvwxyz':
    filtered = part1.replace(unit,'').replace(unit.upper(),'')
    filtered = reduce(filtered)
    if len(filtered) < minimum:
        minimum = len(filtered)

print(minimum)