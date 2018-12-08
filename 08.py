#%%
with open('input/08-in.txt','r') as fin:
    data = fin.read().split('\n')
    data.pop()
    data = list(map(int,data[0].split()))


def process(data, answer=0):
    children, meta = data[:2]
    data = data[2:]
    nodes = []

    if children == 0:
        n = sum(data[:meta])
        answer += n
        return answer, n, data[meta:]

    for i in range(children):
        answer, node, data = process(data, answer)
        nodes.append(node)
    
    n = sum(nodes[k-1] for k in data[:meta] if k > 0 and k <= len(nodes))
    answer += sum(data[:meta])
    return answer, n, data[meta:]

part1, part2, _ = process(data)
print(part1)    
print(part2)
    
#%%
