# 项目创建
```bash
npm create vite@latest
# 项目命名为 sidebar
cd sidebar
npm i
npm run dev
```
# 资源准备
图标与字体，在index.html中head标签中引用
```html
<link
	href="https://fonts.googleapis.com/css2?family=Material+Icons"
	rel="stylesheet"
/>
        
<link 
    rel="preconnect" 
    href="https://fonts.googleapis.com" 
/>
<link 
    rel="preconnect" 
    href="https://fonts.gstatic.com" 
    crossorigin 
/>
<link
	href="https://fonts.googleapis.com/css2?family=Fira+Sans:wght@100&display=swap"
	rel="stylesheet"
/>
```
sass
```bash
npm add -D sass
```
vue-router
```bash
npm i vue-router
```
# 路由配置
以下都是在src文件夹下:
src/views/Home.vue
```html
<template>
    <main class="home-page">
        <h1>Home</h1>
        <p>This is Home page</p>
    </main>
</template>
```
src/views/About.vue类似
```html
<template>
    <main class="about-page">
        <h1>About</h1>
        <p>This is About page</p>
    </main>
</template>
```
src/main.js
```js
import { createApp } from 'vue'
// import './style.css'
import App from './App.vue'
import { createRouter,createWebHistory } from 'vue-router'
import Home from './views/Home.vue'
import About from './views/About.vue'

const router=createRouter({
    history: createWebHistory(),
    routes: [
        {
            path:'/',
            component: Home,
        },
        {
            path:'/about',
            component: About,
        }
    ]
})

createApp(App).use(router).mount('#app')
```
src/App.vue
```html
<script setup>

</script>

<template>
  <div class="app">
    <router-view/>
  </div>
</template>

<style lang="scss">

</style>
```
# 将路由配置独立出来
新建文件夹src/router
将原来main.js有关router的代码移动到`src/router/index.js`，注`import`时文件层级会发生变化;
该文件导出模块命名为router: `export default router`

```js
import { createRouter,createWebHistory } from 'vue-router'
import Home from '../views/Home.vue'
import About from '../views/About.vue'

const router=createRouter({
    history: createWebHistory(),
    routes: [
        {
            path:'/',
            component: Home,
        },
        {
            path:'/about',
            component: About,
        }
    ]
})

export default router
```
在main.js中引用
```js
import { createApp } from 'vue'
// import './style.css'
import App from './App.vue'
import router from './router'

createApp(App).use(router).mount('#app')
```

# 侧边栏准备
src/components/Sidebar.vue
```html
<template>
    <aside>
        Sidebar
    </aside>
</template>

<script setup>

</script>

<style lang="scss" scoped>

</style>
```
src/App.vue
```html
<script setup>
import Sidebar from './components/Sidebar.vue';
</script>

<template>
    <div class="app">
        <Sidebar />
        <router-view />
    </div>
</template>

<style lang="scss">
:root {
    --primary: #4ade80;
    --grey: #64748b;
    --dark: #1e293b;
    --dark-alt: #334155;
    --light: #f1f5f9;
    --sidebar-width: 300px;
}

* {
    margin: 0;
    padding: 0;
    box-sizing: border-box;
    font-family: 'Fira sans', sans-serif;
}

body {
    background: var(--light);
}

button {
    cursor: pointer;
    appearance: none;
    border: none;
    outline: none;
    background: none;
}

.app{
    display: flex;
    main{
        flex: 1 1 0;
        padding: 2rem;

        @media (max-width: 768px) {
            padding-left: 6rem;
        }
    }
}
</style>
```
主要是写整体的css样式排版，我也不是很懂
接下来重点在src/components/Sidebar.vue上

# 侧边导航栏
```html
<div class="menu">
    <router-link class="button" to="/">
        <span class="material-icons">home</span>
        <span class="text">Home</span>
    </router-link>
    <router-link class="button" to="/about">
        <span class="material-icons">description</span>
        <span class="text">About</span>
    </router-link>
</div>
```
css代码有点麻烦，忽略，只了解router-link即可

# Sidebar代码

```html
<template>
    <!-- 当is_expanded为true时，添加类名为'is_expanded'的样式 -->
    <aside :class="`${is_expanded&&'is_expanded'}`">
        <div class="logo">
            <img src="../assets/vue.svg" alt="vue"/>
        </div>
        
        <div class="menu-toggle-wrap">
            <button class="menu-toggle" @click="ToggleMenu">
                <span class="material-icons">keyboard_double_arrow_right</span>
            </button>
        </div>

        <h3>Menu</h3>
        <div class="menu">
            <router-link class="button" to="/">
                <span class="material-icons">home</span>
                <span class="text">Home</span>
            </router-link>

            <router-link class="button" to="/about">
                <span class="material-icons">description</span>
                <span class="text">About</span>
            </router-link>
        </div>

    </aside>
</template>

<script setup>
import { ref } from 'vue';
import { RouterLink } from 'vue-router';

const is_expanded=ref(false)
const ToggleMenu=()=>{
    is_expanded.value=!is_expanded.value
}

</script>

<style lang="scss" scoped>
aside{
    display: flex;
    flex-direction: column;

    background-color: var(--dark);
    color: var(--light);
    width: calc(2rem + 32px);
    overflow: hidden;
    min-height: 100vh;
    padding: 1rem;
    transition: 0.2s ease-out;

    .logo{
        margin-bottom: 1rem;
        img{
            width: 2rem;
        }
    }

    .menu-toggle-wrap{
        display: flex;
        justify-content: flex-end;
        margin-bottom: 1rem;
        
        position: relative;
        top: 0;
        transition: 0.2s ease-out;

        .menu-toggle{
            transition: 0.2s ease-out;
            .material-icons{
                font-size: 2rem;
                color: var(--light);
                transition: 0.2s ease-out;
            }

            &:hover{
                .material-icons{
                    
                    color: var(--primary);
                    transform: translateX(0.5rem);
                }
            }
        }
    }

    h3, .button .text{
        opacity: 1;
        transition: 0.3s ease-out;
    }

    .menu{
        margin: 0 -1rem;

        .button{
            display: flex;
            align-items: center;
            text-decoration: none;

            padding: 0.5rem 1rem;
            transition: 0.2s ease-out;
        }

        .material-icons{
            font-size: 2rem;
            color: var(--light);
            transition: 0.2s ease-out;
        }

        .text{
            color: var(--light);
            transition: 0.2s ease-out;
        }

        &:hover{
            background-color: var(--dark-alt);

            
                
        
        }

        
    }
    &.is_expanded{
        width: var(--sidebar-width);

        .menu-toggle-wrap{
            top: -3rem;
            .menu-toggle{
                transform: rotate(-180deg);
            }
        }
    }

    @media (max-width: 768px) {
        position: fixed;
        z-index: 99;
    }
}
</style>
```
