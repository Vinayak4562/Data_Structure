class _Node:
    __slot__ = '_element', '_next'
    def __init__ (self, element,next):
        self._element = element
        self._next = next

class linkedlist:
    def __init__ (self):
        self._head=None
        self._tail=None
        self._size=0

    def __len__(self):
        return self._size

    def isempty(self):
        return self._size == 0

    def addfirst (self, e):
        newest = _Node (e, None)
        if self.isempty():
            self._head = newest
            self._tail = newest
        else:
            newest._next = self._head
            self._head = newest
        self._size += 1

    def addlast(self,e):
        newest =_Node(e,None)
        if self.isempty():
            self._head=newest
        else:
            self._tail._next = newest
        self._tail = newest
        self._size += 1

    def addany(self, e, position):
        newest = _Node(e,None)
        p = self._head
        i = 1
        while i < position - 1:
            p = p._next
            i += 1
        newest._next = p._next
        p._next = newest
        self._size += 1

    def search(self, key):
        p = self._head
        index = 0
        while p:
            if p._element == key:
                return index
            p = p._next
            index += 1
        return -1

    def remvfirst(self):
        if self.isempty():
            print("List os empty")
            return
        # v = self._head._element
        e = self._head._next
        self._head = self._head._next
        self._size -= 1
        if self.isempty():
            self._tail = None
        return e

    def remvlast(self):
        if self.isempty():
            print("The list is empty")
            return
        p = self._head
        i = 1
        while i < len(self)-1:
            p = p._next
            i += 1
        self._tail = p
        p = p._next
        e = p._element
        self._tail._next = None
        self._size -= 1
        return e

    def remvany(self, position):
        p = self._head
        i = 1
        while i<position - 1:
            p = p._next
            i += 1
        e = p._next
        p._next = p._next._next
        self._size -= 1
        return e 
       
    def display(self):
        p = self._head
        while p:
            print(p._element, end ="--> ")
            p = p._next
        print()

l = linkedlist()
l.addfirst(40)
l.addfirst(30)
l.addfirst(20)
l.addfirst(10)
l.display()
print("Size is:", len(l))
l.addlast(1)
l.display()
print("Size is: ", len(l))
l.addany(15,2)
l.display()
print("Size is: ", len(l))
# l.search(18)
print(l.search(30))
l.remvfirst()
l.display()
print("Size is: ", len(l))
l.remvlast()
l.display()
print("Size is: ", len(l))
l.remvany(2)
l.display()
print("Size is: ", len(l))
  