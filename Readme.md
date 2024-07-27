## Introduction [WIP]

SCM or system-call-montior is monitoring/auditing tool aimed at tracing system calls by attaching itself to a running application process. Under the hood it utilized the [eBPF technology](https://ebpf.io), which lets the program to directly interact with the kernel. After attaching to a process, it exports the metrics to prometheus, then it is consumed by Grafana to show dashboards and generate insights for the data.

## Aim

- Develop a CLI tool to attach the program to any process
- Make the tool interoperable for both linux and windows kernel
- Make the tool os independant, it should only rely on the underlying kernel to trace the calls
- Develop dashboards to show effective insights
- Develop auditing tool to generate audit reports of syscalls for a period
- Develop warning generation mechanism

## Running Locally for Development

- Install libbpf `sudo apt install libbpf-dev` for Debian/Ubuntu and `libbpf-devel` for Fedora
- Install clang `sudo apt install clang` and llvm
- Install Kernel Headers using `sudo apt install linux-headers-$(uname -r)`
- On Debian, you may also need `ln -sf /usr/include/asm-generic/ /usr/include/asm`.
- First generate a system call map for your kernel (Currently supports linux kernels) using `ausyscall $(uname -r) --dump > syscall.csv`
- Run `go get` to install all the required packages
- Run `make generate` to compile the eBPF code to object file and generate go scaffolding
- Run `make build` to compile the app
- Run `sudo ./main -n="name of the process you want to trace"` to run the compiled binary, you can also use `-id=<id of proc>` flag to explicilty provide the process id to track.

## Attaching scm to your docker containers

- Start your docker container using `docker-compose -f /path/to/your/docker-compose.yml up -d`
- Run the script `./docker-monitor /path/to/your/docker-compose.yml`

## Developing Environment

- Go version=1.22.0+
- Linux Kernel=5.7.0+
- [Kernel Config Requirements](https://github.com/iovisor/bcc/blob/master/docs/kernel_config.md) - mostly they are present, but if not, use this to debug

## Debug

- Install bpftool for debugging from source. This [blog](https://thegraynode.io/posts/bpftool_introduction/) covers everything.
- `bpftool prog list` to show all bpf programs loaded
- `bpftool map list` to show all BPF maps loaded
- `bpftool map dump <id>` to display contents of map in json, you can get the id of the map from above list command
