Sure! Here's a well-structured **README** description for your package:

---

# **formatText**

## 🚀 Overview
`formatText` is a lightweight Go package designed to provide text formatting utilities, including bold text conversion using Unicode characters. This package is ideal for developers who need **consistent, reusable** text styling across multiple projects.

## ✨ Features
- Convert plain text to **bold Unicode characters**.
- Convert number string to number string with decimals.
- Easily reusable in any Go project.
- Best used when passing test over HTTP requests.
- Designed for **simplicity and efficiency**.

## 🔧 Installation
To install the package, run:
```sh
go get github.com/EmekaOkonkwo101/formatText
```

## 📌 Usage
Import the package and start using its formatting functions:
```go
package main

import (
	"fmt"
	"github.com/EmekaOkonkwo101/formatText/formatPackages"
)

func main() {
	formattedText := formatPackages.GetTextInBoldFormat("Hello, World!")
	fmt.Println(formattedText) // Output: 𝗛𝗲𝗹𝗹𝗼, 𝗪𝗼𝗿𝗹𝗱!
}
```

## 🔄 Updating the Package
To fetch the latest version:
```sh
go get -u github.com/EmekaOkonkwo101/formatText
```

## 🤝 Contributing
Feel free to fork the repository, submit issues, or create pull requests to improve functionality.

## 📜 License
This package is licensed under [MIT License](LICENSE).

