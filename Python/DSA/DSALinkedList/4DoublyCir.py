class _Node:
    __slots__ = '_element', '_next', '_prev'
    def __init__(self,element,next,prev):
        self._element = element
        self._next = next
        self._prev = prev

class DoublyCLL :
    def __init__(self):
        self._head = None
        self._tail =  None
        self._size = 0

    def __len__(self):
        return self._size

    def isempty(self):
        return self._size == 0

    def addfirst(self, e):
        newest = _Node(e,None,None)
        if self.isempty():
            newest._next = newest
            self._head = newest
            self._tail = newest
        else:
            newest._next = self._head
            self._head._prev = newest
            self._tail._next = newest
            self._head = newest
        self._size += 1

    def addlast(self,e):
        newest = _Node(e,None,None)
        if self.isempty():
            newest._next = newest
            self._head = newest
            self._tail = newest
        else:
            self._tail._next = newest
            newest._prev = self._tail
            self._tail = newest
            self._tail._next = self._head
            self._head._prev = self._tail
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

    def removfirst(self):
        if self.isempty():
            print("List is empty")
            return
        e = self._head._element
        self._head = self._head._next
        self._head._prev = self._tail
        self._size -= 1

    def remvlast(self):
        if self.isempty():
            print("List is empty")
            return
        e = self._tail._element
        self._tail = self._tail._prev
        self._tail._next = self._head
        self._size -= 1
        return e
        
    def remvany(self, position):
        p = self._head
        i = 1
        while i < position - 1:
            p = p._next
            i += 1
        e = p._next._element
        p._next = p._next._next
        p._next._prev = p
        self._size -= 1
        return e
     
  
    def display(self):
        p = self._head
        i = 0
        while i < len(self):
            print(p._element,end='-->')
            p = p._next
            i += 1
        print()

    def displayRev(self):
        p = self._tail
        i = 0
        while i < len(self):
            print(p._element,end='<--')
            p = p._prev
            i += 1
        print()

C1 = DoublyCLL()
C1.addfirst(10)
C1.addfirst(20)
C1.addfirst(30)
C1.addfirst(40)
C1.display()
C1.displayRev()
print("-------------Addlasrt-------------------------------------")
C1.addlast(11)
C1.addlast(22)
C1.display()
C1.displayRev()
print("---------------------------------------------------------------")
C1.addany(2,2)
C1.addany(2,4)
C1.display()
C1.displayRev()
print("---------------------------------------------------------------")
C1.removfirst()
C1.display()
C1.displayRev()
print("---------------------------------------------------------------")
C1.remvlast()
C1.display()
C1.displayRev()
print("---------------------------------------------------------------")
C1.remvany(2)
C1.display()
C1.displayRev()
C1.addany(222,6)
C1.display()
C1.displayRev()
print("---------------------------------------------------------------")