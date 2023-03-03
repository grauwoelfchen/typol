# TypoL Converter

The world's most inaccurate converter.

It's not even powered by neural networks ;)


## Usage

### CLI

Do you need to type your Dvorak keyboards as Qwerty input? Try as follow:

```zsh
% typol convert -in Qwerty -out Dvorak "loadkeys dvorak"
nraet.fo ekrpat
```

Or, opposite?

```zsh
# same as -in Dvorak -out Qwerty
% typol convert "loadkeys /usr/share/keymaps/i386/qwerty/us.map.gz"
psahvdt; [f;o[;jaod[vdtmar;[g386[x,dokt[f;emareu\
```

See `typol convert --help` for usage.

### Server

```zsh
% make build:server
% ./dst/typol-server
```

```zsh
% curl localhost:3000/convert -F 'text=Hello'
TODO
```

## Development

```zsh
# run go generate (for stringer)
% make setup
```

```zsh
# run only unit tests in typol package
% make test

# run only integration tests
% make test:integration
```


## License

`AGPL-3.0-or-later`


```txt
TypoL
Copyright (C) 2023 Yasuhiro Яша Asaka

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.
```
