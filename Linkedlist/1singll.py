# Creating node using Node class
class _Node:
    __slots__ = '_element', '_next'
    def __init__ (self, element,next):
        self._element = element
        self._next = next

# Creating linkedlist using linkedlist class
class linkedlist:
    def __init__ (self):
        self._head=None
        self._tail=None
        self._size=0

    # creating lenth method for counting the nodes
    def __len__(self):
        return self._size

    # creating isempty method for checking the cobition that empty or not
    def isempty(self):
        return self._size == 0

    # creating adding node in starting the node using addfirst method
    def addfirst (self, e):
        newest = _Node (e, None)
        if self.isempty():
            self._head = newest
            self._tail = newest
        else:
            newest._next = self._head
            self._head = newest
        self._size += 1

    #  adding node at the last in the list using addlast method
    def addlast(self,e):
        newest =_Node(e,None)
        if self.isempty():
            self._head=newest
        else:
            self._tail._next = newest
        self._tail = newest
        self._size += 1

# creating adding node anywhere in the list using addany method
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

# creating function searching element in list using search method
    def search(self, key):
        p = self._head
        index = 0
        while p:
            if p._element == key:
                return index
            p = p._next
            index += 1
        return -1

# removing the fisrt node in the list using remvfirst method
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

# removing the last node in the list using remvlast method
    def remvlast(self):
        if self.isempty():
            print("The list is empty")
            return
        p = self._head
        i = 1
        while i < len(self)-1:
            p = p._next
            i += 1
        e = p._element
        self._tail = p
        p = p._next
        self._tail._next = None
        self._size -= 1
        return e

# removing the node at anywhere in the list using remvany method
    def remvany(self,position):
        p = self._head
        i = 1
        while i < position - 1:
            p = p._next
            i += 1
        e = p._next
        p._next = p._next._next
        self._size -= 1
        return e
       
# creating the display method 
    def display(self):
        p = self._head
        while p:
            print(p._element, end ="--> ")
            p = p._next
        print()

# creating objects a clling them
l = linkedlist()
l.addfirst(40)
l.addfirst(30)
l.addfirst(20)
l.addfirst(10)
l.display() 

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

# l.remvany(3)
# l.display()
# print("Size is: ", len(l))
l.remvany(1)
l.display()
# l.remvany(1)
# l.display()
# l.remvany(1)
# l.display()
l.remvany(0)
l.display()
 