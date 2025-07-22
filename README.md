# PasswordMaker

**PasswordMaker** is a lightweight and secure password generator built in Go, featuring a simple graphical user interface (GUI). Just enter the domain (e.g., `www.baidu.com`, `QQ`, etc.) and it will generate a strong password containing uppercase, lowercase letters, numbers, and special characters — suitable for most websites.

## Features

- 🔐 Strong password generation based on domain input
- 🖥️ User-friendly GUI
- 📦 Portable `.exe` build (no installation required)
- ⚙️ Developed in Go using the [Fyne](https://github.com/fyne-io/fyne) GUI framework

## Getting Started

Clone the repo and run:

```bash
fyne package -os windows --icon keys.ico --name PasswordMaker