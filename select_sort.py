import random
lst = list(range(10000))
random.shuffle(lst)

def select_sort(lst):
    for i in range(len(lst)-1):
        for j in range(i+1,len(lst)):
            if lst[i] > lst[j]:
                lst[i],lst[j] = lst[j],lst[i]

select_sort(lst)
print(lst)