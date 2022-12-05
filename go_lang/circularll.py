# creating node
class _Node:
    __slots__ = '_element', '_next'

    def __init__(self, element, next):
        self._element = element
        self._next = next
        

# Creating doubly linked list
class CircularLinkedList:
    def __init__(self):
        self._head = None
        self._tail = None
        self._size = 0
        # Creating length for size of list
    def __len__(self):
        return self._size

        # creating isempty method for checking the cobition that empty or not
    def isempty(self):
        return self._size == 0

        #adding node at the first in list
    def addfirst(self,e):
        newest = _Node(e, None)
        if self.isempty():
            newest._next = newest
            self._head = newest
            self._tail = newest
        else:
            self._tail._next = newest
            newest._next = self._head
            self._head = newest
        self._size += 1

        #adding node at the Last in list
    def addlast(self,e):
        newest = _Node(e, None)
        if self.isempty():
            newest._next = newest
            self._head = newest
        else:
            newest._next = self._tail
            self._tail._next = newest
        self._tail = newest
        self._size += 1

        #adding the nodes in at given index posssition
    def addany(self,e,possition):
        newest = _Node(e,None)
        p = self._head
        i  = 1
        while i < possition - 1 :
            p = p._next
            i += 1
        newest._next = p._next
        p._next = newest
        self._size += 1

        #deleting the first node at list using removingfirstmethod
    def removingfirst(self):
        if self.isempty():
            print ("list is empty")
            return
        e = self._head
        self._tail._next = self._head._next
        self._head = self._head._next
        self._size -= 1
        if self.isempty():
            self._head = None
            self._tail = None
        return e

        #removing the last node
    def removelast(self):
        if self.isempty():
            print ("list is empty")
            return
        p = self._head
        i = 1
        while i < len(self) - 1:
            p = p._next
            i += 1
        self._tail = p
        p = p._next
        self._tail._next = self._head
        e = p._element
        self._size -= 1
        return e

        #removing node with a preside index value
    def removeany(self, position):
        p = self._head
        i = 1
        while i < position - 1:
            p = p._next
            i += 1
        e = p._element
        p._next = p._next._next
        self._size -= 1
        return e

        #displaying the list
    def display(self):
        p = self._head
        i = 0
        while i<len(self):
            print(p._element, end = "-->")
            p = p._next
            i +=1
        print()

 
C = CircularLinkedList()
C.addfirst(40)
C.addfirst(30)
C.addfirst(20)
C.display()
# print('Size:',len(C))
print (len(C))

C.addlast(50)
C.addlast(60)
C.display()

C.addany(15,2)
C.display()

C.removingfirst()
C.display()

C.removelast()
C.display()

C.removeany(2)
C.display()


