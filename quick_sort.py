import time
import random


lst = list(range(10000))
random.shuffle(lst)


def cal_time(func):
    def inner(*args,**kwargs):
        prev_time = time.time()
        ret = func(*args,**kwargs)
        print('用时%s'%(time.time()-prev_time))
        return ret
    return inner


def _quick_sort(lst,left,right):
    if left < right:
        middle = partition(lst,left,right)
        _quick_sort(lst,left,middle-1)
        _quick_sort(lst,middle+1,right)

def partition(lst,left,right):
    temp = lst[left]
    while left < right:
        while left < right and lst[right] >= temp:
            right -= 1
        lst[left] = lst[right]
        while left < right and lst[left] <= temp:
            left += 1
        lst[right] = lst[left]
    lst[left] = temp
    return left

@cal_time
def quick_sort(lst):
    left = 0
    right = len(lst) - 1
    _quick_sort(lst,left,right)

quick_sort(lst)
print(lst)


lst2 = list(range(10000))
random.shuffle(lst2)


def _quick_sort2(lst):
    if len(lst) < 2:
        return lst
    temp = lst[0]
    left = [i for i in lst[1:] if i <= temp]
    right = [i for i in lst[1:] if i > temp]
    left = _quick_sort2(left)
    right = _quick_sort2(right)
    return left + [temp] + right


@cal_time
def quick_sort2(lst):
    _quick_sort2(lst)


lst2 = quick_sort2(lst2)
print(lst2)