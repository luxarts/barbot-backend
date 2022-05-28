# Build
```
go build -o builds/backend cmd/main.go
```

# Run
1. Downlaod last binary from [releases](https://github.com/luxarts/barbot-backend/releases)
2. Download [`/data/drinks.json`](https://raw.githubusercontent.com/luxarts/barbot-backend/main/data/drinks.json) and [`/data/mixeddrinks.json`](https://raw.githubusercontent.com/luxarts/barbot-backend/main/data/mixedDrinks.json) and put them on the same folder as the binary
3. Create `.env` file with the following environment variables
  ```
  BARBOT_JWT_SECRET=<secret_base64Encoded>
  BARBOT_DB_FILEPATH_DRINKS=/drinks.json
  BARBOT_DB_FILEPATH_MIXEDDRINKS=/mixeddrinks.json
  ```
4. Run the binary
  ```
  ./backend
  ```
