[![Test](https://github.com/lmllrjr/svg-github-readme/actions/workflows/test.yaml/badge.svg)](https://github.com/lmllrjr/svg-github-readme/actions/workflows/test.yaml)

# svg-github-readme

svg-github-readme provides a simple solution of creating custom SVG's (e.g. headings, eye-catching content ...) in your github readme file.

## Animated Text
### Examples
![](https://svg-github-readme.vercel.app/api?type=animated_text&text=%23%23%20Animated%20Text&color=7ED7C1&font_size=50&font_family=Open%20Sans)

#### Full example query
```markdown
![](https://svg-github-readme.vercel.app/api?type=animated_text&text=helloworld&color=161A30&font_size=100&font_family=Arial)
```

#### Minimal example query
> [!IMPORTANT]
> The `text` query param is not mandatory, but nothing will be displayed when no text is provided.

```markdown
![](https://svg-github-readme.vercel.app/api?type=animated_text&text=helloworld)
```

### Query paremters
| Query Paremeter name | Example Value | Mandatory |
|----------------------|--------------:|----------:|
| type                 | animated_text | true      |
|                      |               |           |
| color                |        00ADD8 | false     |
| font_family          |       fantasy | false     |
| font_size            |           200 | false     |
| text                 | hello%20world | false     |

## Static Text
![](https://svg-github-readme.vercel.app/api?type=static_text&text=%23%23%20Static%20Text&color=FFA732&font_size=25&font_family=Open%20Sans)
### Examples
#### Full example query
```markdown
![](https://svg-github-readme.vercel.app/api?type=static_text&text=helloworld&color=FFA732&font_size=100&font_family=Courier)
```

#### Minimal example query
> [!IMPORTANT]
> The `text` query param is not mandatory, but nothing will be displayed when no text is provided.

```markdown
![](https://svg-github-readme.vercel.app/api?type=static_text&text=helloworld)
```

### Query paremters
| Query Paremeter name | Example Value | Mandatory |
|----------------------|--------------:|----------:|
| type                 | static_text   | true      |
|                      |               |           |
| color                |        FFA732 | false     |
| font_family          |       courier | false     |
| font_size            |           200 | false     |
| text                 | hello%20world | false     |

## Malformed query will result in an error svg
![](https://svg-github-readme.vercel.app/api?text=helloworld)

> `![](https://svg-github-readme.vercel.app/api?text=helloworld)`
