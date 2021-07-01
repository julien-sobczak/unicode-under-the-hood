ℌ = "Me"
H = "Funny"
print(ℌ == H) # True

# Explanation
# Python accepts non-ASCII but normalizes all identifiers.

import unicodedata
print(unicodedata.normalize('NFKC', "ℌ")) # "H"

