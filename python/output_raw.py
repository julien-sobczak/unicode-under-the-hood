import os

with os.fdopen(1, 'wb') as stdout:
    stdout.write("Hello 👋\n".encode("utf-8"))
