# Go-Reloaded Project

![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)

## Table of Contents

- [Project Overview](#project-overview)
- [Objectives](#objectives)
- [Introduction](#introduction)
- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
- [Examples](#examples)
- [Testing](#testing)
- [Contributing](#contributing)
- [License](#license)
- [Acknowledgements](#acknowledgements)
- [Contact](#contact)

---

## Project Overview

This project is a **simple text completion/editing/auto-correction tool** built using Go. Leveraging previously developed functions from an existing repository, the tool processes input text files and applies a series of modifications based on embedded commands. The primary goal is to enhance text readability and correctness automatically.

The project emphasizes good coding practices and includes comprehensive unit tests to ensure reliability. Additionally, it fosters a collaborative learning environment where auditors (students) review each other's work, promoting mutual growth and understanding.

---

## Objectives

- **Leverage Existing Functions:** Utilize functions from a previous repository to build the auto-correction tool.
- **Command-Based Text Modification:** Implement a system where embedded commands in the text dictate specific modifications.
- **Collaborative Auditing:** Engage in peer-review processes where students audit each other's code, enhancing learning outcomes.
- **Comprehensive Testing:** Develop and maintain unit tests to validate functionality and ensure robustness.
- **Adherence to Best Practices:** Ensure the codebase follows Go's best practices for readability, efficiency, and maintainability.

---

## Introduction

The **Text Auto-Correction Tool** is designed to process text files and apply a variety of transformations based on embedded commands. Written in Go, the tool emphasizes clean code, efficiency, and ease of use. By interpreting specific commands within the text, the tool can perform actions such as case conversion, numerical base transformation, punctuation formatting, and more, enhancing text readability and correctness automatically.

This project is particularly useful for tasks like automated proofreading, text formatting, and preparing documents for publication or presentation.

---

## Features

- **Hexadecimal to Decimal Conversion:** Transforms hexadecimal numbers to their decimal equivalents.
- **Binary to Decimal Conversion:** Converts binary numbers to decimal format.
- **Case Transformations:** Supports converting words to lowercase, uppercase, or capitalized forms.
- **Punctuation Formatting:** Adjusts spacing and grouping of punctuation marks for improved readability.
- **Article Correction:** Automatically adjusts articles ("a" to "an") based on the following word.
- **Batch Word Transformations:** Applies transformations to a specified number of preceding words.
- **Robust Command Parsing:** Accurately interprets and applies embedded commands within the text.
- **Comprehensive Testing:** Includes unit tests to ensure all functionalities work as intended.

---
The **Go-Reloaded** is a Go-based application that processes text files to apply various modifications based on embedded commands. It enhances text readability and correctness by performing tasks such as case conversion, numerical base transformations, and punctuation formatting automatically.

## Installation

1. **Clone the Repository**

   ```bash
   git clone https://github.com/TerZoro/go-reloaded.git
   cd go-reloaded
   ```

## Usage

The tool operates by taking two command-line arguments:

1. Input File: The path to the text file that requires modifications.
2. Output File: The path where the modified text will be saved.

### Syntax:

   ```bash
   go run . <input_file> <output_file>
   ```
Or, if you've built the executable:

   ```bash
   ./autocorrect <input_file> <output_file>
   ```
