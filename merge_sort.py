
lst1 = list(range(100))
lst2 = list(range(100))
lst = lst1 + lst2

def merge_sort(lst, low, mid, high):
    # 列表分为两段有序列表 [low,mid] [mid+1,high]
    i = low
    j = mid + 1
    li_temp = []
    while i <= mid and j <= high:
        if lst[i] <= lst[j]:
            li_temp.append(lst[i])
            i += 1
        else:
            li_temp.append(lst[j])
            j += 1
    while i <= mid:
        li_temp.append(lst[i])
        i += 1
    while j <= high:
        li_temp.append(lst[j])
        j += 1
    lst[low:high+1] = li_temp

merge_sort(lst, 0, len(lst)//2, len(lst)-1)
print(lst)