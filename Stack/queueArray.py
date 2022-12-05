class queuearray:
    def __init__(self):
      self._data = []

    def __len__(self):
        return len(self._data)

    def isempty(self):
        return len(self._data) == 0

    def enqueue(self,e):
        self._data.append(e)

    def dequeue(self):
        if self.isempty():
            print("Queue is empty")
            return
        return self._data.pop(0)

    def first (self):
        if self.isempty():
            print("Queue is empty")
            return
        return self._data[0]

    # def display(self):
    #     p = self._top
    #     while p:
    #         print(p._element, end = " --> ")
    #         p = p._next
    #     print()

S = queuearray()
S.enqueue(20)
S.enqueue(30)
S.enqueue(40)
S.enqueue(50)
# S.display()
print(S._data)
print(len(S))
print(S.isempty())

S.dequeue()
print(S._data)
print(len(S))
S.dequeue()
S.dequeue()
S.dequeue()
print(S._data)
print(len(S))
print(S.isempty())




