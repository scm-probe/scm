//go:build ignore

#include <linux/bpf.h>
#include <bpf/bpf_helpers.h>

struct {
    __uint(type, BPF_MAP_TYPE_HASH);
    __type(key, __u64);
    __type(value, __u64);
    __uint(max_entries, 600);
} sys_calls SEC(".maps");

SEC("raw_tracepoint/sys_enter")
static __always_inline void bpf_prog(struct bpf_raw_tracepoint_args *ctx){
    unsigned long call_id = ctx->args[1];
    __u64 count = 0;
    __u64 *calls;

    calls = bpf_map_lookup_elem(&sys_calls, &call_id);
    if(calls != 0){
        count = *calls;
    }
    count++;
    bpf_map_update_elem(&sys_calls, &call_id, &count, BPF_ANY);
}

char __license[] SEC("license") = "Dual MIT/GPL";