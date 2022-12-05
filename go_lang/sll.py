class _Node:
    __slots__ = '_element', '_next'

    def __init__(self, element, next):
        self._element = element
        self._next = next

class LinkedList:
    def __init__(self):
        self._head = None
        self._tail = None
        self._size = 0

    def __len__(self):
        return self._size

    def isempty(self):
        return self._size == 0

    def addfirst(self, e):
        newest = _Node(e, None)
        if self.isempty():
            self._head = newest
            self._tail = newest
        else:
            newest._next = self._head
            self._head = newest
        self._size += 1

    def display(self):
        p = self._head
        while p:
            print(p._element,end='-->')
            p = p._next
        print()

    
l = LinkedList()
l.addfirst(1)
l.display ()
l.addfirst(2)
l.display()
l.addfirst(3)
l.display()