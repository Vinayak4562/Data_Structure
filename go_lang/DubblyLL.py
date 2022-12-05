class _Node :
    __slots__ = '_element', '_next', '_prev'
    def __init__ (self,element,next,prev):
        self._element = element
        self._next = next
        self._prev = prev

class dublyLinkedList :
    def __init__ (self):
        self._head = None
        self._tail = None
        self._size = 0
    def __len__ (self):
        return self._size

    def isempty(self):
        return self._size == 0

    def addfirst(self,e):
        newest = _Node(e,None,None)
        if self.isempty():
            self._head = newest
            self._tail = newest
        else:
            newest._next = self._head
            self._head._prev = newest
            self._head = newest
        self._size += 1

    def addlast(self,e):
        newest = _Node(e,None,None)
        if self.isempty():
            self._head = newest
            self._tail = newest
        else:
            self._tail._next = newest
            newest._prev = self._tail
            self._tail = newest
        self._size += 1

    def addany(self,e,position):
        newest = _Node(e,None,None)
        p = self._head
        i = 1
        while i < position - 1:
            p = p._next
            i += 1
        newest._next = p._next
        p._next._prev = newest
        p._next = newest
        newest._prev = p
        self._size += 1

    def remvfirst(self):
        if self.isempty():
            print("List is empty")
            return
        e = self._head._element
        self._head =  self._head._next
        self._head._prev = None
        self._size -= 1
        if self.isempty():
            self._tail = None
        return e

    def remvlast(self):
        if self.isempty():
            print("List is empty")
            return
        e = self._tail._element
        self._tail = self._tail._prev
        self._tail._next = None
        self._size -= 1
        return e
        
    def  remvany(self, position):
        p = self._head
        i = 1
        while i < position - 1:
            p = self._next
            i += 1
        e = p._next._element
        p._next = p._next._next
        p._next._prev = p
        self._size -= 1
        return e

        
    def display (self):
        p = self._head
        while p:
            print(p._element, end = "-->")
            p = p._next
        print()

    def displayRev (self):
        p = self._tail
        while p:
            print(p._element, end ="<--")
            p = p._prev
        print()


D = dublyLinkedList()
D.addfirst(40)
D.addfirst(30)
D.addfirst(20)
D.addfirst(10)
D.display()
D.displayRev()

D.addlast(500)
D.display()

D.addany(222,2)
D.display()

D.remvfirst()
D.display()

D.remvlast()
D.display()

D.remvany(1)
D.display()