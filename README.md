## 

<img title="" src="https://lavender-neighbouring-coral-264.mypinata.cloud/ipfs/QmPnMCgJRbDKQ7S6xp4HJu9vN4XQAU5ovVK7yQC764EQGg/altv.svg" alt="" width="59"> <img title="" src="https://lavender-neighbouring-coral-264.mypinata.cloud/ipfs/QmPnMCgJRbDKQ7S6xp4HJu9vN4XQAU5ovVK7yQC764EQGg/plus.png" alt="" width="58"> <img title="" src="https://lavender-neighbouring-coral-264.mypinata.cloud/ipfs/QmPnMCgJRbDKQ7S6xp4HJu9vN4XQAU5ovVK7yQC764EQGg/rust.png" alt="" width="89"> <img title="" src="https://lavender-neighbouring-coral-264.mypinata.cloud/ipfs/QmPnMCgJRbDKQ7S6xp4HJu9vN4XQAU5ovVK7yQC764EQGg/equal.png" alt="" width="59"> <img title="" src="https://lavender-neighbouring-coral-264.mypinata.cloud/ipfs/QmPnMCgJRbDKQ7S6xp4HJu9vN4XQAU5ovVK7yQC764EQGg/go.png" alt="" width="97" data-align="inline">

## RAltGo

##### Server-side alt:V API for Go (Developed based on [altv-rust]([xxshady/altv-rust: Server-side alt:V API for Rust. Client-side via JS WASM WIP (github.com)](https://github.com/xxshady/altv-rust)))

---

#### **This module has reduced the performance overhead caused by CGO, allowing you to get closer to docking Go performance**

###### Thanks to the contributors of the altv-rust module for allowing me to make this Go module in Rust

###### 

**It is still under development and is not recommended for use in any production environment at the moment.**

<strong>You can try to use it and raise a related issue if you encounter problems</strong>

**<font color="red">Note: Since the author has other things to do, he can only try to spend as much time as possible to update. You can rest assured that I will keep updating in my free time</font>**

**We know many people may be concerned about the performance of this module, so we've added benchmarks to the examples resources. The current code and results are below. We think it is great to be able to achieve the current data. Of course, we will look for opportunities to continue to optimize it in the future as we continue to update it.**

**Note: If you have fully tested part of the benchmark performance on the sample resources, you can calculate the TPS based on the time taken and the number of processes in the benchmark performance to calculate whether it can support your server. After our many tests, we can calculate that the current TPS of this module can fully support: 0-1800 people on the server side. With continued optimization and development in the future, we hope that it can eventually support 0-4000 people. Server side. Of course, what we need now is to implement all server-side APIs**

<img title="" src="https://lavender-neighbouring-coral-264.mypinata.cloud/ipfs/QmPnMCgJRbDKQ7S6xp4HJu9vN4XQAU5ovVK7yQC764EQGg/code.png" alt="" data-align="center">

Measured screenshots:

<img title="" src="https://lavender-neighbouring-coral-264.mypinata.cloud/ipfs/QmPnMCgJRbDKQ7S6xp4HJu9vN4XQAU5ovVK7yQC764EQGg/server.png" alt="" data-align="left" width="648">

Some APIs that have been implemented:

```
Object:
1. Blip - Mostly done
2. Player - Mostly done
3. Vehicle - Mostly done
4. Ped - Done √
5. Colshape - Done √
6. Checkpoint - Done √
7. Marker- Done √
8. Object - Done √
9. VisualEntity - Done √
10. VirtualEntityGroup - Done √
11. Pools - The object pool has implemented all the above objects
12. VoiceChannel - To be implemented
12. ConnectionInfo - To be implemented
.... TODO: The remaining objects to be implemented, 
.... and the real optimization after the module is completed

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
.... TODO: The remaining events to be implemented, 
.... and the real optimization after the module is completed
```

<font color="red">Note: Because it is developed based on altv-rust, the problems that occur in altv-rust may also occur in this module!</font>

[I need an example?]([GitHub - StanZzzz222/RAltGo-example-server: RAltGo example server](https://github.com/StanZzzz222/RAltGo-example-server))

[中文文档](ZH_CH.md)
