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

You can dump the default templates to the current directory by calling `people2md dump-templates`.

To use changes templates pass -template-directory to people2md. There are five different templates:

* markdown.tmpl - provides the file structure
* phone.tmpl - used to print out phone numbers
* email.tmpl - used to print out email addresses
* address.tmpl - used to print out addresses
* person.tmpl - used to print out data of the person

To print out the call history and SMS you can provide SMS and calls using a backup file from SMS Backup and Restore (Pro).
This is an Android only app. Due to a lack of iPhone and knowledge of it's ecosystem there is no support for it. Feel free
to contribute.

## Required data calls

### Overview

- ResourceName
- Tags
- MainLinkName
- Groups

### List

- ResourceName
- Tags
- MainLinkName
- GroupKey
- CallData

## Upcoming feature

- [ ] allow dumping of properties

## History

|Version|Description|
|---|---|
|0.1.0|initial version|
