## 一致性哈希算法
一致性哈希算法往往会和分布式系统相关，准确的说，是分布式缓存。

在Web服务中，缓存是介于数据库和服务端程序之间的一个东西。在网站的业务还不是很大的时候，一般不会需要这个东西，每次都可以从数据库中查询获得数据，但是随着网站的访问量增大，每次请求都访问数据库的话，数据库的压力就会非常大，导致响应变慢，于是就需要一种缓存数据库，它们的存储结构一般比较简单，而且它们的数据都是存储在内存中。典型的缓存数据库有Redis，Memcache。有了缓存之后，服务端的程序可以先从缓存中读取数据，如果发现没有数据，那么从数据库中读取，并且将它放入缓存中，局部性原理，它们下次被访问的概率会很大，因此，当再次访问该数据的时候，就可以直接从缓存中读取数据。由于磁盘和内存的性能差别是很大的，因此，命中缓存和不命中缓存的性能差距也是很大的。要增加缓存的命中概率，很显然，增加机器是一种简单有效的解决办法，但是一台机器的内存是非常有限的，因此，我们需要多台机器一起组成一个缓存集群。然后将数据分散到各个机器上，在这个分散过程中，我就需要一种hash算法。如果不加额外考虑，我们可以根据机器的台数n，给机器编号成0到n-1，使用模n的方法分散到各个机器中，根据概率的分布原理，每台机器获得的数据大致是相同的。

随着网站规模的增大，n台机器已经不足以支撑现有的数据缓存，这个时候，需要增加机器，如果还是按照取模的办法，假如现在的机器增加到n+1台，那么缓存的命中率会变成原来的1/(n+1)，这个命中率实在是太低了。根据理论计算，要保证最大的负载均衡的情况下，我们增加一台机器后的最大命中率是可以达到原来的n/(n+1)的。就相当于从原来的每台机器上拿出1/(n+1)的数据，然后放到新的机器里面。根据数学归纳法可以求得在n台机器上增加m台机器，在达到最大化负载均衡的情况下，缓存命中率最高可以达到原来的n/(n+m)。那么怎样的一种哈希算法才可以使得扩容后的命令率接近这个最高理论值呢？这就是一致性哈希算法。

现在，我们要做的其实就是设计一个满足要求的哈希算法。我们的期望是这样的，当加入新的成员（缓存机器）的时候，我们可以分别从原来的机器中抽取一小部分的数据到我们新加入的机器中，这样就保证了原有机器的数据不变化。举个形象的例子，大家一起喝可乐，开始八个人，每个人一小满杯，刚好倒完。又来了两个人，这个时候，我们不可能把所有可乐倒回一个大瓶子，然后分给大家吧，最好最公平的方式就是大家都匀出一小点给新的两个空杯子。这只是个例子，不一定实际，但是道理是一样的。在这种分布式缓存系统中，我们应该怎么做呢？在量子力学理论中，能力是一份一份的，而不是连续的。受这个启发，我们可以把机器分成一小份一小份，然后每一小份负责一个哈希范围。这样一台机器就负责多个哈希范围。这些一小份一小份，我们称之为虚拟节点。再来考虑在n台机器基础上增加m台机器的情况。其实只要把原来机器每台总份数的m/(n+m)分给新的机器就可以了。至于整个哈希范围分成多少份，这个问题需要实践和经验来获得，显然，分多了的话，管理和操作的效率下降，而分少了呢，随着机器增加容易造成机器的负载不均匀。

首先，我们把整个哈希范围设想成一条线段，一般用一个无符号的32位整型来表示整个范围。在线段上，我们可以将它分成一段一段的，每一个小线段由头和尾两个坐标来表示，同时为了避免相邻两个线段交界处重复，我们可以预先定义每一个小线段都是左闭右开。然后每台机器都会管理多个小线段。当有新的机器加入的时候，我们根据加入的机器的数量，根据最大限度负载均衡原则，计算出需要从每个现有的机器中拿出多少份给新的机器。也就是说每个机器都会维护一个自己的数据结构，里面包括自己需要管理的区间的集合，而每一次机器的变动操作都会更新这个数据机构。而当要查询一个key的数据对应哪一台机器的时候，为了加快查询速度，需要建立一张倒排索引，以便快速查到对应的机器。这种方法是我自己根据一致性哈希达到的目的而想的。它可以实现一致性哈希的要求，但是明显管理起来比较麻烦，而且具体分成多少份也是一个难定夺的值。下面，对它进行简化。

上面，每一台机器管理一条线段，很明显这样子比较复杂，对于一段区间，需要维护一个包含两个值的元组。事实上，只需要维护一个点即可，我们可以定义这个点只管辖它本身以及顺时针方向直到下一个点（不包括下一个点）的区间。但是对于最后一个点和第一个点会有点麻烦，于是，把原来的线段两头相接，这样就形成了一个环。这样，每个点都可以往顺时针方向寻找，并且能够寻找到相应的点。因此，一台机器只要维护他自己的点的列表就可以了。在上面，当增加机器的时候，我们选择的是将已有的份移向新的机器，这样需要维护一个算法，保证最大负载均衡，而实际上可以在已有的环上增加点来达到分割已有机器数据的目的。例如在在线段一中加上一个点，则线段一为分割成两个线段，这样就巧妙的把数据分割问题解决了，当然这么做之后，导致总的份数不再是确定值（其实我们根本不在意总的份数），总的份数由机器数量以及每个机器占有的份数决定。每个机器占用多少份比较合理呢？有人给出的答案是150。当然这也是个经验值。现在，现在的一致性哈希算法变成了这样：每个机器有用固定个数的点，这些点分布在一个由32位无符号整形组成的环上。每个点管理本身和顺时针方向直到下一个点的区域。如何生成这些点呢？随机。使用哈希函数对机器名字进行哈希，得到一个无符号32位整型。在Golang中使用crc32.ChecksumIEEE函数生成。因此对于一个key，需要查询它对应的机器的时候，只需要查找它逆时针方向的第一个点，然后找到这个点对应的机器即可。

下面是Golang语言的一个简单实现，仅供参考。主要借鉴了[这份代码](https://github.com/stathat/consistent)的实现。其实对于节点的查询，这里使用排序后二分查询（Golang自带sort包）。复杂度为logN，但是插入和删除的复杂度是nlogN。如果使用红黑树实现，两者的复杂度都是nlogN，从算法角度来讲，使用红黑树比较合适，但是实际上一致性哈希使用最多的是查找，而对节点的改动相对较少，因此差距并不是很大。

```go
import (
    "hash/crc32"
    "sort"
    "strconv"
    "sync"
)

type errorString struct {
    s string
}

func (e *errorString) Error() string {
    return "ConsistentError: " + e.s
}

// 定义错误类型
func ConsistentError(text string) error {
    return &errorString{text}
}

// 定义环类型
type Circle []uint32

func (c Circle) Len() int {
    return len(c)
}

func (c Circle) Less(i, j int) bool {
    return c[i] < c[j]
}

func (c Circle) Swap(i, j int) {
    c[i], c[j] = c[j], c[i]
}

type Hash func(date []byte) uint32

type Consistent struct {
    hash         Hash  // 产生uint32类型的函数
    circle       Circle  // 环
    virtualNodes int  // 虚拟节点个数，文中所说150
    virtualMap   map[uint32]string  // 点到主机的映射
    members      map[string]bool  // 主机列表
    sync.RWMutex
}

func NewConsisten() *Consistent {
    return &Consistent{
        hash:         crc32.ChecksumIEEE,
        circle:       Circle{},
        virtualNodes: 150,
        virtualMap:   make(map[uint32]string),
        members:      make(map[string]bool),
    }
}

// generate a string key for an element with an index
func (c *Consistent) eltKey(key string, idx int) string {
    return key + "|" + strconv.Itoa(idx)
}

func (c *Consistent) updateCricle() {
    c.circle = Circle{}
    for k := range c.virtualMap {
        c.circle = append(c.circle, k)
    }
    sort.Sort(c.circle)
}

func (c *Consistent) Members() []string {
    c.RLock()
    defer c.RUnlock()

    m := make([]string, len(c.members))

    var i = 0
    for k := range c.members {
        m[i] = k
        i++
    }

    return m
}

func (c *Consistent) Get(key string) string {
    hashKey := c.hash([]byte(key))
    c.RLock()
    defer c.RUnlock()

    i := c.search(hashKey)

    return c.virtualMap[c.circle[i]]
}

// search nearly vnode around key
// sort.Search uses binary search to find key
// every vnode cover its self and clockwise area
func (c *Consistent) search(key uint32) int {
    f := func(x int) bool {
        return c.circle[x] >= key
    }

    i := sort.Search(len(c.circle), f)
    i = i - 1
    if i < 0 {
        i = len(c.circle) - 1
    }
    return i
}

// this function is beautiful
func (c *Consistent) ForceSet(keys ...string) {
    mems := c.Members()
    for _, elt := range mems {
        var found = false

    FOUNDLOOP:
        for _, k := range keys {
            if k == elt {
                found = true
                break FOUNDLOOP
            }
        }
        if !found {
            c.Remove(elt)
        }
    }

    for _, k := range keys {
        c.RLock()
        _, ok := c.members[k]
        c.RUnlock()

        if !ok {
            c.Add(k)
        }
    }
}

func (c *Consistent) Add(elt string) {
    c.Lock()
    defer c.Unlock()

    if _, ok := c.members[elt]; ok {
        return
    }

    c.members[elt] = true

    for idx := 0; idx < c.virtualNodes; idx++ {
        c.virtualMap[c.hash([]byte(c.eltKey(elt, idx)))] = elt
    }

    c.updateCricle()
}

func (c *Consistent) Remove(elt string) {
    c.Lock()
    defer c.Unlock()

    if _, ok := c.members[elt]; !ok {
        return
    }

    delete(c.members, elt)

    for idx := 0; idx < c.virtualNodes; idx++ {
        delete(c.virtualMap, c.hash([]byte(c.eltKey(elt, idx))))
    }

    c.updateCricle()
}
```