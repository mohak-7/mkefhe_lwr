# mkefhe_lwr

This is a Go implementation of Multi-Key Extension Fully Homomorphic Encryption using LWR by Mohakjot Dhiman during SPARK 2025 Summer Internship at Department of Mathematics, IIT Roorkee. This encryption scheme is proposed by Mansi Goyal and Dr. Aditi Gangopadhyay of Department of Mathematics at IIT Roorkee.

## Project Structure

```
mkefhe_lwr/
├── mkefhe/         # Project-specific packages
├── utils/          # Utility and helper functions
├── main.go         # Main application entry point
├── go.mod          # Go module definition
├── go.sum          # Go module checksums
└── README.md       # Project documentation
```

## Getting Started

### Prerequisites

- Go 1.16 or newer

### Installation

1. **Clone the repository:**
    ```
    git clone https://github.com/mohak-7/mkefhe_lwr.git
    cd mkefhe_lwr
    ```

2. **Download dependencies:**
    ```
    go mod tidy
    ```

## Usage

To run the application:

```
go run main.go
```

To build and execute:

```
go build -o mkefhe_lwr
./mkefhe_lwr
```

## Contributing

Pull requests and issues are welcome! Please fork the repository and submit your changes via a pull request.

## Author

[mohak-7](https://github.com/mohak-7)

## References

Mansi Goyal and Aditi Kar Gangopadhyay. 2025. Key Extension: Multi-Key
FHE Utilizing LWR. In *ACM Asia Conference on Computer and Communications Security (ASIA CCS ’25), 2025, Hanoi, Vietnam.* ACM,
New York, NY, USA, 14 pages. https://doi.org/10.1145/3708821.3736209