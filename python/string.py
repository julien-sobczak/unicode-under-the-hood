import unicodedata
import re

def check_length():
    # Ʃ (https://unicode-table.com/en/01A9/)
    # This single character is encoded using at least two bytes
    l = len("Ʃ") # \u01a9
    if l == 1:
        print("👍 Unicode character length")
    else:
        print("👎 Unicode character length. Got: %d" % l)

    # Test with a combining character
    l = len("á") # a\u0301
    if l == 1:
        print("👍 Unicode character length with combining character")
    else:
        print("👎 Unicode character length with combining character. Got: %d" % l)

    # Test with combining characters
    sequence1 = "Voil\N{LATIN SMALL LETTER A WITH GRAVE}"
    sequence2 = "Voil\N{LATIN SMALL LETTER A}\N{COMBINING GRAVE ACCENT}"
    if len(sequence1) == len(sequence2):
        print("👍 Unicode sequence length")
    else:
        print("👎 Unicode sequence length")

    norm1 = unicodedata.normalize("NFKC", sequence1)
    norm2 = unicodedata.normalize("NFKC", sequence2)
    if len(norm1) == len(norm2):
        print("👍 Unicode sequence length using normalization")
    else:
        print("👎 Unicode sequence length using normalization")


def check_comparison():
    sequence1 = "Voil\N{LATIN SMALL LETTER A WITH GRAVE}"
    sequence2 = "Voil\N{LATIN SMALL LETTER A}\N{COMBINING GRAVE ACCENT}"

    if sequence1 == sequence2:
        print("👍 Unicode string comparison using ==")
    else:
        print("👎 Unicode string comparison using ==")

    # The Unicode Standard defines four normalization forms.
    # Roughly speaking, NFD and NFKD decompose characters where possible,
    # while NFC and NFKC compose characters where possible.
    # For more information, see Unicode Standard.
    def NFD(s):
        return unicodedata.normalize('NFD', s)

    if NFD(sequence1) == NFD(sequence2):
        print("👍 Unicode string comparison using unicodedata module")
    else:
        print("👎 Unicode string comparison using unicodedata module")


def check_case_folding():
    sequence1 = "Voil\u00E0" # à
    sequence2 = "Voil\u00C0" # À

    res = sequence1.upper()
    if res != sequence2:
        print("👍 Unicode case folding using strings.ToUpper()")
    else:
        print("👎 Unicode case folding using strings.ToUpper()")

    res = sequence2.lower()
    if res != sequence1:
        print("👍 Unicode case folding using strings.ToLower()")
    else:
        print("👎 Unicode case folding using strings.ToLower()")

    if sequence1.casefold() == sequence2.casefold():
        print("👍 Unicode case folding using strings.EqualFold()")
    else:
        print("👎 Unicode case folding using strings.EqualFold()")


def check_regex():
    s = "100 µAh 10 mAh"
    res = re.findall(r'\\d+ \\wAh', s)
    if len(res) == 2:
        print("👍 Unicode regex using \\w in bytes")
    else:
        print("👎 Unicode regex using \\w in bytes")

    # If the regex pattern is in bytes, \w matches any alphanumeric character
    # ([a-zA-Z0-9_]).
    # If the regex pattern is a string, \w matches all the characters marked
    # as letters in the Unicode database.

    res = re.findall("\\d+ \\wAh", s)
    if len(res) == 2:
        print("👍 Unicode regex using \\w in string")
    else:
        print("👎 Unicode regex using \\w in string")


if __name__ == "__main__":
    check_length()
    check_comparison()
    check_case_folding()
    check_regex()

# Output:
# ------
# 1. Length
# 👍 Unicode character length
# 👎 Unicode character length with combining character. Got: 2
# 👎 Unicode sequence length
# 👍 Unicode sequence length using normalization
# 2. Text Comparison
# 👎 Unicode string comparison using ==
# 👍 Unicode string comparison using unicodedata module
# 3. Text Case Folding
# 👍 Unicode case folding using strings.ToUpper()
# 👍 Unicode case folding using strings.ToLower()
# 👍 Unicode case folding using strings.EqualFold()
# 4. Regex
# 👎 Unicode regex using \w in bytes
# 👍 Unicode regex using \w in string
