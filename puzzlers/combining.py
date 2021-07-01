print("à" == "à") # False

# Why? characters vs glyphs
print("\N{LATIN SMALL LETTER A WITH GRAVE}" ==
      "\N{LATIN SMALL LETTER A}\N{COMBINING GRAVE ACCENT}") # False
print("\u00E0" == "\u0061\u0300") # False

# Solution
from unicodedata import normalize
print(normalize("NFKC", "à") == normalize("NFKC", "à")) # True
print(normalize("NFKC", "\u00E0") == normalize("NFKC", "\u0061\u0300")) # True
