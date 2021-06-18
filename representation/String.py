

class String:

    UnicodeData = {}

    def __init__(self, codepoints=[]):
        self.codepoints = codepoints

    def __len__(self):
        return len(self.codepoints)

    def __getitem__(self, index):
        # TODO check boundaries
        return self.codepoints[index]

    def upper(self):
        res = []
        for cl in self.codepoints:
            cu = String.UnicodeData[cl].get("Simple_Uppercase_Mapping", None)
            if cu:
                res.append(int("0x" + cu, 0))
            else:
                res.append(cl)
        return String(res)

def loadUCD():
    ucd = {}
    with open('./UnicodeData.txt') as fp:
        for line in fp:
            (codepoint, *_, upper, lower, title) = line.split(";")
            ucd[int("0x" + codepoint, 0)] = {
                "Simple_Uppercase_Mapping": upper,
                "Simple_Lowercase_Mapping": lower,
                "Simple_Titlecase_Mapping": title,
            }
    String.UnicodeData = ucd

loadUCD()

s = String([0x0068, 0x0065, 0x0079, 0x1F600]) # "hey\N{Grinning Face Emoji}"
print(len(s)) # 4
print(s[0] == ord("h")) # True

print("".join(map(chr, s.upper()))) # Convert bytes to string
