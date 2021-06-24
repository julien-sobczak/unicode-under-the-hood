print(len("✋")) # 1
print(len("🤚")) # 1

for c in "Hey 🤚!":
    print(c, hex(ord(c)))
    # H 0x48
    # e 0x65
    # y 0x79
    #   0x20
    # 🤚 0x1f91a
    # ! 0x21
