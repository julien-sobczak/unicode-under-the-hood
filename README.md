# Unicode Under the Hood

This repository is the companion of the blog post [Unicode for Curious Developers Loving Emojis and Code 😉](https://www.juliensobczak.com/inspect/2021/06/19/unicode-for-developers.html).

It contains short programs written to illustrate different aspects of Unicode.

```
encodings/       # Go basic implementation of Unicode encoding algorithms
go/              # Demo programs written in Go
haskell/         # Demo programs written in Haskell
java/            # Demo programs written in Java
php/             # Demo programs written in PHP
python/          # Demo programs written in Python
rendering/       # C program to illustrate the text rendering
representation/  # Basic implementation of a string type in Python
```

Directories named after a programming language contains some of the following files:

* `hello.*`: Test using only ASCII characters.
* `hello_UTF-*.*`: Test using various Unicode encodings.
* `voila.*`: Test using non-ASCII characters.
* `string.*`: Test various string manipulation functions.
* `emojis.*`: Test the support of emojis.
* `representation.*`: Expose the internal string representation subtleties.

