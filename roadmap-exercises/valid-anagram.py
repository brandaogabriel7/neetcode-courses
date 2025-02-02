class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        if len(s) != len(t):
            return False

        hash_s = {}
        hash_t = {}

        for i in range(len(s)):
            hash_s[s[i]] = 0 if s[i] not in hash_s else hash_s[s[i]] + 1
            hash_t[t[i]] = 0 if t[i] not in hash_t else hash_t[t[i]] + 1

        return hash_s == hash_t
