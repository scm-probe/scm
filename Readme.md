## Introduction [WIP]

SCM or system-call-montior is monitoring/auditing tool aimed at tracing system calls by attaching itself to a running application process. Under the hood it utilized the [eBPF technology](https://ebpf.io), which lets the program to directly interact with the kernel. After attaching to a process, it exports the metrics to prometheus, then it is consumed by Grafana to show dashboards and generate insights for the data.

## Aim

- Develop a CLI tool to attach the program to any process
- Make the tool interoperable for both linux and windows kernel
- Make the tool os independant, it should only rely on the underlying kernel to trace the calls
- Develop dashboards on Grafana to show effective insights

## Running Locally for Development

- First generate a system call map for your kernel (Currently supports linux kernels) using `ausyscall $(uname -r) --dump > syscall.csv`
- Run `go get` to install all the required packages
- Run `make generate` to compile the eBPF code to object file and generate go scaffolding
- Run `make build` to compile the app
- Run `sudo ./main -n="name of the process you want to trace"` to run the compiled binary

## Developing Environment

- Go version=1.22.0+
- Linux Kernel=5.7.0+
- [Kernel Config Requirements](https://github.com/iovisor/bcc/blob/master/docs/kernel_config.md) - mostly they are present, but if not, use this to debug

## Debug

- Install bpftool for debugging from source. This [blog](https://thegraynode.io/posts/bpftool_introduction/) covers everything.
- `bpftool prog list` to show all bpf programs loaded
- `bpftool map list` to show all BPF maps loaded
- `bpftool map dump <id>` to display contents of map in json, you can get the id of the map from above list command
