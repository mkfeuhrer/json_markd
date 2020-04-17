# JSON-Markd

Go library for converting markdown lists to json.

Super useful for documenting APIs, maintaining todo list which can be used anywhere once converted to json.

## Installation

```go
go get -u "github.com/mkfeuhrer/json-markd"
```

Then use it using the ParseMarkdown functions -

```go
parser.SetTabSpaceValue(2) // This is important step, library will crash without this.
result, err := parser.ParseMarkdown("data/sample_api.md") // your markdown file path
if err != nil {
    logger.Log.Panic(err)
    return
}
fmt.Println(result)
```

### Markdown file

```markdown
- data : object
  - name : string
  - age : integer
  - income : double
  - vehicles: array
    - cars: object
      - name: string
      - price: double
  - apps : array
    - transport : array
      - gojek : string
      - uber : string
    - food : array
      - zomato : string
      - swiggy : string
- errors: object
  - type: string
```

### Resulting JSON

```json
{
  "data": {
    "name": "random string",
    "age": 0,
    "income": 0.0,
    "vehicles": [
      {
        "name": "random string",
        "price": 0.0
      }
    ],
    "apps": [
      ["random string", "random string"],
      ["random string", "random string"]
    ]
  },
  "errors": {
    "type": "random string"
  }
}
```

## Contraints

Since, this is a relatively new project, I built it for my personal use. Also, the logic code is a bit messy, will improve it.

It has some restrictions -

- You can only use nested lists.
- Ensure your markdown contains no empty lines
- Ensure for lists you use `-` , will include `*` and numbered list later.

**Note** - Check data folder for sample markdown

## Contribute

Feel free to contribute. Add new features or improve existing codebase.

Fork and make a PR :)

Star and share the library with your community.

<a href="https://www.buymeacoffee.com/chHAzigTb" target="_blank"><img src="https://cdn.buymeacoffee.com/buttons/default-orange.png" alt="Buy Me A Coffee" height="41" width="174"></a>

## Todos

- Add remaining tests
- Support all list formats - `*` and numbered list maybe

## Author

- [Mohit Khare](https://mohitkhare.me)
