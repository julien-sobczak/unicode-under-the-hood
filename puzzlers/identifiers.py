𝛀 = "U+1D6C0"
Ω = "U+03A9"
print(𝛀, Ω) # U+03A9 U+03A9

# But not all non-ASCII characters are allowed:
ツ = "Letter in Unicode Character Database" # OK
# 🙂 = "Symbol in Unicode Character Database" # KO
# SyntaxError: invalid character '🙂' (U+1F642)
