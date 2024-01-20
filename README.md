# rpgsave-decode

A tool to decode RPG Maker MV save files.

`.rpgsave` files are json files compressed with [lz-string](https://pieroxy.net/blog/pages/lz-string/index.html) to base64.

## Usage

Drag and drop a `.rpgsave` file onto `rpgsave-decode.exe` or run the following command:

```
rpgsave-decode <input>
```

> `rpgsave-decode` will automatically output to a `.json` file of the same name. **(indented for readability)**

Releases can be found [here](https://github.com/gteditor99/rpgsave-decrpyt/releases).

## Building

`rpgsave-decode` is a `golang` project. To build it, you will need to have `golang` installed.

Clone the repository:

```bash
git clone https://github.com/gteditor99/rpgsave-decrpyt.git
```

Build the project:

```bash
go build
```

## Misc

`rpgsave-decode` depends on [@daku10/go-lz-string](https://github.com/daku10/go-lz-string) for decompression.

## License

`rpgsave-decode` is licensed under the `MIT` license. See `LICENSE` for more information.
