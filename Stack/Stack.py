class stackarray:
    def __init__(self):
        self._data = [] # which is use to store the data

    def __len__(self):
        return len (self._data)

    def isempty(self):
        return len(self._data) == 0

    def push(self,e):
       self._data.append(e)

    def pop(self):
        if self.isempty():
            print("Stack is empty")
            return
        return self._data.pop()

    def top(self):
        if self.isempty():
            print("Stack is empty")
            return
        return self._data[-1]

S = stackarray()
S.push(10)
S.push(30)
S.push(40)
S.push(50)
S.push(60)
print("List is", S._data)
print("Element: ", len(S))
print(S.isempty())

print(S.pop())
print(S.top())
print(S._data)
