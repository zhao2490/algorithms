import random
lst = list(range(10000))
random.shuffle(lst)

def insert_sort(lst):
    for i in range(1,len(lst)):
        temp = lst[i]
        j = i - 1
        while j >= 0 and lst[j] > temp:
            lst[j+1] = lst[j]
            j -= 1
        lst[j+1] = temp

insert_sort(lst)
print(lst)