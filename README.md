# [koyomi] -- 日本のこよみ

[![check vulns](https://github.com/spiegel-im-spiegel/koyomi/workflows/vulns/badge.svg)](https://github.com/spiegel-im-spiegel/koyomi/actions)
[![lint status](https://github.com/spiegel-im-spiegel/koyomi/workflows/lint/badge.svg)](https://github.com/spiegel-im-spiegel/koyomi/actions)
[![GitHub license](https://img.shields.io/badge/license-Apache%202-blue.svg)](https://raw.githubusercontent.com/spiegel-im-spiegel/koyomi/master/LICENSE)
[![GitHub release](https://img.shields.io/github/release/spiegel-im-spiegel/koyomi.svg)](https://github.com/spiegel-im-spiegel/koyomi/releases/latest)

「[国立天文台 天文情報センター 暦計算室](http://eco.mtk.nao.ac.jp/koyomi/)」より日本の暦情報を取得する [Go 言語]用パッケージです。 Google Calendar を経由して取得しています。

取得可能な情報は以下の通りです。

```go
const (
    Holiday   CalendarID = iota + 1 //国民の祝日および休日
    MoonPhase                       //朔弦望
    SolarTerm                       //二十四節気・雑節
    Eclipse                         //日食・月食・日面経過
    Planet                          //惑星現象
)
```

## Usage

```go
start, _ := koyomi.DateFrom("2019-05-01")
end, _ := koyomi.DateFrom("2019-05-31")
k, err := koyomi.NewSource(
    koyomi.WithCalendarID(koyomi.Holiday, koyomi.SolarTerm),
    koyomi.WithStartDate(start),
    koyomi.WithEndDate(end),
).Get()
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

## Links

- [国立天文台 天文情報センター 暦計算室](https://eco.mtk.nao.ac.jp/koyomi/)
    - [今月のこよみ powered by Google Calendar - 国立天文台暦計算室](https://eco.mtk.nao.ac.jp/koyomi/cande/calendar.html)
- [日本の暦情報を取得するパッケージを作ってみた — リリース情報 | text.Baldanders.info](https://text.baldanders.info/release/2020/05/koyomi/)

[Go 言語]: https://golang.org/ "The Go Programming Language"
[koyomi]: https://github.com/spiegel-im-spiegel/koyomi "spiegel-im-spiegel/koyomi: 日本のこよみ"
