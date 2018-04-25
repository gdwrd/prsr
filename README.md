# Parser

## Installation

```bash
$> go get golang.org/x/net/html
$> cd prsr && go build
```

## Usage

```bash
$> ./prsr -h
Usage of ./crawler:
  -d int
    	Parser page level deep (default 3)
  -l int
    	Max size of parsed links (default 50)
  -tag string
    	Tag name you want to find (default "input")
  -uri string
    	URI to parse (default "https://sheremet.pw/")
```

### Example

```bash
$> ./prsr -uri=https://google.com/ -tag=img -d=2 -l=10

...........
URI:  https://www.google.com/calendar?tab=wc has 2 img
URI:  https://www.google.com.ua/intl/uk/options/ has 258 img
URI:  https://maps.google.com.ua/maps?hl=uk&tab=wl has 0 img
URI:  https://drive.google.com/?tab=wo has 2 img
URI:  https://google.com/ has 2 img
URI:  https://accounts.google.com/ServiceLogin?hl=uk&passive=true&continue=https://www.google.com/ has 2 img
URI:  https://www.google.com.ua/imghp?hl=uk&tab=wi has 1 img
URI:  https://mail.google.com/mail/?tab=wm has 2 img
URI:  https://www.youtube.com/?gl=UA&tab=w1 has 80 img
URI:  http://www.google.com.ua/history/optout?hl=uk has 1 img

LINKS PARSED:  10
```

## Tests

```bash
$> go test -cover
```