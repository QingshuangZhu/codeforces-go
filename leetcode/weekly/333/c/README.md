# 方法一：转换成 0-1 背包方案数

把无平方因子数的数字记作 NSQ。

对于每个 $[2,30]$ 内的 NSQ，通过预处理得到对应的质数集合，用二进制表示。

二进制从低到高第 $i$ 个比特为 $1$ 表示第 $i$ 个质数在集合中，为 $0$ 表示第 $i$ 个质数不在集合中。

那么把每个是 NSQ 的 $\textit{nums}[i]$ 转换成对应的质数集合，题目就变成「选一些**不相交**的质数集合，它们的**并**集**恰好**为集合 $j$ 的方案数」。

这是 0-1 背包求方案数的模型，具体可以看[【基础算法精讲 18】](https://www.bilibili.com/video/BV16Y411v7Y6/)。

附：[本题视频讲解](https://www.bilibili.com/video/BV1jM411J7y7/)

```py [sol1-Python3]
PRIMES = 2, 3, 5, 7, 11, 13, 17, 19, 23, 29
NSQ_TO_MASK = [0] * 31  # NSQ_TO_MASK[i] 为 i 对应的质数集合（用二进制表示）
for i in range(2, 31):
    for j, p in enumerate(PRIMES):
        if i % p == 0:
            if i % (p * p) == 0:  # 有平方因子
                NSQ_TO_MASK[i] = -1
                break
            NSQ_TO_MASK[i] |= 1 << j  # 把 j 加到集合中

class Solution:
    def squareFreeSubsets(self, nums: List[int]) -> int:
        MOD = 10 ** 9 + 7
        M = 1 << len(PRIMES)
        f = [0] * M  # f[j] 表示恰好组成集合 j 的方案数
        f[0] = 1  # 空集的方案数为 1
        for x in nums:
            mask = NSQ_TO_MASK[x]
            if mask >= 0:  # x 是 NSQ
                for j in range(M - 1, mask - 1, -1):
                    if (j | mask) == j:  # mask 是 j 的子集
                        f[j] = (f[j] + f[j ^ mask]) % MOD  # 不选 mask + 选 mask
        return (sum(f) - 1) % MOD  # -1 去掉空集
```

```java [sol1-Java]
class Solution {
    private static final int[] PRIMES = new int[]{2, 3, 5, 7, 11, 13, 17, 19, 23, 29};
    private static final int MOD = (int) 1e9 + 7, MX = 30, N_PRIMES = PRIMES.length, M = 1 << N_PRIMES;
    private static final int[] NSQ_TO_MASK = new int[MX + 1]; // NSQ_TO_MASK[i] 为 i 对应的质数集合（用二进制表示）

    static {
        for (int i = 2; i <= MX; ++i)
            for (int j = 0; j < N_PRIMES; ++j) {
                int p = PRIMES[j];
                if (i % p == 0) {
                    if (i % (p * p) == 0) { // 有平方因子
                        NSQ_TO_MASK[i] = -1;
                        break;
                    }
                    NSQ_TO_MASK[i] |= 1 << j; // 把 j 加到集合中
                }
            }
    }

    public int squareFreeSubsets(int[] nums) {
        var f = new int[M]; // f[j] 表示恰好组成集合 j 的方案数
        f[0] = 1; // 空集的方案数为 1
        for (int x : nums) {
            int mask = NSQ_TO_MASK[x];
            if (mask >= 0) // x 是 NSQ
                for (int j = M - 1; j >= mask; --j)
                    if ((j | mask) == j)  // mask 是 j 的子集
                        f[j] = (f[j] + f[j ^ mask]) % MOD; // 不选 mask + 选 mask
        }
        var ans = 0L;
        for (int v : f) ans += v;
        return (int) ((ans - 1) % MOD); // -1 去掉空集
    }
}
```

```cpp [sol1-C++]
class Solution {
    static constexpr int PRIMES[] = {2, 3, 5, 7, 11, 13, 17, 19, 23, 29};
    static constexpr int MOD = 1e9 + 7, MX = 30, N_PRIMES = 10, M = 1 << N_PRIMES;
public:
    int squareFreeSubsets(vector<int> &nums) {
        int nsq2mask[MX + 1]{}; // nsq2mask[i] 为 i 对应的质数集合（用二进制表示）
        for (int i = 2; i <= MX; ++i)
            for (int j = 0; j < N_PRIMES; ++j) {
                int p = PRIMES[j];
                if (i % p == 0) {
                    if (i % (p * p) == 0) { // 有平方因子
                        nsq2mask[i] = -1;
                        break;
                    }
                    nsq2mask[i] |= 1 << j; // 把 j 加到集合中
                }
            }

        int f[M]{1}; // f[j] 表示恰好组成集合 j 的方案数，其中空集的方案数为 1
        for (int x : nums)
            if (int mask = nsq2mask[x]; mask >= 0) // x 是 NSQ
                for (int j = M - 1; j >= mask; --j)
                    if ((j | mask) == j)  // mask 是 j 的子集
                        f[j] = (f[j] + f[j ^ mask]) % MOD; // 不选 mask + 选 mask
        return (accumulate(f, f + M, 0L) - 1) % MOD; // -1 去掉空集
    }
};
```

```go [sol1-Go]
var primes = [...]int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}
var nsq2mask = [31]int{} // nsq2mask[i] 为 i 对应的质数集合（用二进制表示）

func init() {
	for i := 2; i <= 30; i++ {
		for j, p := range primes {
			if i%p == 0 {
				if i%(p*p) == 0 { // 有平方因子
					nsq2mask[i] = -1
					break
				}
				nsq2mask[i] |= 1 << j // 把 j 加到集合中
			}
		}
	}
}

func squareFreeSubsets(nums []int) int {
	const mod int = 1e9 + 7
	const m = 1 << len(primes)
	f := [m]int{1} // f[j] 表示恰好组成集合 j 的方案数，其中空集的方案数为 1
	for _, x := range nums {
		if mask := nsq2mask[x]; mask >= 0 { // x 是 NSQ
			for j := m - 1; j >= mask; j-- {
				if j|mask == j { // mask 是 j 的子集
					f[j] = (f[j] + f[j^mask]) % mod // 不选 mask + 选 mask
				}
			}
		}
	}
	ans := 0
	for _, v := range f {
		ans += v
	}
	return (ans - 1) % mod // -1 去掉空集
}
```

### 复杂度分析

- 时间复杂度：$O(n2^m)$，其中 $n$ 为 $\textit{nums}$ 的长度，$m=\pi(\max(\textit{nums}))$，$\pi(N)$ 表示 $N$ 以内的质数个数，对于本题的数据范围，$m\le 10$。
- 空间复杂度：$O(2^m)$。

# 方法二：子集状压 DP

还可以枚举 $\textit{mask}$ 的补集 $\textit{other}$ 的子集 $j$，按照「选或不选」的思想，有如下转移：

$$
f[j|mask] += f[j] \cdot \textit{cnt}[x]
$$

其中 $x$ 是 $\textit{mask}$ 对应的 NSQ，$\textit{cnt}[x]$ 为 $x$ 在 $\textit{nums}$ 中的出现次数。

代码实现时，$1$ 可以分开计算。

附：[视频讲解](https://www.bilibili.com/video/BV1jM411J7y7/)

```py [sol2-Python3]
PRIMES = 2, 3, 5, 7, 11, 13, 17, 19, 23, 29
NSQ_TO_MASK = [0] * 31  # NSQ_TO_MASK[i] 为 i 对应的质数集合（用二进制表示）
for i in range(2, 31):
    for j, p in enumerate(PRIMES):
        if i % p == 0:
            if i % (p * p) == 0:  # 有平方因子
                NSQ_TO_MASK[i] = -1
                break
            NSQ_TO_MASK[i] |= 1 << j  # 把 j 加到集合中

class Solution:
    def squareFreeSubsets(self, nums: List[int]) -> int:
        MOD = 10 ** 9 + 7
        cnt = Counter(nums)
        M = 1 << len(PRIMES)
        f = [0] * M  # f[j] 表示恰好组成集合 j 的方案数
        f[0] = 1  # 空集的方案数为 1
        for x, c in cnt.items():
            mask = NSQ_TO_MASK[x]
            if mask > 0:  # x 是 NSQ
                other = (M - 1) ^ mask  # mask 的补集
                j = other
                while True:  # 枚举 other 的子集 j
                    f[j | mask] = (f[j | mask] + f[j] * c) % MOD  # 不选 mask + 选 mask
                    j = (j - 1) & other
                    if j == other: break
        return (sum(f) * pow(2, cnt[1], MOD) - 1) % MOD  # -1 去掉空集
```

```java [sol2-Java]
class Solution {
    private static final int[] PRIMES = new int[]{2, 3, 5, 7, 11, 13, 17, 19, 23, 29};
    private static final int MOD = (int) 1e9 + 7, MX = 30, N_PRIMES = PRIMES.length, M = 1 << N_PRIMES;
    private static final int[] NSQ_TO_MASK = new int[MX + 1]; // NSQ_TO_MASK[i] 为 i 对应的质数集合（用二进制表示）

    static {
        for (int i = 2; i <= MX; ++i)
            for (int j = 0; j < N_PRIMES; ++j) {
                int p = PRIMES[j];
                if (i % p == 0) {
                    if (i % (p * p) == 0) { // 有平方因子
                        NSQ_TO_MASK[i] = -1;
                        break;
                    }
                    NSQ_TO_MASK[i] |= 1 << j;
                }
            }
    }

    public int squareFreeSubsets(int[] nums) {
        var cnt = new int[MX + 1];
        int pow2 = 1;
        for (int x : nums)
            if (x == 1) pow2 = pow2 * 2 % MOD;
            else ++cnt[x];

        var f = new long[M]; // f[j] 表示恰好组成集合 j 的方案数
        f[0] = 1; // 空集的方案数为 1
        for (int x = 2; x <= MX; ++x) {
            int mask = NSQ_TO_MASK[x], c = cnt[x];
            if (mask > 0 && c > 0) {
                int other = (M - 1) ^ mask, j = other; // mask 的补集 other
                do { // 枚举 other 的子集 j
                    f[j | mask] = (f[j | mask] + f[j] * cnt[x]) % MOD; // 不选 mask + 选 mask
                    j = (j - 1) & other;
                } while (j != other);
            }
        }
        var ans = 0L;
        for (var v : f) ans += v;
        return (int) ((ans % MOD * pow2 - 1) % MOD); // -1 去掉空集
    }
}
```

```cpp [sol2-C++]
class Solution {
    static constexpr int PRIMES[] = {2, 3, 5, 7, 11, 13, 17, 19, 23, 29};
    static constexpr int MOD = 1e9 + 7, MX = 30, N_PRIMES = 10, M = 1 << N_PRIMES;
public:
    int squareFreeSubsets(vector<int> &nums) {
        int nsq2mask[MX + 1]{};
        for (int i = 2; i <= MX; ++i)
            for (int j = 0; j < N_PRIMES; ++j) {
                int p = PRIMES[j];
                if (i % p == 0) {
                    if (i % (p * p) == 0) {
                        nsq2mask[i] = 0;
                        break;
                    }
                    nsq2mask[i] |= 1 << j;
                }
            }

        int cnt[MX + 1]{}, pow2 = 1;
        for (int x : nums)
            if (x == 1) pow2 = pow2 * 2 % MOD;
            else ++cnt[x];

        long f[M]{1}; // f[j] 表示恰好组成集合 j 的方案数，其中空集的方案数为 1
        for (int x = 2; x <= MX; ++x) {
            int mask = nsq2mask[x], c = cnt[x];
            if (mask && c) {
                int other = (M - 1) ^ mask, j = other; // mask 的补集 other
                do { // 枚举 other 的子集 j
                    f[j | mask] = (f[j | mask] + f[j] * cnt[x]) % MOD; // 不选 mask + 选 mask
                    j = (j - 1) & other;
                } while (j != other);
            }
        }
        return (accumulate(f, f + M, 0L) % MOD * pow2 - 1 + MOD) % MOD; // -1 去掉空集，+MOD 保证非负
    }
};
```

```go [sol2-Go]
var primes = [...]int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}
var nsq2mask = [31]int{} // nsq2mask[i] 为 i 对应的质数集合（用二进制表示）

func init() {
	for i := 2; i <= 30; i++ {
		for j, p := range primes {
			if i%p == 0 {
				if i%(p*p) == 0 { // 有平方因子
					nsq2mask[i] = -1
					break
				}
				nsq2mask[i] |= 1 << j // 把 j 加到集合中
			}
		}
	}
}

func squareFreeSubsets(a []int) int {
	const mod int = 1e9 + 7
	cnt, pow2 := [31]int{}, 1
	for _, x := range a {
		if x == 1 {
			pow2 = pow2 * 2 % mod
		} else {
			cnt[x]++
		}
	}

	const m = 1 << len(primes)
	f := [m]int{1} // f[j] 表示恰好组成集合 j 的方案数，其中空集的方案数为 1
	for nsq, mask := range nsq2mask {
		if mask > 0 && cnt[nsq] > 0 {
			other := (m - 1) ^ mask // mask 的补集
			for j := other; ; { // 枚举 other 的子集 j
				f[j|mask] = (f[j|mask] + f[j]*cnt[nsq]) % mod // 不选 mask + 选 mask
				j = (j - 1) & other
				if j == other {
					break
				}
			}
		}
	}
	ans := 0
	for _, v := range f {
		ans += v
	}
	return (ans%mod*pow2 - 1 + mod) % mod // -1 去掉空集，+mod 保证非负
}
```

### 复杂度分析

- 时间复杂度：$O(n+q2^m)$，其中 $n$ 为 $\textit{nums}$ 的长度，$q$ 为 $\max(\textit{nums})$ 以内的 NSQ 个数，$m=\pi(\max(\textit{nums}))$，$\pi(N)$ 表示 $N$ 以内的质数个数，对于本题的数据范围，$q\le 18, m\le 10$。
- 空间复杂度：$O(2^m)$。