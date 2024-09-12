## 

<img src="https://lavender-neighbouring-coral-264.mypinata.cloud/ipfs/QmcLCicAeuFBQ3XR7njaCUe4NpYGTznTtQxDWPmZJwBMcE/altv.svg" title="" alt="" width="72">      <img title="" src="https://lavender-neighbouring-coral-264.mypinata.cloud/ipfs/QmbPKVJ2pNK57AQSBN4BFNtM2ukf9Lx8MveCGsVYUKzoaS" alt="" width="50" data-align="inline">   <img title="" src="https://lavender-neighbouring-coral-264.mypinata.cloud/ipfs/QmcLCicAeuFBQ3XR7njaCUe4NpYGTznTtQxDWPmZJwBMcE/rust.png" alt="" width="100" data-align="inline">    <img title="" src="https://lavender-neighbouring-coral-264.mypinata.cloud/ipfs/QmWHwr5trLNo6YccKisvZ1GKpu17EJs4a2YCJ7zFXtXyYD" alt="" width="54">  <img title="" src="https://lavender-neighbouring-coral-264.mypinata.cloud/ipfs/QmcLCicAeuFBQ3XR7njaCUe4NpYGTznTtQxDWPmZJwBMcE/go.png" alt="" width="97" data-align="inline">

## RAltGo

##### 适用于 AltV 服务端的 Go 模块 (基于 [altv-rust]([xxshady/altv-rust: Server-side alt:V API for Rust. Client-side via JS WASM WIP (github.com)](https://github.com/xxshady/altv-rust)) 进行开发)

---

##### **该模块已经尽可能的减少 CGO 造成的性能开销，使得您能更接近原生的 Go 性能**

###### 感谢 altv-rust 模块的贡献者，让我可以利用 Rust 来制作这个 Go 模块

###### 

**它依旧还在进行开发中，暂时不建议用于任何生产环境。**

**但是您可以尝试使用它，如果遇到问题可以提出相关 issue**

*我们知道很多人可能关心这个模块的性能，因此我们在示例资源中添加了基准测试。目前的代码和结果如下。如果您有任何疑问，可以尝试示例服务器。因为资源调度只会在onTick上进行，所以您的性能问题将会被服务端resourceManager提醒*

![](https://lavender-neighbouring-coral-264.mypinata.cloud/ipfs/QmSvxRdFriNN3MddHJvVc6A7E7sShdA4JoodbewPCxKZ3m)



<font color="red">注：因基于 altv-rust 进行开发，所以 altv-rust 出现的问题该模块也都可能会出现！</font>

[我需要一个案例？]([GitHub - StanZzzz222/RAltGo-example-server: RAltGo example server](https://github.com/StanZzzz222/RAltGo-example-server))

[English](README.md)
