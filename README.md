# gin-error-handler

Automatically returns JSON when you don't want to bother with the err!=nil, with Gin framework.

## Usage

Copy the `eh.go` and `preset.go` file to wherever you want.

Define your Data Transfer Object with "eh" tags.

There are 3 meaningful tags, others stands for the default value:

1. `err` - the field with `err` tag will be returned as a string, its value equals to `err.Error()`
2. `nil` - the field with `nil` tag will be omitted
3. `pre: ` - set the field with value in `Preset` type, e.g. `pre: CodeOK` will set the field with `Preset.CodeOK` (200,
   int64). Be aware of the spelling: `p`,`r`,`e`,`colon`,`space`. The type of the original field and the Preset must be
   the same, otherwise it will panic.

## Example