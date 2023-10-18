# ctcc
see your tcc db

This is a simple port of [tcc](https://github.com/Mac-Nerd/tcctool/blob/main/tcctool.sh) to go with a few extra features.

## Usage
```sh
./ctcc -h             
Usage of ./ctcc:
  -csv
        Output to CSV
  -file string
        File name to output to (default "ctcc.csv")
  -json
        Output to JSON
  -sal
        Output for sal to pick it up
  -version
        Print version and exit
```

For this to run properly you will need to provide Full-Disk Access to the binary or whatever is executing it.

## Sal
If you want to add this extra information as a [checkin module](https://github.com/salopensource/sal/wiki/Checkin-and-checkin-modules) for [Sal](https://github.com/salopensource/sal) you can use the `-sal` flag to output the data in the correct format.

This will create a file `/usr/local/sal/tcc_results.json` that the [tcc_returner.py](package/payload/usr/local/sal/checkin_modules/tcc_returner.py) will pick up and send to Sal.

Note: This program takes about 10 seconds to run on my machine, so it might be worth running it independently of the checkin module to avoid slowing down the checkin process.

## CSV
Similar to the tool that inspired this, you can output the data to a CSV file with the `-csv` flag. This will create a file called `ctcc.csv` in the current directory as well as `MDM-ctcc.csv` for the MDM overrides.
