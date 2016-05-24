# AlchemyAPI

 Simple AlchemyAPI client for Go.

## Usage
```
apiKey := "your api key"
config := alchemyapi.NewConfig(apiKey)
client := alchemyapi.New(config)
url := "http://www.dailymail.co.uk/sciencetech/article-2355833/Apples-iPhone-5-hated-handset--majority-people-love-Samsung-Galaxy-S4-study-finds.html"
input := client.NewURLGetRankedNamedEntitiesInput(url)
output, _ := client.URLGetRankedNamedEntities(input)
```

## Implemented
- [x] [URLGetRankedNamedEntities](https://www.alchemyapi.com/api/entity/urls.html)
- [x] [TextGetRankedNamedEntities](https://www.alchemyapi.com/api/entity/textc.html)
- [x] [URLGetAuthors](https://www.alchemyapi.com/api/authors-web-api)
- [x] [URLGetCombinedData](https://www.alchemyapi.com/api/combined/urls.html)
- [x] [URLGetText](https://www.alchemyapi.com/api/text/urls.html#text)

## Testing

 To manually run tests use your api key:

```
$ export ALCHEMYAPI_API_KEY=<your api key>
$ go test -v
```

# License

MIT