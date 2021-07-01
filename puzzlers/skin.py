print("👋\N{Emoji Modifier Fitzpatrick Type-1-2}") # 👋🏻
print("👋\N{Emoji Modifier Fitzpatrick Type-3}")   # 👋🏼
print("👋\N{Emoji Modifier Fitzpatrick Type-4}")   # 👋🏽
print("👋\N{Emoji Modifier Fitzpatrick Type-5}")   # 👋🏾
print("👋\N{Emoji Modifier Fitzpatrick Type-6}")   # 👋🏿

print("👋🏻" == "👋🏿") # False

# Normalization doesn't help
from unicodedata import normalize
print(normalize("NFKC", "👋🏻") == normalize("NFKC", "👋🏿")) # False

# Ignore skin tones does
print("👋🏻"[0] == "👋🏿"[0]) # True
