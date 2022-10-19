### 调度算法

> 操作系统的进程调度算法，以及优缺点分析，包括 FIFO 算法、最短任务优先算法、轮转算法、多级反馈队列算法、彩票调度算法，以及多处理器的调度算法。只总结了各算法的原理，没有介绍 Linux 调度算法的具体实现。

#### 1. 调度指标

> 进程调度算法都遵循哪些指标。目前只考虑周转时间和响应时间。

#### 2. FIFO

> FIFO 就是先到先出调度 (First In First Out, FIFO)，有时也叫先到先服务调度 (First Come First Served, FCFS)。优点：很简单，易于实现。缺点：当先到达的任务耗时较长时，FIFO 调度算法的平均周转时间会很长。

![alt 属性文本](https://blog.cdn.updev.cn/2020-07-26-131504.jpg)

#### 3. 最短任务优先 (SJF)

> 最短任务优先 (Shortest Job First, SJF) 算法的策略：先运行最短的任务，然后是次短的任务，如此下去。在所有任务同时到达的假设下，SJF 确实是最优调度算法。但任务有先有后的时候，SJF 就不理想了，如下图所示。

![alt 属性文本](https://blog.cdn.updev.cn/2020-07-26-131522.jpg)

#### 4. 最短完成时间优先 (STCF)

> 为了解决 SJF 的问题，引入了最短完成时间优先 (Shortest Time-to-Completion First, STFT) 算法，前提是操作系统需要允许任务抢占。STCF 算法在 SJF 的基础上添加抢占，又称抢占式最短作业优先 (Preemptive Shortest Job First, PSJF)。每当新工作进入系统时，就会确定剩余工作和新工作中，谁的剩余时间最少，然后调度该工作。

![alt 属性文本](https://blog.cdn.updev.cn/2020-07-26-131532.jpg)

#### 5. 轮转算法

> 为了解决 STCF 算法响应时间长的问题，引入了轮转调度 (Round-Robin, RR)。轮转调度的思想：RR 在一个时间片 (time slice，有时称为调度因子，scheduling quantum) 内运行一个工作，然后切换到运行队列中的下一个任务，而不是运行一个任务直到结束。RR 有时被称为时间切片，时间片长度必须是时钟中断周期的倍数。
> ![alt 属性文本](https://blog.cdn.updev.cn/2020-07-26-131546.jpg)

> 时间片长度对于 RR 是至关重要的。时间片越短，RR 在响应时间上表现越好。然而，时间片太短是有问题的：突然上下文切换的成本将影响整体性能。因此，系统设计者需要权衡时间片的长度，使其足够长，以便摊销上下文切换成本，而又不会使系统不及时响应。

> 注：上下文切换的成本不仅仅来自保存和恢复少量寄存器的操作系统操作。程序运行时，它们在 CPU 高速缓存、TLB、分支预测器和其他片上硬件中建立了大量的状态。切换到另一个工作会导致此状态被刷新，且与当前运行的作业相关的新状态被引入，可能导致显著的性能成本

#### 6. 结合 I/O

> 当运行的程序在进行 I/O 操作的时候，在 I/O 期间不会使用 CPU，但它被阻塞等待 I/O 完成，这时调度程序应该在 CPU 上安排另一项工作。而在 I/O 完成时，会产生中断，操作系统运行并将发出 I/O 的进程从阻塞状态移回就绪状态。当然，它甚至可以决定在那个时候运行该项工作。操作系统应该如何处理每项工作？假设有两项工作 A 和 B，每项工作需要 50ms 的 CPU 时间。但是 A 先运行 10ms，然后发出 I/O 请求（假设 I/O 每个都需要 10ms），而 B 只是使用 CPU 50ms，不执行 I/O。调度程序先运行 A，然后运行 B。

> 左图所示的调度是非常糟糕的。常见的方法是将 A 的每个 10ms 的子工作视为一项独立的工作。因此，当系统启动时，它的选择是调度 10ms 的 A，还是 50ms 的 B。STCF 会选择较短的 A。然后，A 的工作已完成，只剩下 B，并开始运行。然后提交 A 的一个新子工作，它抢占 B 并运行 10ms。这样做可以实现重叠，一个进程在等待另一个进程的 I/O 完成时使用 CPU，系统因此得到更好的利用。

![alt 属性文本](https://blog.cdn.updev.cn/2020-07-26-131557.jpg)

#### 7. 多级反馈队列 MLFQ （Multi-level Feedback Queue，简称 MLFQ

> 操作系统常常不知道工作要运行多久，而这又是 SJF 等算法所必需的；而轮转调度虽然降低了响应时间，周转时间却很差。因此引入了多级反馈队列（Multi-level Feedback Queue，简称 MLFQ）。MLFQ 需要解决两方面的问题：首先它要优化周转时间，这可以通过优先执行较短的工作来实现；其次，MLFQ 希望给用户提供较好的交互体验，因此需要降低响应时间。

- MLFQ 中有许多独立的队列，每个队列有不同的优先级。任何时刻，一个工作只能存在于一个队列中。MLFQ 总是优先执行较高优先级的工作（即那些在较高级队列中的工作）。每个队列中可能会有多个工作，它们具有同样的优先级。在这种情况下，我们就对这些工作采用轮转调度。
-

###### 基本规则

- 如果 A 的优先级大于 B 的优先级，运行 A 不运行 B。
- 如果 A 的优先级等于 B 的优先级，轮转运行 A 和 B。
- 工作进入系统时，放在最高优先级（最上层）队列。这一规则使得多级反馈队列算法类似 SJF，保证了良好的响应时间。
- 一旦工作用完了其在某一层中的时间配额（无论中间主动放弃了多少次 CPU），就降低其优先级（移入低一级队列）。这一规则防止进程主动放弃 CPU，从而造成其他进程饥饿。
- 每经过一段时间，就将系统中所有工作重新加入最高优先级队列。这一规则解决了两个问题：一是防止长进程饥饿，二是如果一个 CPU 密集型工作变成了交互型，当它优先级提升时，调度程序会正确对待它。

#### 8. 比例份额

> 比例份额（proportional-share）调度程序，有时也称为公平份额（fair-share）调度程序。比例份额算法认为，调度程序的最终目标是确保每个工作获得一定比例的 CPU 时间，而不是优化周转时间和响应时间。它的基本思想很简单：每隔一段时间，都会举行一次彩票抽奖，以确定接下来应该运行哪个进程。越是应该频繁运行的进程，越是应该拥有更多地赢得彩票的机会。

##### 彩票数表示份额

> 在彩票调度中，彩票数代表了进程占有某个资源的份额。一个进程拥有的彩票数占总彩票数的百分比，就是它占有资源的份额。假设有两个进程 A 和 B，A 拥有 75 张彩票，B 拥有 25 张。因此我们希望 A 占用 75%的 CPU 时间，而 B 占用 25%. 通过不断且定时地抽取彩票，彩票调度从概率上获得这种份额比例。抽取彩票的过程很简单：调度程序知道总共的彩票数（在我们的例子中，有 100 张）。调度程序抽取中奖彩票，这是从 0 和 99 之间的一个数，拥有这个数对应的彩票的进程中奖。假设进程 A 拥有 0 到 74 共 75 张彩票，进程 B 拥有 75 到 99 的 25 张，中奖的彩票就决定了运行 A 或 B。调度程序然后加载中奖进程的状态，并运行它。彩票调度利用了随机性，这导致了从概率上满足期望的比例。随着这两个工作运行得时间越长，它们得到的 CPU 时间比例就会越接近期望。

##### 实现

> 彩票调度实现起来非常简单，只需要一个随机数生成器来选择中奖彩票和一个记录系统中所有进程的数据结构，以及所有彩票的总数。假设我们使用列表记录进程，下面的例子中有 A、B 和 C 这 3 个进程，每个进程有一定数量的彩票。