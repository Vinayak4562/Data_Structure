class _Node:
    __slots__ = '_element', '_next'
    def __init__(self,element,next):
        self._element = element
        self._next = next

class StackLinked:
    def __init__(self):
        self._top = None
        self._size = 0

    def __len__(self):
        return self._size

    def isempty(self):
        return self._size == 0

    def push(self,e):
        newest = _Node(e, None)
        if self.isempty():
            self._top = newest
        else:
            newest._next = self._top # link is created between top to newest
            self._top = newest  # top is moved to newest
        self._size += 1

    def pop(self):
        if self.isempty():
            print("Stack is empty")
            return
        e = self._top._element
        self._top = self._top._next
        self._size -= 1
        return e

    def top(self):
        if self.isempty(self):
            print("Stack is empty")
            return
        return self._top._element

        
    def display(self):
        p = self._top
        while p:
            print(p._element, end = " --> ")
            p = p._next
        print()


S = StackLinked()
print("----------push method O/P---------")
# S.push(2)
# S.push(3)
S.display()
S.push(4)
S.push(5)
S.display()
print("-------------------------pop method O/P---------")
S.pop()
S.display()
S.pop()
S.display()
print("-------------------------len method O/P---------")
print(len(S))
print("---------------------isempty method O/P---------")
print(S.isempty())