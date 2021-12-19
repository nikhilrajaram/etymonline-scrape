# etymonline-scrape
Program to scrape all etymologies on https://www.etymonline.com/

## How to use
Compile the program via
```bash
go build cmd/etymonline-scrape/scrape.go
```

And execute it with
```bash
./scrape
```

Once executed, the program will dump the scraped etymologies to `output/etymologies.json`.