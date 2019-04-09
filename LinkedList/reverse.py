class ListNode():
    def __init__(self, val):
        """
        节点
        :param val: 节点的值
        """
        self.val = val
        self.next = None


class LinkedList():
    def __init__(self, lst, method='tail'):
        """
        :param lst: 列表 list
        :param method: 使用方法 头插法/尾插法
        """
        self.head = None
        self.tail = None
        if method == 'tail':
            # 尾插法
            self.create_linklist_head(lst)
        elif method == 'head':
            # 头插法
            self.create_linklist_tail(lst)
        else:
            raise ValueError('Unsupported value %s' % method)

    def create_linklist_head(self, lst):
        """
        将列表转化为链表
        :param lst: 列表
        """
        self.head = ListNode(0)
        for i in lst:
            temp = ListNode(i)
            temp.next = self.head.next
            self.head.next = temp
            self.head.val += 1

    def create_linklist_tail(self, lst):
        """
        将列表转化为反转链表
        :param lst: 列表
        """
        self.head = ListNode(0)
        self.tail = self.head
        for i in lst:
            temp = ListNode(i)
            self.tail.next = temp
            self.tail = temp
            self.head.val += 1

    def __iter__(self):
        """
        返回一个链表的生成器
        :return: 节点值
        """
        temp = self.head.next
        while temp:
            yield temp.val
            temp = temp.next

    def __len__(self):
        return self.head.val

    def reverse(self):
        root_node = ListNode(self.head.val)
        self.head = self.head.next
        while self.head:
            temp = ListNode(self.head.val)
            temp.next = root_node.next
            root_node.next = temp
            self.head = self.head.next
        self.head = root_node


lst = [1, 2, 3, 4, 5, 6]
link_lst = LinkedList(lst, method='head')
link_generator = link_lst.__iter__()
for i in link_generator:
    print(i)
link_lst.reverse()
link_generator = link_lst.__iter__()
for i in link_generator:
    print(i)
