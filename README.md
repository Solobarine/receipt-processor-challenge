# Receipt Processing API

This project is a Receipt Processing API built in Go, using the Gin Web Framework. The API allows you to process receipts, calculate points based on specific rules, and retrieve receipt details.

## Features

- Process Receipts: Submit receipts with relevant details, including items and prices.
- Calculate Points: Earn points based on receipt attributes and defined rules.
- View Receipts(Additional): Retrieve all processed receipts or query points for a specific receipt by ID.

## Setup

### Prerequisites

- Go 1.17+ installed on your machine.

### Installation

1. Clone this repository:

```bash
git clone https://github.com/solobarine/receipt-processing-api.git
cd receipt-processing-api
```

2. Install the required Go modules:

```bash
go mod tidy
```

3. Run the Server
   To start the server, run:

```bash
go run main.go
```

The server will start on localhost:8080.

## Endpoints

1. Get All Receipts

- Endpoint: GET /receipts
- Description: Retrieves a list of all processed receipts.
- Response: JSON array of all receipts.

2. Process a New Receipt

- Endpoint: POST /receipts/process
- Description: Processes a new receipt by saving its details and generating a unique ID.
- Request Body: JSON object representing the receipt.

```json
{
  "retailer": "Retailer Name",
  "purchaseDate": "YYYY-MM-DD",
  "purchaseTime": "HH:MM",
  "items": [
    {
      "shortDescription": "Item Description",
      "price": "Item Price"
    }
  ],
  "total": "Total Price"
}
```

- Response: JSON object with the receipt ID.

```json
{
  "id": "unique-receipt-id"
}
```

3. Get Points for a Receipt

- Endpoint: GET /receipts/:id/points
- Description: Retrieves the points for a specific receipt based on predefined rules.
- Response: JSON object with the calculated points.

```json
{
  "points": "calculated points"
}
```

## Rules for Points Calculation

- **Alphanumeric Retailer Name:** If the retailer's name is alphanumeric, points are awarded.
- **Round Dollar Total:** If the total is a round dollar amount, 50 points are awarded.
- **Multiple of 0.25 Total:** If the total is a multiple of 0.25, 25 points are awarded.
- **Item Pairs:** Every pair of items earns 5 points.
- **Description Multiple of 3:** If an item description length is a multiple of 3, points equal to 20% of the item's price are awarded.
- **Odd Purchase Day:** If the purchase day is odd, 6 points are awarded.
- **Purchase Time Range:** If the purchase time is between 14:00 and 16:00, 10 points are awarded.

## Example Usage

1. Process a Receipt:

```bash
curl -X POST "http://localhost:8080/receipts/process" -H "Content-Type: application/json" -d '{
  "retailer": "Walmart",
  "purchaseDate": "2024-11-12",
  "purchaseTime": "15:30",
  "items": [{"shortDescription": "Apple", "price": "1.00"}],
  "total": "1.00"
}'
```

2. Get Points for a Receipt:

```bash
curl -X GET "http://localhost:8080/receipts/{id}/points"
```

Replace {id} with the ID returned in the POST /receipts/process response.

## License

This project is licensed under the MIT License.
