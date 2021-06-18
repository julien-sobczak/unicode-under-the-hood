#!/usr/bin/env python3
# Parse Scripts.txt and output an HTML file rendering a subset of characters in every script.

import re
import yaml

def extract_character_samples():
    p = re.compile('^(?P<start>[0-9A-Z]{4,5})(?:[.][.](?P<end>[0-9A-Z]{4,5}))?\s+;\s(?P<script>\w+)\s#\s+(?P<comment>.*)$')
    # Ex:
    # 0000..001F    ; Common # Cc  [32] <control-0000>..<control-001F>
    # 0020          ; Common # Zs       SPACE

    with open('Scripts.txt', 'r') as reader:
        scripts = {}
        for line in reader.readlines():

            # Detect script separator
            if line.startswith("# ================================="):
                new_script = True
                continue

            # New character (range)?
            m = p.match(line)
            if m:
                start = m.group("start")
                end = m.group("end")
                script = m.group("script")
                comment = m.group("comment")
                if script not in scripts: #
                    scripts[script] = []
                if "COMBINING" not in comment:
                    scripts[script].append(start)

    return scripts

def format_html(scripts):
    print("""<!doctype html>
<html lang="fr">
<head>
  <meta charset="utf-8">
  <title>Unicode</title>
  <link rel="preconnect" href="https://fonts.gstatic.com">
  <link href="https://fonts.googleapis.com/css2?family=Oi&display=swap" rel="stylesheet">
  <style type="text/css">
    body {
        background: #f5f5f5;
    }
    h1 {
        text-align: center;
    }
    .container {
        display: flex;
        flex-wrap: wrap;
        justify-content: center;
    }
    section {
        margin: 1rem;
        width: 10rem;
        height: 10rem;
        padding: 0.5rem;
        overflow: hidden;
        box-shadow: 1px 1px 5px grey;
        border-radius: 1rem;
        background: #fff475;
    }
    section h2 {
        margin: 0;
        text-align: center;
    }
    section ul {
        list-style-type: none;
        margin: 0.25rem;
        padding: 0;
    }
    section li {
        display: inline-block;
        font-family: 'Oi', Serif;
    }
  </style>
</head>
<body>
    <!--
        See for the list of characters per script
        https://www.unicode.org/Public/UCD/latest/ucd/Scripts.txt
    -->
    <h1>Unicode Scripts</h1>

    <div class="container">
""")

    for name, letters in scripts.items():
        print("""
        <section>
            <h2>%s</h2>
            <ul>""" % name)
        for letter in letters:
            print("                <li>&#x%s</li>" % letter)
        print("""            </ul>
        </section>""")

    print("""
    </div>
</body>
</html>
""")


scripts = extract_character_samples()
# print(yaml.dump(scripts))
format_html(scripts)



