# [koyomi] -- 日本のこよみ

## Usage

```go
start, _ := koyomi.DateFrom("2019-05-01")
end, _ := koyomi.DateFrom("2019-05-31")
k, err := koyomi.NewSource(
    koyomi.WithCalendarID(koyomi.Holiday, koyomi.SolarTerm),
    koyomi.WithStartDate(start),
    koyomi.WithEndDate(end),
).Import()
if err != nil {
    return
}

csv, err := k.EncodeCSV()
if err != nil {
    return
}
io.Copy(os.Stdout, bytes.NewReader(csv))
//Output:
//"Date","Title"
//"2019-05-01","休日 (天皇の即位の日)"
//"2019-05-02","休日"
//"2019-05-02","八十八夜"
//"2019-05-03","憲法記念日"
//"2019-05-04","みどりの日"
//"2019-05-05","こどもの日"
//"2019-05-06","休日"
//"2019-05-06","立夏"
//"2019-05-21","小満"
```

[koyomi]: https://github.com/spiegel-im-spiegel/koyomi "spiegel-im-spiegel/koyomi: 日本のこよみ"
