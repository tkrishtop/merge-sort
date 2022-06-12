def merge(l1: list, l2: list) -> list:
    '''
    Two pointers: merge two sorted lists
    '''
    output = []
    p1 = p2 = 0
    while p1 < len(l1) and p2 < len(l2):
        if l1[p1] <= l2[p2]:
            output.append(l1[p1])
            p1 += 1
        else:
            output.append(l2[p2])
            p2 += 1
        
    if p1 < len(l1):
        output.extend(l1[p1:])
    else:
        output.extend(l2[p2:])

    return output    


def merge_sort(l: list) -> list:
    '''
    Top-down merge sort
    '''
    n = len(l)
    if n < 2:
        return l
    
    mid = n // 2
    left = merge_sort(l[:mid])
    right = merge_sort(l[mid:])
    return merge(left[:], right[:])


if __name__=='__main__':
    l1 = [1, 3, 5, 4, 6, 8, 9, 10, 2, 7]
    assert merge_sort(l1) == sorted(l1)
