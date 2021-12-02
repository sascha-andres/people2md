# People 2 Markdown

This program takes the output of [goobook](https://gitlab.com/goobook/goobook) for contacts and groups and creates 
markdown files for [obsidian](https://obsidian.md).

To create the contacts.json you could run something like this:

    goobook dump_contacts > contacts.json

To create the groups.json you can run something like this:

    goobook dump_groups > groups.json

Setup of goobook is described on the homepage of the program.

Then you can run people2md to create the markdown files. people2md accepts the following
arguments:

* tags - a list of Google labels that will be applied in addition to the contact tag
* contacts - pass the path to contacts.json, defaulting to just the file name
* groups - pass the path to groups.json, defaulting to just the file name

## Upcoming feature

- [ ] copy file only if changed

## History

|Version|Description|
|---|---|
|0.1.0|initial version|
