## 

<img title="" src="https://lavender-neighbouring-coral-264.mypinata.cloud/ipfs/QmPnMCgJRbDKQ7S6xp4HJu9vN4XQAU5ovVK7yQC764EQGg/altv.svg" alt="" width="72">      <img title="" src="https://lavender-neighbouring-coral-264.mypinata.cloud/ipfs/QmPnMCgJRbDKQ7S6xp4HJu9vN4XQAU5ovVK7yQC764EQGg/plus.png" alt="" width="50" data-align="inline">   <img title="" src="https://lavender-neighbouring-coral-264.mypinata.cloud/ipfs/QmPnMCgJRbDKQ7S6xp4HJu9vN4XQAU5ovVK7yQC764EQGg/rust.png" alt="" width="100" data-align="inline">    <img title="" src="https://lavender-neighbouring-coral-264.mypinata.cloud/ipfs/QmPnMCgJRbDKQ7S6xp4HJu9vN4XQAU5ovVK7yQC764EQGg/equal.png" alt="" width="54">  <img title="" src="https://lavender-neighbouring-coral-264.mypinata.cloud/ipfs/QmPnMCgJRbDKQ7S6xp4HJu9vN4XQAU5ovVK7yQC764EQGg/go.png" alt="" width="97" data-align="inline">

## RAltGo

##### 适用于 AltV 服务端的 Go 模块 (基于 [altv-rust]([xxshady/altv-rust: Server-side alt:V API for Rust. Client-side via JS WASM WIP (github.com)](https://github.com/xxshady/altv-rust)) 进行开发)

---

##### **该模块已经尽可能的减少 CGO 造成的性能开销，使得您能更接近原生的 Go 性能**

###### 感谢 altv-rust 模块的贡献者，让我可以利用 Rust 来制作这个 Go 模块

###### 

**它依旧还在进行开发中，暂时不建议用于任何生产环境。**

**但是您可以尝试使用它，如果遇到问题可以提出相关 issue**

<font color="red">**注意：因作者现实也有其他事情，所以只能尽量花时间进行更新。但是您可以放心我会在空闲时间坚持保持更新**</font>

*我们知道很多人可能关心这个模块的性能，因此我们在示例资源中添加了基准测试。目前的代码和结果如下。如果您有任何疑问，可以尝试示例服务器。因为资源调度只会在onTick上进行，所以您的性能问题将会被服务端resourceManager提醒*

<img src="https://lavender-neighbouring-coral-264.mypinata.cloud/ipfs/QmPnMCgJRbDKQ7S6xp4HJu9vN4XQAU5ovVK7yQC764EQGg/code.png" title="" alt="" data-align="center">

基准性能实测截图：

<img src="https://lavender-neighbouring-coral-264.mypinata.cloud/ipfs/QmPnMCgJRbDKQ7S6xp4HJu9vN4XQAU5ovVK7yQC764EQGg/server.png" title="" alt="" data-align="left">

目前已实现的部分API如下：

```
对象:
1. Blip - 创建、设置、获取
2. Player - 创建、设置、获取、Emit
3. Vehicle - 创建、设置、获取
4. Ped - 创建、设置、获取
5. Colshape - 创建、设置、获取
6. Pools - 对象池，已实现上方所有对象
7. Marker - 待实现
8. Checkpoint - 待实现
9. VisualEntity - 待实现
10. Object - 待实现
12. VirtualEntityGroup - 待实现
12. VoiceChannel - 待实现
13. ConnectionInfo - 待实现

事件:
1. onStart
2. onServerStarted
3. onStop
4. onPlayerConnect
5. onPlayerDisconnect
6. onEnterVehicle
7. onLeaveVehicle
8. onClientEvent


```

<font color="red">注：因基于 altv-rust 进行开发，所以 altv-rust 出现的问题该模块也都可能会出现！</font>

[我需要一个案例？]([GitHub - StanZzzz222/RAltGo-example-server: RAltGo example server](https://github.com/StanZzzz222/RAltGo-example-server))

[English](README.md)
