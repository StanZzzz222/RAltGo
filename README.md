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

We have tested some ALTV API performance and the results show that the current Go module is no worse than C# and JS resources in ALTV API calls, and is even faster than both of them. Below we only show 5000 created using C# and Go modules respectively. If you have any objections, you can still try other resources or other content:

Csharp:

![](https://lavender-neighbouring-coral-264.mypinata.cloud/ipfs/QmcLCicAeuFBQ3XR7njaCUe4NpYGTznTtQxDWPmZJwBMcE/csharp.png)

Golang:

![](https://lavender-neighbouring-coral-264.mypinata.cloud/ipfs/QmcLCicAeuFBQ3XR7njaCUe4NpYGTznTtQxDWPmZJwBMcE/golang.png)

<font color="red">Note: Because it is developed based on altv-rust, the problems that occur in altv-rust may also occur in this module!</font>

[I need an example?]([GitHub - StanZzzz222/RAltGo-example-server: RAltGo example server](https://github.com/StanZzzz222/RAltGo-example-server))

[中文文档](ZH_CH.md)
