# SilentSAM

```bash
    █████████   ███  ████                       █████     █████████    █████████   ██████   ██████
 ███░░░░░███ ░░░  ░░███                      ░░███     ███░░░░░███  ███░░░░░███ ░░██████ ██████ 
░███    ░░░  ████  ░███   ██████  ████████   ███████  ░███    ░░░  ░███    ░███  ░███░█████░███ 
░░█████████ ░░███  ░███  ███░░███░░███░░███ ░░░███░   ░░█████████  ░███████████  ░███░░███ ░███ 
 ░░░░░░░░███ ░███  ░███ ░███████  ░███ ░███   ░███     ░░░░░░░░███ ░███░░░░░███  ░███ ░░░  ░███ 
 ███    ░███ ░███  ░███ ░███░░░   ░███ ░███   ░███ ███ ███    ░███ ░███    ░███  ░███      ░███ 
░░█████████  █████ █████░░██████  ████ █████  ░░█████ ░░█████████  █████   █████ █████     █████
 ░░░░░░░░░  ░░░░░ ░░░░░  ░░░░░░  ░░░░ ░░░░░    ░░░░░   ░░░░░░░░░  ░░░░░   ░░░░░ ░░░░░     ░░░░░ 
                                                                                                                                                                                               
                                                                                                                                                                                           
```



SilentSAM is a tool designed to extract the Security Account Manager (SAM) database by leveraging raw disk access and parsing the NTFS Master File Table (MFT). By avoiding high-level system APIs, SilentSAM minimizes detection by security tools, offering a novel approach to credential extraction.

---

## Features

- **Raw Disk Access**: Reads disk sectors directly, bypassing OS-level file protections.
- **NTFS MFT Parsing**: Extracts file data using low-level NTFS structures.
- **Stealthy Operations**: Avoids generating high-level system logs or alerts.
- **Minimal Detection**: Bypasses conventional monitoring mechanisms like Windows APIs.

## Installation

To use SilentSAM, ensure you have [Go](https://golang.org/) installed on your system. Clone the repository and build the binary:

```bash
git clone https://github.com/Ryukk33/SilentSAM.git
cd SilentSAM
GOOS=windows go build SilentSAM.go
```

## Usage

SilentSAM requires elevated permissions to access raw disk data. Run the tool as an administrator.

```bash
SilentSAM.exe system-output-file sam-output-file
```

## How It Works

SilentSAM employs a low-level approach to bypass traditional file access restrictions:

1. **Locate the System Volume**:
   Identifies the volume containing the Windows installation.

2. **Read the Boot Sector**:
   Extracts essential NTFS parameters, including the MFT location.

3. **Parse the MFT**:
   Identifies records corresponding to the `SAM` and `SYSTEM` files.

4. **Extract Data**:
   Reads raw disk clusters directly to retrieve file content.

## Requirements

- **Operating System**: Windows
- **Permissions**: Administrator access is required.
- **Dependencies**:
  - [golang.org/x/sys/windows](https://pkg.go.dev/golang.org/x/sys/windows)
  - [github.com/t9t/gomft/mft](https://github.com/t9t/gomft)



## Disclaimer

This tool is intended for educational and research purposes only. The author is not responsible for misuse or damages caused by the use of this tool.

## Contributing

Contributions are welcome! If you'd like to contribute, feel free to fork the repository, make your changes, and submit a pull request.



## License

This project is licensed under the GNU GPL v3  License. See the [LICENSE](LICENSE) file for details.

## Author

- **Ryukk33**  
  [GitHub](https://github.com/Ryukk33) | [Twitter](https://twitter.com/ryukk33) | [Blog](https://blog.ryukk33.fr/)

## Screenshot

![](https://github.com/ryukk33/SilentSAM/blob/main/img/run.png)

---

## Acknowledgements

- Inspired by NTFS forensic research.
- Special thanks to the creators of [gomft](https://github.com/t9t/gomft) for the MFT parsing library.
