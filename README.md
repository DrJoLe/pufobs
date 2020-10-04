# pufobs

## Usage

```
pufobs is a small tool to list and download "DAS PODCAST UFO" episodes.

Usage:
  pufobs [flags]
  pufobs [command]

Available Commands:
  count       Print the number of currently available "DAS PODCAST UFO" episodes
  current     Print the name of the current "DAS PODCAST UFO" episode
  download    Download a "DAS PODCAST UFO" episode
  help        Help about any command
  list        List all available "DAS PODCAST UFO" episodes

Flags:
  -h, --help   help for pufobs

Use "pufobs [command] --help" for more information about a command.
```

## Options
`count`, `current` and `list` do not take any arguments.

### download
`download` takes at most 2 arguments. You can leave both blank, `filepath` blank or provide both.
The default value for `title` is the title of the latest episode. The default value for `filepath` is the name of the hosted media file on [podcast-ufo.fail](http://podcast-ufo.fail/) joined to the current working directory.
#### title
The title of the "DAS PODCAST UFO" episode to be downloaded. You do not have to provide the full title. The first matching title will be selected.
##### Examples:
* U -> UFO001 Prolog
* UFO -> UFO001 Prolog
* UFO0 -> UFO001 Prolog
* UG -> UGO081 Freestyle Karate
* UFO0123 -> UFO123 Arrrrr!
#### filepath
The filepath where the downloaded file should be saved at.
