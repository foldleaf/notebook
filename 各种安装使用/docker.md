docker桌面版
https://www.docker.com/products/docker-desktop/

控制面板->程序->启用或关闭 windows 功能，开启 `Windows 虚拟机监控程序平台`和`适用于Linux的Windows子系统`
设置开机启动 Hypervisor
```bash
bcdedit /set hypervisorlaunchtype auto
```

确认任务管理器-性能-CPU面板-虚拟化为已开启，确保 BIOS 已开启虚拟化
使用镜像源
```bash
"registry-mirrors": ["https://registry.docker-cn.com"]
```
