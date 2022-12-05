class _Node :
    __slot__ = '_element', '_next'
    def __init__(self, element, next):
        self._element = element
        self._next = next


class linked :
    def __init__(self):
        self._head = None
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
        self._head = newest 
        self._size += 1

    def addlast(self,e):
        newest = _Node(e,None)
        if self.isempty():
            self._head = newest 
        else:
            self._tail._next = newest
        self._tail = newest._next
        self._size += 1

    def addany(self,e, position):
        newest = _Node(e,None)
        p = self._head
        i = 1
        while i < position-1:
            p = p._next
            i += 1
        newest._next = p._next
        p._next = newest
        self._size +=1                        
    
    def display(self):
        p = self._head
        while p:
            print(p._element, end = "-->")
            p = p._next
        print()


l = linked()
l.addfirst(40)
l.addfirst(30)
l.addfirst(20)
l.addfirst(10)
l.display()
l.addlast(40)
l.display()
print(len(l))
l.addany(1,2)
l.display()
print(len(l))
