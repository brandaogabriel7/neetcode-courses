class Solution:
    def isValid(self, s: str) -> bool:
        parentheses = {"(": ")", "[": "]", "{": "}"}
        stack = []
        for ch in s:
            if ch in parentheses:
                stack.append(ch)
            else:
                if len(stack) == 0:
                    return False

                last_char = stack[-1]
                if parentheses[last_char] == ch:
                    stack.pop()
                else:
                    return False

        return len(stack) == 0
