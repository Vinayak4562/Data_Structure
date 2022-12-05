class _Node:
    __slots__ ='_element','_next'
    def __init__(self,element,next):
        self._element = element
        self._next = next

class DEQueue :
    def __init__(self):
        self._front =  None
        self._rear = None
        self._size = 0
    
    def __len__(self):
        return self._size

    def isempty(self):
        return self._size == 0

    def enqueuefirst(self, e):
        newest = _Node(e,None)
        if self.isempty():
            self._rear = newest
            self._front = newest
        else:
            newest._next = self._front
            self._front = newest
        self._size += 1

    def enqueuelast(self,e):
        newest = _Node(e, None)
        if self.isempty():
            self._front = newest
            self._rear = newest
        else:
            self._rear._next = newest
            self._rear = newest
        self._size += 1

    def dequeueFirst(self):
        if self.isempty():
            print("Queue is empty")
            return
        else:
            e = self._front._element
            self._front = self._front._next
        self._size -= 1
        if self.isempty():
            self._front = None
        print(e)
        
    def dequeLast(self):
        if self.isempty():
            print("Queue is empty")
            return
        p = self._front
        i = 1
        while i < len(self)-1:
            p = p._next
            i += 1
        self._rear = p
        p = p._next
        e = p._element
        self._rear._next = None
        self._size -= 1
        print(e)
          

    def desplay(self):
        p = self._front
        while p:
            print(p._element, end = "<--")
            p = p._next
        print()



Q = DEQueue()
Q.enqueuefirst(10)
Q.enqueuefirst(20)
Q.enqueuefirst(30)
Q.desplay()
print("-----------------------------------------------")
Q.enqueuelast(1)
Q.enqueuelast(2)
Q.enqueuelast(3)
Q.desplay()
print(len (Q))
print("-----------------------------------------------")
Q.dequeLast()
Q.desplay()
print("-----------------------------------------------")
Q.dequeueFirst()
Q.desplay()
print("-----------------------------------------------")

