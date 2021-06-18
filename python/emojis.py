
ツ = "This is a letter in Unicode"
print(ツ)

ℌ = "ℌ"
H = "H"
print(ℌ)

import unicodedata

print(unicodedata.normalize('NFKC', "ℌ"))
