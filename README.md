# Currency Converter

[comment]: <> (![alt text]&#40;./gophers.jpg&#41;)

This is a simple web service that converts Nigerian Naira(NGN), Ghanaian Cedis(GHS), and Kenyan Shillings (KSH) from
anyone to any other.

## Usage

1. Clone the repo

   ```sh

   git clone https://github.com/karokojnr/currency-converter

   ```

2. Install go packages

   ```sh

   go mod download

   ```
3. To make tests <br/>
   You can choose to edit `conversion_test.go` and run:
   ```sh
   
   go test
   
   ```

4. To run the app:

   ```sh

   go run main.go

   ```
5. On your browser open:

   ```sh

   localhost:4000

   ```

## Dependency Used

`github.com/joho/godotenv` <br/>

The acronymns to `currencies` and their respective `currency rates` ought to be secretive and thus they need to be
stored as environment variables, thus the use of this dependency to load them from an `.env` file.

## Assumptions

   ```sh
   1 NGN = 0.014 GHS
   1 NGN = 0.26 KSH
   
   1 GHS = 69.46 NGN
   1 GHS = 18.23 KSH
   
   1 KSH = 3.81 NGN
   1 KSH = 0.055 GHS
   ```

- Cannot convert a negative value or a zero value.


