## 

<img title="" src="https://lavender-neighbouring-coral-264.mypinata.cloud/ipfs/QmPnMCgJRbDKQ7S6xp4HJu9vN4XQAU5ovVK7yQC764EQGg/altv.svg" alt="" width="72">      <img title="" src="https://lavender-neighbouring-coral-264.mypinata.cloud/ipfs/QmPnMCgJRbDKQ7S6xp4HJu9vN4XQAU5ovVK7yQC764EQGg/plus.png" alt="" width="50" data-align="inline">   <img title="" src="https://lavender-neighbouring-coral-264.mypinata.cloud/ipfs/QmPnMCgJRbDKQ7S6xp4HJu9vN4XQAU5ovVK7yQC764EQGg/rust.png" alt="" width="100" data-align="inline">    <img title="" src="https://lavender-neighbouring-coral-264.mypinata.cloud/ipfs/QmPnMCgJRbDKQ7S6xp4HJu9vN4XQAU5ovVK7yQC764EQGg/equal.png" alt="" width="54">  <img title="" src="https://lavender-neighbouring-coral-264.mypinata.cloud/ipfs/QmPnMCgJRbDKQ7S6xp4HJu9vN4XQAU5ovVK7yQC764EQGg/go.png" alt="" width="97" data-align="inline">

## RAltGo

##### 适用于 AltV 服务端的 Go 模块 (基于 [altv-rust]([xxshady/altv-rust: Server-side alt:V API for Rust. Client-side via JS WASM WIP (github.com)](https://github.com/xxshady/altv-rust)) 进行开发)

---

##### **该模块已经尽可能的减少 CGO 造成的性能开销，使得您能更接近原生的 Go 性能**

###### 感谢 altv-rust 模块的贡献者，让我可以利用 Rust 来制作这个 Go 模块

###### 

**它依旧还在进行开发中，暂时不建议用于任何生产环境。**

**您可以尝试使用它，如果遇到问题可以提出相关 issue**

<font color="red">**注意：因作者现实也有其他事情，所以只能尽量花时间进行更新。但是您可以放心我会在空闲时间坚持保持更新**</font>

**我们知道很多人可能关心这个模块的性能，因此我们在示例资源中添加了基准测试。目前的代码和结果如下。我们认为能做到目前这样的数据已经非常棒了，当然未来我们会在持续更新的过程中找机会继续优化它。**

**注：如果您在示例资源完全测试过一部分基准性能，您可以通过基准性能的耗时与处理数量计算出TPS来计算它是否能够支撑您的服务端。在我们的多次测试下，我们能够计算得出目前该模块的 TPS 完全能够支撑：0-1800 人的服务端，随着往后的持续优化与开发我们希望它最终能够支撑 0-4000 人的服务端。当然我们目前紧要的是实现所有的服务端 API**

<img src="https://lavender-neighbouring-coral-264.mypinata.cloud/ipfs/QmPnMCgJRbDKQ7S6xp4HJu9vN4XQAU5ovVK7yQC764EQGg/code.png" title="" alt="" data-align="center">

部分基准性能实测截图：

<img src="https://lavender-neighbouring-coral-264.mypinata.cloud/ipfs/QmPnMCgJRbDKQ7S6xp4HJu9vN4XQAU5ovVK7yQC764EQGg/server.png" title="" alt="" data-align="left">

目前已实现的部分API如下：

```
对象:
1. Blip - 绝大部分已完成
2. Player - 绝大部分已完成
3. Vehicle - 绝大部分已完成
4. Ped - 绝大部分已完成
5. Colshape - 绝大部分已完成
6. Checkpoint - 绝大部分已完成
7. Marker- 绝大部分已完成
8. Object - 绝大部分已完成
9. Pools - 已完全实现上方所有对象的存储与管理
10. VisualEntity - 待实现中
11. VirtualEntityGroup - 待实现中
12. VoiceChannel - 待实现中
13. ConnectionInfo - 待实现中
.... TODO: 剩下的待实现的对象，以及模块完成后的真正优化

Events:
1. OnStart - Done
2. OnServerStarted - Done
3. OnStop - Done
4. OnPlayerConnect - Done
5. OnPlayerDisconnect - Done
6. OnEnterVehicle - Done
7. OnLeaveVehicle - Done
8. OnEnterColshape - Done
9. OnLeaveColshape - Done
10. OnClientEvent - Done
11. OnLocalEvent - Done
12. OnCommandEvent - Done
13. OnConsoleCommand - Done
14. OnEnteringVehicle - Done
15. OnNetOwnerChange - Done
16. OnChangeVehicleSeat - Done
17. OnPlayerSpawn - Done
18. OnInteriorChange - Done
19. OnDimensionChange - Done
.... TODO: 剩下的待实现的事件，以及模块完成后的真正优化
```

<font color="red">注：因基于 altv-rust 进行开发，所以 altv-rust 出现的问题该模块也都可能会出现！</font>

[我需要一个案例？]([GitHub - StanZzzz222/RAltGo-example-server: RAltGo example server](https://github.com/StanZzzz222/RAltGo-example-server))

[English](README.md)
