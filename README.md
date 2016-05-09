# date2unix
Convert date string to UNIX timestamp.

# Installation
```
go get -u github.com/voidint/date2unix
```

# How to use?
- Output current time UNIX timestamp at system time zone.
```
$ date2unix --now
1462761855
```

- Output the specified time UNIX timestamp at UTC time zone.
```
$ date2unix --layout "2006-1-2 15:04:05" -p --utc "2016-10-10 01:01:01"
2016-10-10 01:01:01 => 1476061261
```

- Read date string form stdin
```
$ echo "2016/10/10 01:01:01" | date2unix
1476032461

or 

$ date2unix
2016/10/10 01:01:01
// Enter EOF by CTRL+D(Unix/Linux) or CTRL+Z(Win)
1476032461
```



