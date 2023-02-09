# React教程:概述和演练

自从我开始学习JavaScript时，我就一直有听说过React，但我承认我看了一眼就被它吓到了。我看到的是一堆HTML和JavaScript的混合物，心想，这种情况难道不是我们一直在努力避免的吗?React就这?

相反，我专注于学习普通的 JavaScript，并在专业环境下使用 jQuery工作，在经历几次失败的尝试后，我开始使用React，然后我终于明白了，明白为什么我想要使用React而不是普通的JS或jQuery

我试着将我所学浓缩成一篇优秀的入门教程与你们分享，所以就有了这篇文章。

### 预备知识
在开始使用React之前，有几件事你应该提前知道。例如，如果你以前从未使用过JavaScript或DOM，那应该在尝试使用React之前进一步熟悉它们。

以下是我认为React的预备知识：

>•基本熟悉HTML和CSS。
>•基本的JavaScript和编程知识。
>•对DOM有基本的了解。
>•熟悉ES6语法和特性。
>•全局安装Node.js和npm。

### 目标

>•了解React的基本概念和相关术语，如Babel, Webpack, JSX，组件，道具，状态和生命周期。

>•构建一个非常简单的React应用程序来演示上述概念。

这里会有源代码和最终结果的现场演示。

## 什么是React？
>React是一个JavaScript库——最受欢迎的库之一，在GitHub上有超过10万颗星。

>React不是一个框架(不像Angular，它更固执己见)。

>React是Facebook创建的一个开源项目。

>React用于在前端构建用户界面(UI)。

>React是MVC应用程序的视图层(模型视图控制器)

React最重要的一点是，你可以创建类似于自定义的、可重用的HTML元素的组件(`components`)，快速有效地构建用户界面。React还使用状态(`state`)和属性(`props`)简化了数据的存储和处理方式。
我们将在本文中讨论所有这些内容，现在让我们开始吧。

## 设置与安装
有几种方法可以设置React，我将展示两种，以让你了解它是如何工作的。

### 静态HTML文件
第一种方法不是一种流行的设置React的方法，也不是我们接下来教程的方法。但如果你曾经使用过jQuery这样的库就会很熟悉且容易理解这种方法。如果你不熟悉Webpack, Babel和Node.js，这是最不令人生畏的入门方法。

让我们从创建一个基本的`index.html`文件开始

没用过jQuery，基本没用，不翻译，下一个

### Create React App

我刚才使用的将JavaScript库加载到静态HTML页面并动态呈现React和Babel的方法不是很有效，而且很难维护。

幸运的是，Facebook已经创建了一个预先配置了构建React应用所需的所有内容的环境:`Create React App`。它将创建一个动态开发服务器，使用Webpack自动编译React, JSX和ES6，auto-prefix CSS文件，并使用ESLint测试和警告代码中的错误。

要设置create-react-app，选择你希望项目所在位置的上级目录，在你的终端中运行以下代码，

```bash
npx create-react-app react-tutorial
```
安装完成后，移动到新创建的目录并启动项目。

```bash
cd react-tutorial 
npm start
```
一旦你运行这个命令，你的新React应用程序将弹出一个新的窗口将在localhost:3000。

Create React App非常适合初学者和大型企业应用程序的入门，但它并不适合所有工作流。你也可以为React创建自己的Webpack设置。

如果查看项目结构，您将看到一个`/public`和`/src`目录，以及常规的`node_modules`、`.gitignore`、`README。Md`和`package.json`。

在/public中，我们的重要文件是index.html，它与我们之前创建的静态index.html文件非常相似——只是一个根div块。这一我们没有加载任何库或脚本。/src目录将包含我们所有的React代码。

要查看环境如何自动编译和更新React代码，请在/src/App.js中找到如下代码行:
```html
To get started, edit `src/App.jsànd save to reload.
新版本可能是
<p>Edit <code>src/App.js</code> and save to reload.</p>
```

然后用其他文本替换它。保存文件后，您将注意到localhost:3000编译并使用新数据进行刷新。

继续删除/src目录中的所有文件，我们将创建自己的样板文件,不需要多余的东西，只保留index.css和index.js。

对于index.css，我只是复制并粘贴原始CSS的内容到文件中。如果你愿意，你可以使用Bootstrap或任何你想要的CSS框架，或者什么都不用。我只是觉得这样更好用。

现在我们在index.js中导入React、ReactDOM和CSS文件。
```js
import React from 'react'
import ReactDOM from 'react-dom'
import './index.css'
```

让我们再次创建App组件。之前，我们只有一个`<h1>`,但现在我添加了一个div元素和一个class,您会注意到我们使用className而不是class。这是我们的第一个提示，这里编写的代码是JavaScript，而不是实际的HTML。
```js
class App extends React.Component {
  render() {
    return (
      <div className="App">
        <h1>Hello, React!</h1>
      </div>
    )
  }
}
```
最后，我们将App渲染到根目录
```js
ReactDOM.render(<App />, document.getElementById('root'))
```
下面是完整的index.js。这里我们将Component作为React的属性加载，因此我们不再需要扩展React组件。
```js
import React, { Component } from 'react'
import ReactDOM from 'react-dom'
import './index.css'
class App extends Component {
  render() {
    return (
      <div className="App">
        <h1>Hello, React!</h1>
      </div>
    )
  }
}
ReactDOM.render(<App />, document.getElementById('root'))
```
如果你回到localhost:3000，你会看到“Hello, React!”
我们的React应用程序正式起步。

有一个名为React Developer Tools的扩展，它可以让你在使用React时更加轻松。下载React DevTools for Chrome，或任何你喜欢的浏览器。

在你安装之后，当打开DevTools，你会看到一个React标签。单击它，您将能够在编写组件时检查它们。

您仍然可以转到Elements选项卡查看实际的DOM输出。现在看起来似乎没什么大不了的，但随着应用程序变得越来越复杂，它将变得越来越有必要使用。
现在我们已经有了开始使用React所需的所有工具和设置。

## JSX:JavaScript+XML
正如你所看到的，我们一直在React代码中使用看起来像HTML的东西，但它不是完全的HTML。这是JSX，代表JavaScript XML。
使用JSX，我们可以编写类似HTML的内容，还可以创建和使用我们自己的类似xml的标记。下面是JSX赋值给变量的样子。
```xml
const heading = <h1 className="site-heading">Hello, React</h1>
```
使用JSX并不是编写React的强制要求。在底层，它运行createElement，它接受标签、包含属性的对象和组件的子元素，并呈现相同的信息。下面的代码将具有与上面的JSX相同的输出。
```js
const heading = React.createElement(
    'h1', 
    { className: 'site-heading' }, 
    'Hello, React!'
    )
```
JSX实际上更接近JavaScript，而不是HTML，因此在编写它时需要注意一些关键的区别。

>•使用className代替class来添加CSS类，因为class是JavaScript中保留的关键字。
>•JSX中的属性和方法是驼峰式的——onclick将变成onClick。
>•自关闭标签必须以斜杠结尾，例如`<img/>`

JavaScript表达式也可以使用花括号嵌入到JSX中，包括变量、函数和属性。
```js
const name = 'Tania'
const heading = <h1>Hello, {name}</h1>
```
JSX比在普通JavaScript中创建和添加许多元素更容易编写和理解，这也是人们如此喜欢React的原因之一。

到目前为止，我们已经创建了一个组件——App组件。React中几乎所有的东西都由组件组成，可以是类组件，也可以是简单组件。
大多数React应用程序都有许多小组件，所有内容都加载到主应用程序组件中。组件通常也有自己的文件，所以让我们改变我们的项目来这样做。

从index.js中删除App class，因此像这样:
```js
import React from 'react'
import ReactDOM from 'react-dom'
import App from './App'
import './index.css'
ReactDOM.render(<App />, document.getElementById('root'))
```
我们将创建一个名为App.js的新文件，并将组件放在其中。
```js
import React, { Component } from 'react'
class App extends Component {
  render() {
    return (
      <div className="App">
        <h1>Hello, React!</h1>
      </div>
    )
  }
}
export default App
```
我们将组件导出为App，并在index.js中加载它。将组件分离到文件中并不是强制性的，但是如果不这样做，应用程序就会变得笨拙和难以控制。

让我们创建另一个组件。我们要创建一个表。创建Table.js，并用以下数据填充它。
```js
import React, { Component } from 'react'
class Table extends Component {
  render() {
    return (
      <table>
        <thead>
          <tr>
            <th>Name</th>
            <th>Job</th>
          </tr>
        </thead>
        <tbody>
          <tr>
            <td>Charlie</td>
            <td>Janitor</td>
          </tr>
          <tr>
            <td>Mac</td>
            <td>Bouncer</td>
          </tr>
          <tr>
            <td>Dee</td>
            <td>Aspiring actress</td>
          </tr>
          <tr>
            <td>Dennis</td>
            <td>Bartender</td>
          </tr>
        </tbody>
      </table>
    )
  }
}
export default Table
```
我们创建的这个组件是一个自定义类组件。我们将自定义组件大写以区别于常规HTML元素。

回到App.js，我们可以加载表格，首先导入它:
```js
import Table from './Table'
```
然后将其加载到App的render()中，在此之前我们有“Hello, React!”我还更改了外部容器的类。
```js
import React, { Component } from 'react'
import Table from './Table'
class App extends Component {
  render() {
    return (
      <div className="container">
        <Table />
      </div>
    )
  }
}
export default App
```
如果您检查您的活动环境，您将看到加载的Table。
现在我们已经了解了什么是自定义类组件。我们可以一次又一次地重用这个组件。但是，由于数据是硬编码到其中的，所以目前它并没有太大用处。

## 简单组件


