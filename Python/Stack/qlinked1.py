class _Node :
    __slots__ = '_element','_next'
    def __init__(self,element,next):
        self._element = element
        self._next = next

class QLinkedList:
    def __init__(self):
        self._front = None
        self._rear = None
        self._size = 0

    def __len__(self):
        return self._size

    def isempty(self):
        return self._size == 0

    def enqueue(self,e):
        newest = _Node(e,None)
        if self.isempty():
            self._front = newest
        else:
            self._rear._next = newest            
        self._rear = newest
        self._size += 1

    def dequeue(self):
        if self.isempty():
            print("Que is empty")
            return       
        e = self._front._element
        self._front = self._front._next
        self._size -= 1
        if self.isempty():
            self._rear = None
        print("removed Element:" , e)

    def first(self):
        if self.isempty():
            print("Que is empty")
            return
        return self._front._element

    def desplay(self):
        p = self._front
        while p:
            print(p._element, end = "<--")
            p = p._next
        print()


Q = QLinkedList()
print("------------------------------------enqueue OP--------------------------------------")
Q.enqueue(10)
Q.enqueue(20)
Q.enqueue(30)
Q.enqueue(40)
Q.desplay()
print("------------------------------------dequeue OP--------------------------------------")
Q.desplay()
Q.dequeue()
Q.desplay()
print("length of List: ",len(Q))
print("list is empty: ",Q.isempty())
print("------------------------------------first OP--------------------------------------")
Q.first()
Q.desplay()