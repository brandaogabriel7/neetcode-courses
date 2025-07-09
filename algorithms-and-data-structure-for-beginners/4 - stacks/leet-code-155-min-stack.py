class MinStack:
    def __init__(self):
        self._stack = []

    def push(self, val: int) -> None:
        self._stack.append(val)

    def pop(self) -> None:
        self._stack.pop()

    def top(self) -> int:
        return self._stack[-1]

    def getMin(self) -> int:
        minimum = None
        for num in self._stack:
            if minimum is None or num < minimum:
                minimum = num

        return minimum
