# go-markitdown

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

A Go-based command-line tool and library for converting various file formats to Markdown.

## Features

- Convert DOCX (Word) documents to Markdown
- Convert HTML files to Markdown
- Convert PDF documents to Markdown (including code blocks)
- Convert XLSX (Excel) spreadsheets to Markdown tables
- Preserve basic formatting during conversion
- Lightweight and fast

## Installation

1. Make sure you have Go installed (version 1.20 or higher)
2. Install the tool:

```bash
go install github.com/chenjunqian/go-markitdown@latest
```

## Usage

Basic command structure:
```bash
markitdown [flags] input_file output_file
```

### Examples

1. Convert a Word document:
```bash
markitdown document.docx document.md
```

2. Convert a PDF with code blocks:
```bash
markitdown code.pdf code.md
```

3. Convert an Excel spreadsheet to a Markdown table:
```bash
markitdown data.xlsx table.md
```

4. Convert an HTML file:
```bash
markitdown page.html page.md
```

## Flags

- `-f`: Specify the input file path
- `-o`: Specify the output file path

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
