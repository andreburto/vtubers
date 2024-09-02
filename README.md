# VTubers

## About

This is a test API server.
It's a simple, simple server that will be an endpoint for async client testing.
Do not expect much from this.

## Using

The following should all be done from within the project directory.

### Build

```
go build
```

### Run

MacOS / Linux:
```
./vtubers
```

Windows:
```
.\vtubers.exe
```

## To Do

* Create tests.
* Add YAML config file with port and file info.
* Generate JSON output.
* Create more records in CSV files.
* Better error handling.

## Update Log

**2024-09-01:** Broke up files to make editing them easier, but can go further.
Added basic logic to update the files and add records. (Primo 90's webdev!)
Moved CSV files into `data` directory. 

**2024-04-13:** Initial commit. Server works with a small sample of information.
