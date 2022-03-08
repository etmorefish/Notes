package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second)
		fmt.Println("Hello World")
	}
}

/*
 $ go build trace2.go
 $ GODEBUG=schedtrace=1000 ./trace2
SCHED 0ms: gomaxprocs=16 idleprocs=13 threads=5 spinningthreads=1 idlethreads=0 runqueue=0 [0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
Hello World
SCHED 1004ms: gomaxprocs=16 idleprocs=16 threads=6 spinningthreads=0 idlethreads=4 runqueue=0 [0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
Hello World
SCHED 2013ms: gomaxprocs=16 idleprocs=16 threads=6 spinningthreads=0 idlethreads=4 runqueue=0 [0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
Hello World
SCHED 3021ms: gomaxprocs=16 idleprocs=16 threads=6 spinningthreads=0 idlethreads=4 runqueue=0 [0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
Hello World
SCHED 4030ms: gomaxprocs=16 idleprocs=16 threads=6 spinningthreads=0 idlethreads=4 runqueue=0 [0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
Hello World


SCHED：调试信息输出标志字符串，代表本行是goroutine调度器的输出；
0ms：即从程序启动到输出这行日志的时间；
gomaxprocs: P的数量，本例有2个P, 因为默认的P的属性是和cpu核心数量默认一致，当然也可以通过GOMAXPROCS来设置；
idleprocs: 处于idle状态的P的数量；通过gomaxprocs和idleprocs的差值，我们就可知道执行go代码的P的数量；
threads: os threads/M的数量，包含scheduler使用的m数量，加上runtime自用的类似sysmon这样的thread的数量；
spinningthreads: 处于自旋状态的os thread数量；
idlethread: 处于idle状态的os thread的数量；
runqueue=0： Scheduler全局队列中G的数量；
[0 0]: 分别为2个P的local queue中的G的数量。

*/
