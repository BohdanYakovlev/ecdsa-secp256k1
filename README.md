# Elliptic Curve Digital Signature Algorithm (ECDSA) Implementation

This Go program implements the Elliptic Curve Digital Signature Algorithm (ECDSA) for generating and verifying digital signatures using elliptic curve cryptography.

## How to Run the Code

To run this console app, follow these steps:

1. Clone this repository to your local machine:

```
git clone <repository_url>
```

2. Navigate to the directory containing the Go code:

```
cd ecdsa-secp256k1
```

3. Ensure you have Go installed on your system. If not, you can download it from [here](https://golang.org/dl/).

4. Execute the Go code by running the following command:

```
go run main.go
```
5. The program will output `true` if the signature verification is successful; otherwise, it will output `false`.
## Overview

The program defines functions to perform the following tasks:

1. Generate a private key
2. Derive the corresponding public key
3. Sign a message using ECDSA
4. Verify the signature of a message using ECDSA

## Algorithms and Functions

The program uses the following algorithms and functions:
  - **Point Addition**: Adds two points on an elliptic curve.
  - **Point Doubling**: Doubles a point on an elliptic curve.
  - **Scalar Multiplication**: Multiplies a point on an elliptic curve by a scalar value.
  - **Signature Generation (ECDSA)**: Generates a digital signature for a given message using ECDSA.
  - **Signature Verification (ECDSA)**: Verifies the digital signature of a message using ECDSA.

## Conclusion

This Go program provides a robust implementation of the Elliptic Curve Digital Signature Algorithm (ECDSA) for generating and verifying digital signatures. By leveraging elliptic curve cryptography, it offers efficient and secure cryptographic operations suitable for various applications requiring digital signatures. The program's performance is optimized for the specified hardware specifications, making it suitable for use in real-world scenarios.

Feel free to reach out if you have any questions or encounter any issues while running the code.
