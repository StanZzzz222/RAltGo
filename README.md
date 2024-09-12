## 

<img src="https://lavender-neighbouring-coral-264.mypinata.cloud/ipfs/QmcLCicAeuFBQ3XR7njaCUe4NpYGTznTtQxDWPmZJwBMcE/altv.svg" title="" alt="" width="72">      <img title="" src="https://lavender-neighbouring-coral-264.mypinata.cloud/ipfs/QmbPKVJ2pNK57AQSBN4BFNtM2ukf9Lx8MveCGsVYUKzoaS" alt="" width="50" data-align="inline">   <img title="" src="https://lavender-neighbouring-coral-264.mypinata.cloud/ipfs/QmcLCicAeuFBQ3XR7njaCUe4NpYGTznTtQxDWPmZJwBMcE/rust.png" alt="" width="100" data-align="inline">    <img title="" src="https://lavender-neighbouring-coral-264.mypinata.cloud/ipfs/QmWHwr5trLNo6YccKisvZ1GKpu17EJs4a2YCJ7zFXtXyYD" alt="" width="54">  <img title="" src="https://lavender-neighbouring-coral-264.mypinata.cloud/ipfs/QmcLCicAeuFBQ3XR7njaCUe4NpYGTznTtQxDWPmZJwBMcE/go.png" alt="" width="97" data-align="inline">

## RAltGo

##### Server-side alt:V API for Go (Developed based on [altv-rust]([xxshady/altv-rust: Server-side alt:V API for Rust. Client-side via JS WASM WIP (github.com)](https://github.com/xxshady/altv-rust)))

---

#### **This module has reduced the performance overhead caused by CGO, allowing you to get closer to docking Go performance**

###### Thanks to the contributors of the altv-rust module for allowing me to make this Go module in Rust

###### 

**It is still under development and is not recommended for use in any production environment at the moment.**

<strong>But you can try to use it and raise a related issue if you encounter problems</strong>

We know many people may be concerned about the performance of this module, so we've added benchmarks to the examples resources. The current code and results are below. If you have any questions, you can try the sample server. Because resource scheduling will only be performed onTick, your performance issues will be alerted by the server-side resourceManager.

![](https://lavender-neighbouring-coral-264.mypinata.cloud/ipfs/QmSvxRdFriNN3MddHJvVc6A7E7sShdA4JoodbewPCxKZ3m)

Some APIs that have been implemented:

```
Object:
1. Blip - Create and setup section
2. Player - Spawn and setup section
3. Vehicle - Creation and setup section

Events:
1. onStart
2. onServerStarted
3. onStop
4. onPlayerConnect
5. onEnterVehicle
6. onLeaveVehicle
```



<font color="red">Note: Because it is developed based on altv-rust, the problems that occur in altv-rust may also occur in this module!</font>

[I need an example?]([GitHub - StanZzzz222/RAltGo-example-server: RAltGo example server](https://github.com/StanZzzz222/RAltGo-example-server))

[中文文档](ZH_CH.md)
