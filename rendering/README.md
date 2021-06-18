# Using FreeType to shape text (Demo)

## Running

1. Start the virtual machine

```
$ vagrant up
$ vagrant ssh
```

2. Run the following commands

```
# Install requirements
sudo apt update
sudo apt install -y pkg-config clang libpng-dev libicu-dev libzip-dev freetype2-demos libfreetype6-dev bzip2 libbz2-dev

# Run the program
cd /vagrant
export CXXFLAGS=`pkg-config --cflags freetype2 libpng`
export LDFLAGS=`pkg-config --libs freetype2 libpng`
clang++ -o clfontpng -static $CXXFLAGS clfontpng.cc $LDFLAGS -licuuc -lz -lbz2

# Test with NotoColorEmoji font
./clfontpng NotoColorEmoji.ttf 🏄
# Check out.png
```
