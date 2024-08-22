# PAN-Rule-Usage-Report-Creator

A go utility tool for creating custom report set config for given security rule UUIDs for PAN-OS.

## Usage:

Pre-compiled binaries can be downloaded directly from the latest release (Link here: [Latest Release](https://github.com/enginy88/PAN-Rule-Usage-Report-Creator/releases/latest)). These binaries can readily be used on the systems for which they were compiled. Neither re-compiling any source code nor installing GO is not needed. In case there is no pre-compiled binary presented for your system, you can refer to the [Compilation](#compilation) section.

```shell
# Usage of ./PAN-Rule-Usage-Report-Creator:
  -panorama
        Can be set to use Panorama DB for creating report. (Optional)  [Default: Unset]
  -verbose
        Can be set to introduce verbose logs. (Optional)  [Default: Unset]
```

## Hints:

When using this program under Unix-like OSes like Linux or macOS, you can pipe commands. Here are some usage examples:

```shell
# Examples:

echo "a6b22ec2-e3a8-431c-beb7-a51d6f16943a" | ./PAN-Rule-Usage-Report-Creator -verbose

echo "a6b22ec2-e3a8-431c-beb7-a51d6f16943a" | ./PAN-Rule-Usage-Report-Creator > set-config.txt

echo "a6b22ec2-e3a8-431c-beb7-a51d6f16943a" | ./PAN-Rule-Usage-Report-Creator -panorama -verbose

echo "a6b22ec2-e3a8-431c-beb7-a51d6f16943a" | ./PAN-Rule-Usage-Report-Creator -panorama > set-config.txt

cat uuids.txt | ./PAN-Masterkey-Decoder -verbose

cat uuids.txt | ./PAN-Masterkey-Decoder > set-config.txt

cat uuids.txt | ./PAN-Masterkey-Decoder -panorama -verbose

cat uuids.txt | ./PAN-Masterkey-Decoder -panorama > set-config.txt
```

## Compilation:

If none of the pre-compiled binaries covers your environment, you can choose to compile from source by yourself. Here are the instructions for that:

```shell
git clone https://github.com/enginy88/PAN-Rule-Usage-Report-Creator.git
cd PAN-Rule-Usage-Report-Creator
go mod init
go mod tidy
make local # To compile for your own environment.
make # To compile for pre-selected environments. 
```

**NOTE:** To compile from the source code, GO must be installed in the environment. However, it is not necessary for to run compiled binaries! Please consult the GO website for installation instructions: [Installing Go](https://go.dev/doc/install)

## Why Golang?

Because it (cross) compiles into machine code! You can directly run ready-to-go binaries on Windows, Linux, and macOS. No installation, no libraries, no dependencies, no prerequisites... Unlike Bash/PowerShell/Python it is not interpreted at runtime, which drastically reduces runtime overhead compared to scripting languages. The decision to use a compiled language makes it run lightning fast with lower memory usage. Also, due to the statically typed nature of the Go language, it is more error-proof against possible bugs/typos.