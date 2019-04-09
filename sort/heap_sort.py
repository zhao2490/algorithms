import random

lst = list(range(10000))
random.shuffle(lst)


def shift(lst, low, high):
    temp = lst[low]
    i = low  # 指向根节点
    j = 2 * i + 1  # 指向空位左节点
    while j <= high:
        # 判断是否有右节点并对比左右节点
        if j + 1 <= high and lst[j] < lst[j + 1]:
            # 如果右节点大于左节点j指向右节点
            j += 1
        # 判断子节点是否大于根节点
        if lst[j] > temp:
            lst[i] = lst[j]
            i = j
            j = 2 * i + 1
        else:
            break
    # 将原始根节点的值放入空位子节点
    lst[i] = temp


def heap_sort(lst):
    high = len(lst)
    for low in range(high // 2 - 1, -1, -1):
        shift(lst, low, high-1)
    for high in range(high - 1, -1, -1):
        lst[0], lst[high] = lst[high], lst[0]
        shift(lst, 0, high-1)


heap_sort(lst)
print(lst)
