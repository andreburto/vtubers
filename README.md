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

# Queries

Connect to MongoDB `vtubers` database and try out this query:

```
db.getCollection("vtubers").find({"Name": { $in: ["Jelly Hoshiumi", "Kaneko Lumi"]}})
```

## To Do

* Create tests.
* Add YAML config file with port and file info.
* Generate JSON output.
* Create more records in CSV files.
* Better error handling.

## Update Log

**2024-12-15:** Crude loader works with exiting files. Now I can start converting everything over to MongoDB from CSV files.

**2024-09-01:** Broke up files to make editing them easier, but can go further.
Added basic logic to update the files and add records. (Primo 90's webdev!)
Moved CSV files into `data` directory. 

**2024-04-13:** Initial commit. Server works with a small sample of information.
