//go:build ignore

#include <linux/bpf.h>
#include <bpf/bpf_helpers.h>

struct {
    __uint(type, BPF_MAP_TYPE_HASH);
    __type(key, __u64);
    __type(value, __u64);
    __uint(max_entries, 600);
} sys_calls SEC(".maps");

struct {
    __uint(type, BPF_MAP_TYPE_HASH);
    __type(key, __u32);
    __type(value, __u16);
    __uint(max_entries, 100);
} proc_map SEC(".maps");

struct {
	__uint(type, BPF_MAP_TYPE_QUEUE);
	__uint(max_entries, 100000);
	__type(value, __u64);
} call_queue SEC(".maps");

struct syscall_fork_exit_t {
    unsigned short common_type;
    unsigned char common_flags;
    unsigned char common_preempt_count;
    int common_pid;

    int syscall_nr;
    long ret;
};

SEC("raw_tracepoint/sys_enter")
static void bpf_prog(struct bpf_raw_tracepoint_args *ctx){

    __u32 pid = bpf_get_current_pid_tgid() >> 32;
    __u16 *p;
    p = bpf_map_lookup_elem(&proc_map, &pid);
    if(p != 0){
        if(*p==1){
            __u64 call_id = ctx->args[1];
            __u64 count = 0;
            __u64 *calls;

            calls = bpf_map_lookup_elem(&sys_calls, &call_id);
            if(calls != 0){
                count = *calls;
            }
            count++;
            bpf_map_update_elem(&sys_calls, &call_id, &count, BPF_ANY);
            bpf_printk("Before Queue push: %ld", call_id);
            int queue_push = bpf_map_push_elem(&call_queue, &call_id, BPF_EXIST);
            bpf_printk("Queue push result: %d\n", queue_push);
        }
    }
}

SEC("tp/syscalls/sys_exit_clone")
static __always_inline void add_clone(struct syscall_fork_exit_t* ctx){
    __u32 pid = bpf_get_current_pid_tgid() >> 32;
    __u16 *p;
    p = bpf_map_lookup_elem(&proc_map, &pid);
    if(p != 0 && *p==1){
        long child_id = ctx->ret;
        bpf_printk("Parent: %u, Child: %ld", pid, child_id);
    }
}

char __license[] SEC("license") = "GPL";
