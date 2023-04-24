这应该是目前最简单的创建方法
https://github.com/electron-vite/electron-vite-vue
```bash
npm create electron-vite
# 假设项目名为myapp
cd myapp
npm i
```
在根目录下创建.npmrc文件，设置镜像
> 很重要，国内build打包必需。试错了好多个方法，这个是最简洁有效的
https://segmentfault.com/a/1190000040356146
```bash
ELECTRON_MIRROR=https://npm.taobao.org/mirrors/electron/
ELECTRON_BUILDER_BINARIES_MIRROR=http://npm.taobao.org/mirrors/electron-builder-binaries/
```

```bash
npm run dev # 调试
npm run build # 打包文件在 release文件夹下
```
