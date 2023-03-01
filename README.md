# TypoL

A text exchanger for your key inputs.

e.g. Qwerty -> Dvorak, Colemak -> Neo etc.


## Usage

See `typol convert --help`.

```zsh
% typol convert "Hoi Zäme"
% typol convert -to Dvorak -in "Hoi Zäme"
...
```

## Development

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
