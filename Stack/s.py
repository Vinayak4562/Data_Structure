class _Node:
    __slots__ ='_element','_next'
    def __init__(self,element,next):
        self._element = element
        self._next = next

class sll :
    def __init__(self):
        self._head =  None
        self._tail = None
        self._size = 0
    
    def __len__(self):
        return self._size

    def isempty(self):
        return self._size == 0

    def addfirst(self,e):
        newest = _Node(e,None)
        if self.isempty():
            self._head = newest
            self._tail = newest
        else:
            newest._next = self._head
            self._head= newest
        self._size +=1

    def addlast(self,e):
        newest = _Node (e,None)
        if self.isempty():
            self._head = newest
            self._tail = newest
        else:
            self._tail._next = newest
            self._tail = newest
        self._size += 1

    def addany(self,e,position):
        newest = _Node(e,None)
        p = self._head
        i = 0       
        while i < position - 1:
            p = p._next
            i += 1
        newest._next = p._next
        p._next = newest
        self._size += 1

    def remvfirst(self):
        if self.isempty():
            print("List is empty")
            return 
        e = self._head._element
        self._head = self._head._next
        self._size -= 1
        print(e)

    def remvlast(self):
        if self.isempty():
            print("List is empty")
            return
        p = self._head
        i = 1
        while i < len(self):
            p = p._next
            i += 1
        

    def display(self):
        p = self._head
        while p:
            print(p._element, end ="--> ")
            p = p._next
        print()


l = sll()
l.addfirst(1)
l.addfirst(2)
l.addfirst(3)
l.display()
l.addlast(222)
l.display()
l.addany(10,2)
l.display()
l.remvfirst()
l.display()
l.remvlast()
l.display()
