## Wallet API
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#testing">Testing</a></li>
  </ol>
</details>


## About The Project
You are responsible for managing the wallets of the players of an online casino, and you need to provide an API for getting and updating their account balances.

## Getting Started

### Prerequisites

Before this project can be run locally you will need to install [Go](https://golang.org/doc/install)

### Installation

To utilize the project, the following needs to be done:
1. Clone this repository
2. Install the dependencies, using the following command:
```
go mod tidy
```

## Usage

1. To run the project locally, use the following command:
```
make run
```
2. To retrieve the balance of a given wallet id, use Postman to make a GET request to the following URL:
```
http://localhost:{port}/api/v1/wallets/{wallet_id}/balance
```
3. To credit money on a given wallet id
   , use Postman to make a POST request to the following URL:
```
http://localhost:{port}/api/v1/wallets/{wallet_id}/credit
```
4. To debit money from a given wallet id, use Postman to make a POST request to the following URL:
```
http://localhost:{port}/api/v1/wallets/{wallet_id}/debit

```

## Testing
Tests can be run using the following command:
```
make test
```
