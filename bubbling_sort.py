import random
lst = list(range(10000))
random.shuffle(lst)

def bubbling_sort(lst):
    for i in range(len(lst)-1):
        exchange = False
        for j in range(len(lst)-i-1):
            if lst[j] > lst[j+1]:
                lst[j],lst[j+1] = lst[j+1],lst[j]
                exchange = True
        if not exchange:
            break

bubbling_sort(lst)
print(lst)