# Spam

A program to generate random spam from a template. Inspired by James Curran post NCSS 2k16.

## Example

The template

```
Hello, {Harrison|Barker}
```

Could generate the responses

> Hello, Harrison

or

> Hello, Barker

You can also nest templates.

```
Hello, {how {are you|do you get to central}?|world!}
```

Which has three options:

> Hello, how are you?

> Hello, how do you get to central?

> Hello, world!
